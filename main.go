package main

import (
	"net/http"

	"github.com/arrebole/databox/controllers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	http.HandleFunc("/api/add", controllers.AddCtl)
	http.HandleFunc("/api/read", controllers.ReadCtl)
	http.HandleFunc("/api/logs", controllers.LogsCtl)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
