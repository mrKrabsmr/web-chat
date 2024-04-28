package server

import (
	v1 "mrKrabsmr/web-chat/internal/api/v1"
	core "mrKrabsmr/web-chat/internal/app"
	"mrKrabsmr/web-chat/internal/app/controllers"
	"mrKrabsmr/web-chat/internal/configs"
	"net/http"

	"github.com/sirupsen/logrus"
)

type APIServer struct {
	logger *logrus.Logger
	config *configs.Config
	router *http.ServeMux
}

func NewAPIServer() *APIServer {
	return &APIServer{
		logger: core.GetLogger(),
		config: core.GetConfig(),
		router: http.NewServeMux(),
	}
}

func (s *APIServer) ConfigureRoutes(version int) {
	controller := controllers.NewController()

	switch version {
	case 1:
		v1.ConfigureRoutes(controller, s.router)	
	}
}

func (s *APIServer) Run() {
	if err := http.ListenAndServe(s.config.Address, s.router); err != nil {
		panic(err)
	}
}