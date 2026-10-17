package download

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/httpx"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service       *Service
	countryHeader string
	log           *slog.Logger
}

func NewHandler(service *Service, countryHeader string, log *slog.Logger) *Handler {
	return &Handler{service: service, countryHeader: countryHeader, log: log}
}

// Download serves the installer for the requested platform.
//
// @Summary     Download the installer
// @Description Serves the installer of the current release. Only the request that starts a download is counted; HEAD requests, ranged requests that resume or split a transfer, and repeats within the dedup window are not.
// @Tags        Download
// @Produce     octet-stream
// @Param       os query string true "Target platform" Enums(windows, mac)
// @Success     200 {file} file
// @Failure     400 {object} httpx.ErrorResponse
// @Failure     404 {object} httpx.ErrorResponse
// @Router      /download [get]
func (h *Handler) Download(c *gin.Context) {
	installer, err := h.service.Installer(c.Query("os"))
	switch {
	case errors.Is(err, ErrUnsupportedPlatform):
		httpx.Error(c, http.StatusBadRequest, "os must be one of: windows, mac")
		return
	case errors.Is(err, ErrInstallerNotFound):
		h.log.Error("installer missing", "error", err)
		httpx.Error(c, http.StatusNotFound, "installer not available")
		return
	case err != nil:
		h.log.Error("resolve installer", "error", err)
		httpx.Error(c, http.StatusInternalServerError, "internal error")
		return
	}

	if startsDownload(c.Request) {
		err := h.service.Record(c.Request.Context(), c.Query("os"), c.ClientIP(), c.GetHeader(h.countryHeader))
		if err != nil {
			// Failing to count must not prevent the user from downloading.
			h.log.Error("record download", "error", err)
		}
	}

	c.FileAttachment(installer.Path, installer.FileName)
}

// startsDownload reports whether r begins a new download rather than
// probing it (HEAD) or fetching a later part of it (Range not at offset 0),
// which browsers and download managers do when resuming or parallelizing.
func startsDownload(r *http.Request) bool {
	if r.Method != http.MethodGet {
		return false
	}
	byteRange := strings.ReplaceAll(r.Header.Get("Range"), " ", "")
	return byteRange == "" || strings.HasPrefix(byteRange, "bytes=0-")
}
