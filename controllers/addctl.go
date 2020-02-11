package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/arrebole/databox/model"
	"github.com/arrebole/databox/service"
)

// AddCtl ...
func AddCtl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		w.Write(model.FailResp())
		return
	}

	var postdata model.PostData
	err := json.NewDecoder(r.Body).Decode(&postdata)
	if err != nil {
		w.Write(model.FailResp())
		return
	}

	if postdata.Archive.City == "" || postdata.Archive.Title == "" {
		w.Write(model.FailResp())
		return
	}
	if !service.SaveArchive(postdata.GetArchive()){
		w.Write(model.FailResp())
		return
	}
	service.SaveLog(postdata.GetLog(r.RemoteAddr))
	w.Write(model.SuccessResp())
}
