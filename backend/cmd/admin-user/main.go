// Command admin-user creates an admin panel user or resets its password.
//
// Usage:
//
//	admin-user -username <name>
//
// The password is read from the terminal, or from stdin when it is not a
// terminal. Database settings come from the same environment variables as
// the website.
package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/auth"
	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/config"
	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/database"
	"github.com/joho/godotenv"
	"golang.org/x/term"
)

const minPasswordLength = 12

func main() {
	username := flag.String("username", "", "admin username")
	flag.Parse()

	if err := run(*username); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(username string) error {
	if username == "" {
		return errors.New("-username is required")
	}
	_ = godotenv.Load()

	cfg, err := config.LoadDatabase()
	if err != nil {
		return err
	}

	password, err := readPassword()
	if err != nil {
		return err
	}
	if len(password) < minPasswordLength {
		return fmt.Errorf("password must be at least %d characters", minPasswordLength)
	}

	ctx := context.Background()
	log := slog.New(slog.NewTextHandler(os.Stderr, nil))
	db, err := database.Connect(ctx, cfg, log)
	if err != nil {
		return err
	}
	if err := database.Migrate(db, &auth.Admin{}); err != nil {
		return err
	}

	if err := auth.NewService(db, nil).SetPassword(ctx, username, password); err != nil {
		return err
	}
	fmt.Printf("password set for admin %q\n", username)
	return nil
}

func readPassword() (string, error) {
	fd := int(os.Stdin.Fd())
	if !term.IsTerminal(fd) {
		line, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil && !errors.Is(err, io.EOF) {
			return "", err
		}
		return strings.TrimRight(line, "\r\n"), nil
	}

	fmt.Fprint(os.Stderr, "Password: ")
	first, err := term.ReadPassword(fd)
	fmt.Fprintln(os.Stderr)
	if err != nil {
		return "", err
	}
	fmt.Fprint(os.Stderr, "Repeat password: ")
	second, err := term.ReadPassword(fd)
	fmt.Fprintln(os.Stderr)
	if err != nil {
		return "", err
	}
	if string(first) != string(second) {
		return "", errors.New("passwords do not match")
	}
	return string(first), nil
}
