package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"my-media-hub/backend/internal/errorcode"
	"my-media-hub/backend/internal/model"
	"my-media-hub/backend/internal/response"
	"my-media-hub/backend/internal/service"
)

type MediaAPI struct {
	svc *service.MediaService
}

func NewMediaAPI(svc *service.MediaService) *MediaAPI {
	return &MediaAPI{svc: svc}
}

func (h *MediaAPI) Register(r *gin.RouterGroup) {
	r.POST("/media/page", h.List)
	r.GET("/media/:id", h.GetByID)

	r.POST("/tags/page", h.ListTags)
	r.GET("/tags/:id", h.GetTagByID)

	r.POST("/categories/page", h.ListCategories)
}

func (h *MediaAPI) List(c *gin.Context) {
	var req model.MediaPageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}

	resp, err := h.svc.List(req)
	if err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}

	response.Success(c, resp)
}

func (h *MediaAPI) GetByID(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		response.Error(c, errorcode.ParameterInvalid)
		return
	}

	detail, err := h.svc.GetByID(id)
	if err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	if detail == nil {
		response.Error(c, errorcode.NotFound)
		return
	}

	response.Success(c, detail)
}

func (h *MediaAPI) ListTags(c *gin.Context) {
	var req model.PageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}

	resp, err := h.svc.ListTags(req.Page, req.PageSize)
	if err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}

	response.Success(c, resp)
}

func (h *MediaAPI) GetTagByID(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		response.Error(c, errorcode.ParameterInvalid)
		return
	}

	tag, err := h.svc.GetTagByID(id)
	if err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	if tag == nil {
		response.Error(c, errorcode.NotFound)
		return
	}

	response.Success(c, tag)
}

func (h *MediaAPI) ListCategories(c *gin.Context) {
	var req model.PageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}

	resp, err := h.svc.ListCategories(req.Page, req.PageSize)
	if err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}

	response.Success(c, resp)
}

func parseID(idStr string) (int64, error) {
	var id int64
	for _, b := range []byte(idStr) {
		if b < '0' || b > '9' {
			return 0, fmt.Errorf("invalid id: %s", idStr)
		}
		id = id*10 + int64(b-'0')
	}
	return id, nil
}
