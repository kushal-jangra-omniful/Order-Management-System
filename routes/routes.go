package routes

import (
	
	"oms/controllers"
	"github.com/omniful/go_commons/http"
)

func RegisterRoutes(r *http.Server) {
	// r.POST("/orders/upload", controllers.UploadCSV)
	// r.GET("/orders", controllers.ViewOrders)
	// r.POST("/orders/inventory-check", controllers.InventoryCheck)
	r.GET("/orders/bulk", controllers.GetBulkOrders)

}
