/*
  Copyright (C) 2019 - 2022 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/superhero-match/superhero-get-match/cmd/api/service"
	"github.com/superhero-match/superhero-get-match/internal/config"
)

// Controller holds the Controller data.
type Controller struct {
	Service        service.Service
	MatchKeyFormat string
	Logger         *zap.Logger
	TimeFormat     string
}

// NewController returns new controller.
func NewController(cfg *config.Config) (*Controller, error) {
	srv, err := service.NewService(cfg)
	if err != nil {
		return nil, err
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	defer logger.Sync()

	return &Controller{
		Service:        srv,
		MatchKeyFormat: cfg.Cache.MatchKeyFormat,
		Logger:         logger,
		TimeFormat:     cfg.App.TimeFormat,
	}, nil
}

// RegisterRoutes registers the get_match API route.
func (ctl *Controller) RegisterRoutes() *gin.Engine {
	router := gin.Default()

	sr := router.Group("/api/v1/match")

	// Middleware.
	// sr.Use(c.Authorize)

	// Routes.
	sr.POST("/get_match", ctl.GetMatch)
	sr.GET("/health", ctl.Health)

	return router
}
