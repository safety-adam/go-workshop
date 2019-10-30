package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

type check struct{}

func (c *check) health(ctx context.Context, w http.ResponseWriter, r *http.Request, param map[string]string) error {
	status := struct {
		Status string
	}{
		Status: "OK",
	}
	return json.NewEncoder(w).Encode(status)
}
