package admin

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/httpx"
	"github.com/gin-gonic/gin"
)

const (
	defaultDays     = 30
	maxDays         = 366
	defaultMonths   = 12
	maxMonths       = 60
	maxPage         = 100_000
	defaultPageSize = 25
	maxPageSize     = 100
)

type Handler struct {
	service *Service
	log     *slog.Logger
}

func NewHandler(service *Service, log *slog.Logger) *Handler {
	return &Handler{service: service, log: log}
}

// Summary returns the dashboard totals.
//
// @Summary  Download summary
// @Tags     Admin
// @Security BearerAuth
// @Produce  json
// @Success  200 {object} Summary
// @Failure  401 {object} httpx.ErrorResponse
// @Router   /admin/stats/summary [get]
func (h *Handler) Summary(c *gin.Context) {
	summary, err := h.service.Summary(c.Request.Context())
	h.respond(c, summary, err)
}

// Daily returns the downloads per day.
//
// @Summary  Daily downloads
// @Tags     Admin
// @Security BearerAuth
// @Produce  json
// @Param    days query int false "Number of days, today included" default(30) maximum(366)
// @Success  200 {array} SeriesPoint
// @Failure  400 {object} httpx.ErrorResponse
// @Failure  401 {object} httpx.ErrorResponse
// @Router   /admin/stats/daily [get]
func (h *Handler) Daily(c *gin.Context) {
	days, ok := intQuery(c, "days", defaultDays, maxDays)
	if !ok {
		return
	}
	series, err := h.service.Daily(c.Request.Context(), days)
	h.respond(c, series, err)
}

// Monthly returns the downloads per month.
//
// @Summary  Monthly downloads
// @Tags     Admin
// @Security BearerAuth
// @Produce  json
// @Param    months query int false "Number of months, the current one included" default(12) maximum(60)
// @Success  200 {array} SeriesPoint
// @Failure  400 {object} httpx.ErrorResponse
// @Failure  401 {object} httpx.ErrorResponse
// @Router   /admin/stats/monthly [get]
func (h *Handler) Monthly(c *gin.Context) {
	months, ok := intQuery(c, "months", defaultMonths, maxMonths)
	if !ok {
		return
	}
	series, err := h.service.Monthly(c.Request.Context(), months)
	h.respond(c, series, err)
}

// Downloads returns a page of the download log.
//
// @Summary  Download log
// @Tags     Admin
// @Security BearerAuth
// @Produce  json
// @Param    page     query int    false "Page number" default(1)
// @Param    pageSize query int    false "Page size" default(25) maximum(100)
// @Param    q        query string false "Search in IP address and country"
// @Param    os       query string false "Platform filter" Enums(windows, mac)
// @Success  200 {object} DownloadPage
// @Failure  400 {object} httpx.ErrorResponse
// @Failure  401 {object} httpx.ErrorResponse
// @Router   /admin/downloads [get]
func (h *Handler) Downloads(c *gin.Context) {
	page, ok := intQuery(c, "page", 1, maxPage)
	if !ok {
		return
	}
	pageSize, ok := intQuery(c, "pageSize", defaultPageSize, maxPageSize)
	if !ok {
		return
	}

	result, err := h.service.Downloads(c.Request.Context(), page, pageSize,
		strings.TrimSpace(c.Query("q")), c.Query("os"))
	h.respond(c, result, err)
}

func (h *Handler) respond(c *gin.Context, body any, err error) {
	if err != nil {
		h.log.Error("admin report", "path", c.FullPath(), "error", err)
		httpx.Error(c, http.StatusInternalServerError, "could not load report")
		return
	}
	c.JSON(http.StatusOK, body)
}

// intQuery parses a positive integer query parameter no larger than max,
// writing a 400 response and returning false when it is invalid.
func intQuery(c *gin.Context, key string, fallback, max int) (int, bool) {
	raw := c.Query(key)
	if raw == "" {
		return fallback, true
	}
	value, err := strconv.Atoi(raw)
	if err != nil || value < 1 || value > max {
		httpx.Error(c, http.StatusBadRequest, key+" must be an integer between 1 and "+strconv.Itoa(max))
		return 0, false
	}
	return value, true
}
