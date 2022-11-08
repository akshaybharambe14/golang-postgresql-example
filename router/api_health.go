package router

import (
	"encoding/json"
	"net/http"
	"time"
)

type HealthResp struct {
	StartedAt time.Time `json:"started_at,omitempty"`
	Uptime    string    `json:"uptime,omitempty"`
}

var startedAt = time.Now()

func Health(w http.ResponseWriter, r *http.Request) {
	resp := HealthResp{
		StartedAt: startedAt,
		Uptime:    time.Now().Sub(startedAt).String(),
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
