// Package version tells desktop clients which release is current.
package version

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Info is consumed by the desktop app's update check; keep the field names
// stable.
type Info struct {
	Version     string `json:"version"`
	DownloadURL string `json:"download_url"`
}

type Handler struct {
	info Info
}

func NewHandler(version, downloadURL string) *Handler {
	return &Handler{info: Info{Version: version, DownloadURL: downloadURL}}
}

// Get returns the current release.
//
// @Summary  Current release
// @Tags     Version
// @Produce  json
// @Success  200 {object} Info
// @Router   /version [get]
func (h *Handler) Get(c *gin.Context) {
	c.JSON(http.StatusOK, h.info)
}
