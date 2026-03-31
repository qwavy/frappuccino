package http_adapter

import "net/http"

func Router(handlers *Handlers, r *http.ServeMux) {
	r.HandleFunc("GET /menu", handlers.Menu.GetAll)
	r.HandleFunc("GET /menu/{id}", handlers.Menu.GetById)
	r.HandleFunc("DELETE /menu/{id}", handlers.Menu.DeleteById)
	r.HandleFunc("POST /menu", handlers.Menu.CreateItem)
	r.HandleFunc("PUT /menu/{id}", handlers.Menu.UpdateItem)

	// INVENTORY
	r.HandleFunc("GET /inventory", handlers.Inventory.GetAll)
	r.HandleFunc("GET /inventory/{id}", handlers.Inventory.GetById)
	r.HandleFunc("POST /inventory", handlers.Inventory.Create)
	r.HandleFunc("PUT /inventory/{id}", handlers.Inventory.UpdateById)
	r.HandleFunc("DELETE /inventory/{id}", handlers.Inventory.DeleteById)

	// ORDERS
	r.HandleFunc("POST /orders", handlers.Order.Create)
	r.HandleFunc("GET /orders", handlers.Order.GetAll)
	r.HandleFunc("GET /orders/{id}", handlers.Order.Get)
	r.HandleFunc("DELETE /orders/{id}", handlers.Order.DeleteById)
	r.HandleFunc("PUT /orders/{id}", handlers.Order.UpdateById)
	r.HandleFunc("POST /orders/{id}/close", handlers.Order.Close)

	// REPORTS
	r.HandleFunc("GET /reports/total-sales", handlers.Report.GetTotalSales)
	r.HandleFunc("GET /reports/popular-items", handlers.Report.GetPopularItems)
}
