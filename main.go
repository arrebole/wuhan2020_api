<<<<<<< HEAD
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
=======
package main

import (
	"net/http"

	"github.com/arrebole/databox/controllers"
	
)

func main() {
	http.HandleFunc("/api/add", controllers.AddCtl)
	http.HandleFunc("/api/read", controllers.ReadCtl)
	http.HandleFunc("/api/logs", controllers.LogsCtl)
	http.ListenAndServe("0.0.0.0:80", nil)
}
>>>>>>> fe8d8edab275d63ceb9e5f4734048cd394113922
