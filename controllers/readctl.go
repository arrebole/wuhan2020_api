package controllers

import (
	"net/http"

	"github.com/arrebole/databox/model"
	"github.com/arrebole/databox/service"
)

// ReadCtl ...
func ReadCtl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	city := r.FormValue("city")
	if city == "" {
		w.Write(model.FailResp())
		return
	}
	w.Write(model.ArchiveResp(service.GetArchives(city)))
}
