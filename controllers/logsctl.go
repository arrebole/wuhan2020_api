package controllers

import (
	"net/http"
	"strconv"

	"github.com/arrebole/databox/model"
	"github.com/arrebole/databox/service"
)

// LogsCtl ...
func LogsCtl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	limit := r.FormValue("limit")
	if limit == "" {
		limit = "10"
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}
	w.Write(model.LogResp(service.GetLogs(limitInt)))
}
