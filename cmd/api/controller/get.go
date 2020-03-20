/*
  Copyright (C) 2019 - 2020 MWSOFT
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
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	ctrl "github.com/superhero-match/superhero-get-match/cmd/api/model"
)

// GetMatch fetches a match.
func (ctl *Controller) GetMatch(c *gin.Context) {
	var req ctrl.GetRequest

	err := c.BindJSON(&req)
	if checkGetError(err, c) {
		ctl.Service.Logger.Error(
			"failed to bind JSON to value of type GetRequest",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.Service.TimeFormat)),
		)

		return
	}

	// Get Superhero from Cache, if empty, get from Elasticsearch.
	result, err := ctl.Service.GetMatch(req.SuperheroID, req.MatchedSuperheroID)
	if checkGetError(err, c) {
		ctl.Service.Logger.Error(
			"failed while executing service.GetMatch()",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.Service.TimeFormat)),
		)

		return
	}

	if result != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"match": result,
		})

		return
	}

	superhero, err := ctl.Service.GetESSuggestion(req.MatchedSuperheroID)
	if checkGetError(err, c) {
		ctl.Service.Logger.Error(
			"failed while executing service.GetMatch()",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.Service.TimeFormat)),
		)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"match": superhero,
	})
}

func checkGetError(err error, c *gin.Context) bool {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"match": nil,
		})

		return true
	}

	return false
}
