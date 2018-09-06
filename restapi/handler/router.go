package handler

import (
	"net/http"

	"github.com/7phs/coding-challenge-search/config"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func DefaultRouter(conf *config.Config) http.Handler {
	log.Info("http/router: init")

	router := gin.New()

	//router.Use(gin.Recovery())
	router.Use(gin.Logger())
	if conf.Cors {
		router.Use(AllowCors())
	}
	// SEARCH
	router.GET("/search", Search)
	// HEALTH CHECK
	router.GET("/health/check", HealthCheck)

	return router
}
