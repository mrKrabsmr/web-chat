package controllers

import (
	"log"
	core "mrKrabsmr/web-chat/internal/app"
	"net/http"
	"text/template"

	"github.com/sirupsen/logrus"
)

type Controller struct {
	logger *logrus.Logger
}

func NewController() *Controller {
	return &Controller{
		logger: core.GetLogger(),
	}
}

func (c *Controller) MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
