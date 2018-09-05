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

const (
	DefaultLimit = 20
)

type SearchHandler struct {
	request struct {
		model.Location

		SearchTerm string `form:"searchTerm"`

		keywords *model.SearchKeyword
	}
	response struct {
		common.RespError
		Data model.ItemsList    `json:"data"`
		Meta MetaSearchResponse `json:"meta"`
	}
}

func (o *SearchHandler) Bind(c *gin.Context) (errList common.ErrorRecordList) {
	if err := c.ShouldBindWith(&o.request, binding.Default(c.Request.Method, c.ContentType())); err != nil {
		errList.AddError(errCode.ErrDataUnmarshal, "search: ", err)
	}

	o.request.keywords = model.NewSearchKeywords(o.request.SearchTerm)

	return
}

func (o *SearchHandler) Validate() (errList common.ErrorRecordList) {
	if !o.request.Location.Empty() {
		if err := o.request.Location.Validate(); err != nil {
			errList.AddError(errCode.ErrDataValidation, "location: "+err.Error())
		}
	} else if o.request.keywords.Empty() {
		errList.AddError(errCode.ErrDataValidation, "empty params")
	}

	return
}

func Search(c *gin.Context) {
	var (
		handler = SearchHandler{}
		err     error
	)
	// BIND PARAMS
	if err := handler.Bind(c); err != nil {
		log.Error("categories-get: failed to bind parameters - ", err)

		handler.response.AppendError(err)
		c.JSON(http.StatusBadRequest, handler.response)
		return
	}
	// OPTIONS
	paging := &model.Paging{
		Start: 0,
		Limit: DefaultLimit,
	}
	filter := &model.SearchFilter{
		Keywords: handler.request.keywords,
		Location: handler.request.Location,
	}

	handler.response.Meta.Filter = filter
	handler.response.Meta.Paging = paging
	//
	loggedId := "search: " + filter.String()

	log.Info(loggedId + ", handle")
	// VALIDATE
	if err := handler.Validate(); err != nil {
		log.Error(loggedId+", failed to validate parameters - ", err)

		handler.response.AppendError(err)
		c.JSON(http.StatusUnprocessableEntity, handler.response)
		return
	}
	// REQUEST
	handler.response.Data, err = model.SearchModel.List(filter, paging)
	if err != nil {
		log.Error(loggedId+", failed to request an items - ", err)

		handler.response.AddError(errCode.ErrProcessSearch, err)
		c.JSON(http.StatusUnprocessableEntity, handler.response)
		return
	}

	c.JSON(http.StatusOK, handler.response)
}
