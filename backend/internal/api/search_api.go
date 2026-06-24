package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"my-media-hub/backend/internal/errorcode"
	"my-media-hub/backend/internal/model"
	"my-media-hub/backend/internal/response"
	"my-media-hub/backend/internal/service"
)

type SearchAPI struct {
	svc *service.SearchService
}

func NewSearchAPI(svc *service.SearchService) *SearchAPI {
	return &SearchAPI{svc: svc}
}

func (h *SearchAPI) Register(r *gin.RouterGroup) {
	r.POST("/search/page", h.Search)
	r.GET("/search/suggestions", h.Suggestions)

	r.POST("/search/history/page", h.ListHistory)
	r.DELETE("/search/history/:id", h.DeleteHistory)
	r.DELETE("/search/history", h.ClearHistory)
}

func (h *SearchAPI) Search(c *gin.Context) {
	var req model.SearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}

	if req.Keyword == "" {
		response.Error(c, errorcode.SearchKeywordEmpty)
		return
	}

	resp, err := h.svc.Search(req)
	if err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}

	response.Success(c, resp)
}

func (h *SearchAPI) Suggestions(c *gin.Context) {
	keyword := c.Query("keyword")

	items, err := h.svc.Suggestions(keyword)
	if err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}

	response.Success(c, gin.H{"items": items})
}

func (h *SearchAPI) ListHistory(c *gin.Context) {
	var req model.PageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}

	resp, err := h.svc.ListHistory(req.Page, req.PageSize)
	if err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}

	response.Success(c, resp)
}

func (h *SearchAPI) DeleteHistory(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, errorcode.ParameterInvalid)
		return
	}

	if err := h.svc.DeleteHistory(id); err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}

	response.Success(c, struct{}{})
}

func (h *SearchAPI) ClearHistory(c *gin.Context) {
	if err := h.svc.ClearHistory(); err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}

	response.Success(c, struct{}{})
}
