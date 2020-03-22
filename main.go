package main

import (
	"net/http"

	"github.com/arrebole/wuhan2020_api/controllers"
)

func main() {
	http.HandleFunc("/api/add", controllers.AddCtl)
	http.HandleFunc("/api/read", controllers.ReadCtl)
	http.HandleFunc("/api/read/all", controllers.ReadAllCtl)
	http.HandleFunc("/api/logs", controllers.LogsCtl)
	http.ListenAndServe("0.0.0.0:80", nil)
}
