package download

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/database/databasetest"
	"gorm.io/gorm"
)

func TestRecordUniqueCountsConcurrentBurstOnce(t *testing.T) {
	db := databasetest.Open(t, &Download{})
	repo := NewRepository(db)
	now := time.Now()

	const requests = 25
	warmUpConnections(t, db, requests)

	var wg sync.WaitGroup
	start := make(chan struct{})
	stored := make(chan bool, requests)
	for range requests {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			ok, err := repo.RecordUnique(context.Background(),
				&Download{OS: "windows", IPAddress: "203.0.113.7", CreatedAt: now}, 10*time.Minute)
			if err != nil {
				t.Errorf("RecordUnique: %v", err)
			}
			stored <- ok
		}()
	}
	close(start)
	wg.Wait()
	close(stored)

	storedCount := 0
	for ok := range stored {
		if ok {
			storedCount++
		}
	}
	if storedCount != 1 {
		t.Errorf("stored %d downloads, want 1", storedCount)
	}

	var rows int64
	db.Model(&Download{}).Count(&rows)
	if rows != 1 {
		t.Errorf("table has %d rows, want 1", rows)
	}
}

// warmUpConnections opens n pooled connections up front so the goroutines
// of a test start their transactions at the same moment instead of one
// after another as each connection is established.
func warmUpConnections(t *testing.T, db *gorm.DB, n int) {
	t.Helper()
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(n)

	var wg sync.WaitGroup
	for range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sqlDB.Exec("SELECT pg_sleep(0.1)")
		}()
	}
	wg.Wait()
}

func TestRecordUniqueWindow(t *testing.T) {
	db := databasetest.Open(t, &Download{})
	repo := NewRepository(db)
	ctx := context.Background()
	start := time.Now()
	window := 10 * time.Minute

	tests := []struct {
		name string
		d    Download
		want bool
	}{
		{"first download", Download{OS: "windows", IPAddress: "198.51.100.1", CreatedAt: start}, true},
		{"repeat inside window", Download{OS: "windows", IPAddress: "198.51.100.1", CreatedAt: start.Add(5 * time.Minute)}, false},
		{"other platform", Download{OS: "mac", IPAddress: "198.51.100.1", CreatedAt: start.Add(5 * time.Minute)}, true},
		{"other IP", Download{OS: "windows", IPAddress: "198.51.100.2", CreatedAt: start.Add(5 * time.Minute)}, true},
		{"repeat after window", Download{OS: "windows", IPAddress: "198.51.100.1", CreatedAt: start.Add(11 * time.Minute)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.RecordUnique(ctx, &tt.d, window)
			if err != nil {
				t.Fatalf("RecordUnique: %v", err)
			}
			if got != tt.want {
				t.Errorf("stored = %v, want %v", got, tt.want)
			}
		})
	}
}
