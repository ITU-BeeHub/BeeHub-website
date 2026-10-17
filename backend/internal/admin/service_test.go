package admin

import (
	"context"
	"testing"
	"time"

	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/database/databasetest"
	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/download"
)

func TestBuildSeriesFillsEmptyPeriods(t *testing.T) {
	periods := []string{"2026-10-01", "2026-10-02", "2026-10-03"}
	counts := []bucketCount{
		{Bucket: "2026-10-01", OS: "windows", Count: 3},
		{Bucket: "2026-10-01", OS: "mac", Count: 1},
		{Bucket: "2026-10-03", OS: "windows", Count: 2},
		{Bucket: "2026-09-30", OS: "windows", Count: 9}, // outside the range
	}

	series := buildSeries(periods, counts)

	wantTotals := []int64{4, 0, 2}
	for i, point := range series {
		if point.Period != periods[i] || point.Total != wantTotals[i] {
			t.Errorf("series[%d] = %+v, want period %s total %d", i, point, periods[i], wantTotals[i])
		}
	}
	if series[0].ByPlatform["mac"] != 1 || series[0].ByPlatform["windows"] != 3 {
		t.Errorf("series[0].ByPlatform = %v", series[0].ByPlatform)
	}
}

func TestReportsAgainstDatabase(t *testing.T) {
	db := databasetest.Open(t, &download.Download{})
	istanbul, err := time.LoadLocation("Europe/Istanbul")
	if err != nil {
		t.Fatal(err)
	}
	now := time.Date(2026, 10, 17, 12, 0, 0, 0, istanbul)

	// 2026-10-16 22:30 UTC is already 2026-10-17 in Istanbul (UTC+3).
	seed := []download.Download{
		{OS: "windows", IPAddress: "10.0.0.1", CreatedAt: time.Date(2026, 10, 16, 22, 30, 0, 0, time.UTC)},
		{OS: "mac", IPAddress: "10.0.0.2", CreatedAt: now.Add(-time.Hour)},
		{OS: "windows", IPAddress: "10.0.0.1", CreatedAt: now.AddDate(0, 0, -3)},
		{OS: "windows", IPAddress: "10.0.0.3", CreatedAt: now.AddDate(0, -2, 0)},
	}
	if err := db.Create(&seed).Error; err != nil {
		t.Fatal(err)
	}

	service := NewService(NewRepository(db), istanbul)
	service.now = func() time.Time { return now }
	ctx := context.Background()

	summary, err := service.Summary(ctx)
	if err != nil {
		t.Fatalf("Summary: %v", err)
	}
	want := Totals{Total: 4, Today: 2, Last7Days: 3, Last30Days: 3, UniqueIPs: 3}
	if summary.Totals != want {
		t.Errorf("Totals = %+v, want %+v", summary.Totals, want)
	}
	if len(summary.ByPlatform) != 2 || summary.ByPlatform[0].OS != "windows" || summary.ByPlatform[0].Count != 3 {
		t.Errorf("ByPlatform = %+v", summary.ByPlatform)
	}

	daily, err := service.Daily(ctx, 7)
	if err != nil {
		t.Fatalf("Daily: %v", err)
	}
	if len(daily) != 7 || daily[6].Period != "2026-10-17" || daily[6].Total != 2 || daily[3].Total != 1 {
		t.Errorf("Daily = %+v", daily)
	}

	monthly, err := service.Monthly(ctx, 3)
	if err != nil {
		t.Fatalf("Monthly: %v", err)
	}
	if len(monthly) != 3 || monthly[0].Period != "2026-08" || monthly[0].Total != 1 || monthly[2].Total != 3 {
		t.Errorf("Monthly = %+v", monthly)
	}

	page, err := service.Downloads(ctx, 1, 2, "10.0.0.1", "")
	if err != nil {
		t.Fatalf("Downloads: %v", err)
	}
	if page.Total != 2 || len(page.Items) != 2 || !page.Items[0].CreatedAt.After(page.Items[1].CreatedAt) {
		t.Errorf("Downloads = %+v", page)
	}
}
