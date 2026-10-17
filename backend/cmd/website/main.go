// Command website runs the BeeHub website API.
package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/ITU-BeeHub/BeeHub-website/backend/docs"
	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/admin"
	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/auth"
	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/config"
	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/database"
	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/download"
	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/server"
	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/version"
	"github.com/joho/godotenv"
)

// @title                      BeeHub Website API
// @version                    2.0
// @description                Installer downloads, release information and the admin reporting API.
// @BasePath                   /api
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization

const shutdownTimeout = 15 * time.Second

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	if err := run(log); err != nil {
		log.Error("server stopped", "error", err)
		os.Exit(1)
	}
}

func run(log *slog.Logger) error {
	// A .env file is optional and only meant for local development; in
	// production the environment comes from the systemd unit.
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	db, err := database.Connect(ctx, cfg.Database, log)
	if err != nil {
		return err
	}
	if err := database.Migrate(db, &download.Download{}, &auth.Admin{}); err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	tokens := auth.NewTokenManager(cfg.Auth.JWTSecret, cfg.Auth.TokenTTL)
	downloads := download.NewService(
		download.NewCatalog(cfg.Download.InstallersDir, cfg.Download.AppVersion),
		download.NewRepository(db),
		cfg.Download.DedupWindow,
		log,
	)

	router, err := server.NewRouter(server.Dependencies{
		Log:            log,
		TrustedProxies: cfg.TrustedProxies,
		SwaggerEnabled: cfg.SwaggerEnabled,
		Database:       sqlDB,
		Tokens:         tokens,
		Auth:           auth.NewHandler(auth.NewService(db, tokens), log),
		Download:       download.NewHandler(downloads, cfg.Download.CountryHeader, log),
		Version:        version.NewHandler(cfg.Download.AppVersion, cfg.PublicURL),
		Admin:          admin.NewHandler(admin.NewService(admin.NewRepository(db), cfg.StatsLocation), log),
	})
	if err != nil {
		return err
	}

	srv := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       2 * time.Minute,
		// No write timeout: installers are large and slow clients need
		// minutes to download them.
	}

	serveErr := make(chan error, 1)
	go func() {
		log.Info("server listening", "addr", cfg.HTTPAddr, "version", cfg.Download.AppVersion)
		serveErr <- srv.ListenAndServe()
	}()

	select {
	case err := <-serveErr:
		return err
	case <-ctx.Done():
	}

	log.Info("shutting down")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
