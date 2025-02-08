package routes

import (
	"oms/controllers"

	"github.com/omniful/go_commons/http"
)

func RegisterRoutes(r *http.Server) {
	
	r.GET("/orders/bulk", controllers.BulkOrder)
}
