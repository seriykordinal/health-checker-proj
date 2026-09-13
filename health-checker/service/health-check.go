package service

import (
	"context"
	"health-checker/config"
	"net/http"
	"sync"
)

type HealthCheck struct {
	cfg    *config.Config
	client *http.Client
}

func NewHealthCheck(cfg *config.Config) *HealthCheck {
	return &HealthCheck{
		cfg: cfg,
		client: &http.Client{
			Timeout: cfg.RequestTimeout(),
		},
	}

}

func (hc *HealthCheck) Check(ctx context.Context) map[string]string {
	results := make(map[string]string, len(hc.cfg.URLs()))
	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)

	for _, url := range hc.cfg.URLs() {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()

			status := hc.checkOne(ctx, u)

			mu.Lock()
			results[u] = status
			mu.Unlock()
		}(url)
	}

	wg.Wait()
	return results
}

func (hc *HealthCheck) checkOne(ctx context.Context, base string) string {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base+hc.cfg.HealthPath(), nil)
	if err != nil {
		return "down"
	}

	resp, err := hc.client.Do(req)
	if err != nil {
		return "down"
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return "up"
	}
	return "down"
}
