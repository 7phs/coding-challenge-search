package handler

import (
	"net/http"

	"github.com/7phs/coding-challenge-search/errCode"
	"github.com/7phs/coding-challenge-search/model"
	"github.com/7phs/coding-challenge-search/restapi/common"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
)

type SearchHandler struct {
	request struct {
		model.Location
		Keywords string `form:"searchTerm"`
	}
	response struct {
		common.RespError
		Data model.SearchResult `json:"data"`
	}
}

func (o *SearchHandler) Bind(c *gin.Context) (errList common.ErrorRecordList) {
	if err := c.ShouldBindWith(&o.request, binding.Default(c.Request.Method, c.ContentType())); err != nil {
		errList.AddError(errCode.ErrDataUnmarshal, "search: ", err)
	}

	return
}

func (o *SearchHandler) Validate() (errList common.ErrorRecordList) {
	if !o.request.Location.Empty() {
		if err := o.request.Location.Validate(); err != nil {
			// TODO:
		}
	} else if len(o.request.Keywords) == 0 {
		// TODO:
	}

	return
}

func Search(c *gin.Context) {
	handler := SearchHandler{}
	// BIND PARAMS
	if err := handler.Bind(c); err != nil {
		log.Error("categories-get: failed to bind parameters - ", err)

		handler.response.AppendError(err)
		c.JSON(http.StatusOK, handler.response)
		return
	}

	loggedId := "search: '" + handler.request.Keywords + "'+(" + handler.request.Location.String() + ")"

	log.Info(loggedId + ", handle")
	// VALIDATE
	if err := handler.Validate(); err != nil {
		log.Error(loggedId+", failed to validate parameters - ", err)

		handler.response.AppendError(err)
		c.JSON(http.StatusOK, handler.response)
		return
	}

	c.JSON(http.StatusOK, handler.response)
}
