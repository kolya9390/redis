package modules

import (
	"studentgit.kata.academy/Nikolai/historysearch/internal/infrastructure/responder"
	controllerauth "studentgit.kata.academy/Nikolai/historysearch/internal/modules/auth/controllerAuth"
	controllergeo "studentgit.kata.academy/Nikolai/historysearch/internal/modules/geoservis/controller_geo"
	"studentgit.kata.academy/Nikolai/historysearch/internal/modules/geoservis/repository"
	"studentgit.kata.academy/Nikolai/historysearch/internal/modules/geoservis/servis"
)

type Controller struct {
	AuthController	controllerauth.Auther
	GeoController	controllergeo.GeoServiceController
}

func NewControllers(services servis.DadataService, responder responder.Responder, geoRepo repository.GeoRepository) *Controller {

	authController := controllerauth.NewAuth(responder)
	geoController := controllergeo.NewGeoController(services,responder,geoRepo)

	return &Controller{
		AuthController: authController,
		GeoController: geoController,

	}
}
