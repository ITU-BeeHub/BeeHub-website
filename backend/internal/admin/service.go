package admin

import (
	"context"
	"time"

	"github.com/ITU-BeeHub/BeeHub-website/backend/internal/download"
)

// Summary is the headline data of the admin dashboard.
type Summary struct {
	Totals
	ByPlatform []PlatformCount `json:"byPlatform"`
}

// SeriesPoint is the number of downloads in one day or month.
type SeriesPoint struct {
	Period     string           `json:"period"`
	Total      int64            `json:"total"`
	ByPlatform map[string]int64 `json:"byPlatform"`
}

// DownloadPage is one page of the download log.
type DownloadPage struct {
	Items    []download.Download `json:"items"`
	Total    int64               `json:"total"`
	Page     int                 `json:"page"`
	PageSize int                 `json:"pageSize"`
}

// Service computes the reports shown in the admin panel. Days and months
// are calendar periods in loc.
type Service struct {
	repo *Repository
	loc  *time.Location
	now  func() time.Time
}

func NewService(repo *Repository, loc *time.Location) *Service {
	return &Service{repo: repo, loc: loc, now: time.Now}
}

// Summary returns overall totals and the split per platform.
func (s *Service) Summary(ctx context.Context) (Summary, error) {
	today := startOfDay(s.now().In(s.loc))
	totals, err := s.repo.Totals(ctx, today, today.AddDate(0, 0, -6), today.AddDate(0, 0, -29))
	if err != nil {
		return Summary{}, err
	}

	byPlatform, err := s.repo.CountByPlatform(ctx)
	if err != nil {
		return Summary{}, err
	}
	return Summary{Totals: totals, ByPlatform: byPlatform}, nil
}

// Daily returns the downloads of each of the last `days` days, today included.
func (s *Service) Daily(ctx context.Context, days int) ([]SeriesPoint, error) {
	first := startOfDay(s.now().In(s.loc)).AddDate(0, 0, -(days - 1))
	periods := make([]string, days)
	for i := range periods {
		periods[i] = first.AddDate(0, 0, i).Format(time.DateOnly)
	}

	counts, err := s.repo.countByBucket(ctx, "YYYY-MM-DD", s.loc, first)
	if err != nil {
		return nil, err
	}
	return buildSeries(periods, counts), nil
}

// Monthly returns the downloads of each of the last `months` months, the
// current month included.
func (s *Service) Monthly(ctx context.Context, months int) ([]SeriesPoint, error) {
	now := s.now().In(s.loc)
	first := time.Date(now.Year(), now.Month()-time.Month(months-1), 1, 0, 0, 0, 0, s.loc)
	periods := make([]string, months)
	for i := range periods {
		periods[i] = first.AddDate(0, i, 0).Format("2006-01")
	}

	counts, err := s.repo.countByBucket(ctx, "YYYY-MM", s.loc, first)
	if err != nil {
		return nil, err
	}
	return buildSeries(periods, counts), nil
}

// Downloads returns the requested page of the download log.
func (s *Service) Downloads(ctx context.Context, page, pageSize int, search, os string) (DownloadPage, error) {
	items, total, err := s.repo.ListDownloads(ctx, DownloadFilter{
		Search: search,
		OS:     os,
		Limit:  pageSize,
		Offset: (page - 1) * pageSize,
	})
	if err != nil {
		return DownloadPage{}, err
	}
	return DownloadPage{Items: items, Total: total, Page: page, PageSize: pageSize}, nil
}

// buildSeries returns one point per period, in order, including periods
// without downloads.
func buildSeries(periods []string, counts []bucketCount) []SeriesPoint {
	index := make(map[string]int, len(periods))
	series := make([]SeriesPoint, len(periods))
	for i, period := range periods {
		index[period] = i
		series[i] = SeriesPoint{Period: period, ByPlatform: map[string]int64{}}
	}

	for _, c := range counts {
		i, ok := index[c.Bucket]
		if !ok {
			continue
		}
		series[i].Total += c.Count
		series[i].ByPlatform[c.OS] += c.Count
	}
	return series
}

func startOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
