package v1

import (
	"mrKrabsmr/web-chat/internal/app/controllers"
	"net/http"
)

func ConfigureRoutes(c *controllers.Controller, r *http.ServeMux) {
	go c.Mailing()

	r.HandleFunc("/", c.MainPage)
	r.HandleFunc("/chat", c.Chat)
}
