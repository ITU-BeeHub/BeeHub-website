package download

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

type fakeStore struct {
	recorded []Download
}

func (f *fakeStore) RecordUnique(_ context.Context, d *Download, _ time.Duration) (bool, error) {
	f.recorded = append(f.recorded, *d)
	return true, nil
}

func newTestRouter(t *testing.T) (*gin.Engine, *fakeStore) {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "installer_1.2.1.exe"), []byte("windows installer"), 0o644); err != nil {
		t.Fatal(err)
	}

	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	store := &fakeStore{}
	service := NewService(NewCatalog(dir, "1.2.1"), store, time.Minute, log)
	handler := NewHandler(service, "CF-IPCountry", log)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/download", handler.Download)
	r.HEAD("/download", handler.Download)
	return r, store
}

func TestDownload(t *testing.T) {
	tests := []struct {
		name        string
		method      string
		query       string
		header      http.Header
		wantStatus  int
		wantCounted bool
	}{
		{"full download", http.MethodGet, "os=windows", nil, http.StatusOK, true},
		{"range from start", http.MethodGet, "os=windows", http.Header{"Range": {"bytes=0-"}}, http.StatusPartialContent, true},
		{"range resume", http.MethodGet, "os=windows", http.Header{"Range": {"bytes=5-"}}, http.StatusPartialContent, false},
		{"head probe", http.MethodHead, "os=windows", nil, http.StatusOK, false},
		{"unknown platform", http.MethodGet, "os=linux", nil, http.StatusBadRequest, false},
		{"missing platform", http.MethodGet, "", nil, http.StatusBadRequest, false},
		{"installer missing", http.MethodGet, "os=mac", nil, http.StatusNotFound, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, store := newTestRouter(t)
			req := httptest.NewRequest(tt.method, "/download?"+tt.query, nil)
			for key, values := range tt.header {
				req.Header[key] = values
			}
			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
			if counted := len(store.recorded) == 1; counted != tt.wantCounted {
				t.Errorf("counted = %v, want %v", counted, tt.wantCounted)
			}
		})
	}
}

func TestDownloadServesAttachmentAndRecordsClient(t *testing.T) {
	r, store := newTestRouter(t)
	req := httptest.NewRequest(http.MethodGet, "/download?os=windows", nil)
	req.Header.Set("CF-IPCountry", "tr")
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if got, want := rec.Header().Get("Content-Disposition"), `attachment; filename="BeeHub-1.2.1.exe"`; got != want {
		t.Errorf("Content-Disposition = %q, want %q", got, want)
	}
	if got := rec.Body.String(); got != "windows installer" {
		t.Errorf("body = %q", got)
	}
	if len(store.recorded) != 1 {
		t.Fatalf("recorded %d downloads, want 1", len(store.recorded))
	}
	if got := store.recorded[0]; got.OS != "windows" || got.Country != "TR" || got.IPAddress == "" {
		t.Errorf("recorded %+v", got)
	}
}

func TestNormalizeCountry(t *testing.T) {
	tests := map[string]string{"TR": "TR", " de ": "DE", "XX": "", "": "", "TUR": ""}
	for in, want := range tests {
		if got := normalizeCountry(in); got != want {
			t.Errorf("normalizeCountry(%q) = %q, want %q", in, got, want)
		}
	}
}
