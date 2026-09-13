package handlers

import (
	"encoding/json"
	"health-checker/config"
	"health-checker/service"
	"net/http"
)

type Router struct {
	mux *http.ServeMux
	cfg *config.Config
}

func NewRouter(cfg *config.Config) *Router {
	return &Router{
		mux: http.NewServeMux(),
		cfg: cfg,
	}
}

func (r *Router) RegisterRoutes() {
	r.mux.HandleFunc("/health", r.HealthCheckHandler)
	r.mux.HandleFunc("/check", r.checkURLHandler)
}

func (r *Router) HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "up"})
}

func (r *Router) checkURLHandler(w http.ResponseWriter, req *http.Request) {

	hc := service.NewHealthCheck(r.cfg)
	results := hc.Check(req.Context())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
