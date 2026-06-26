package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"my-media-hub/backend/internal/errorcode"
	"my-media-hub/backend/internal/model"
	"my-media-hub/backend/internal/response"
	"my-media-hub/backend/internal/service"
)

type BehaviorAPI struct {
	svc *service.BehaviorService
}

func NewBehaviorAPI(svc *service.BehaviorService) *BehaviorAPI {
	return &BehaviorAPI{svc: svc}
}

func (h *BehaviorAPI) Register(r *gin.RouterGroup) {
	r.POST("/behavior", h.Record)
	r.GET("/behavior/statistics", h.Statistics)

	r.POST("/favorites", h.Favorite)
	r.DELETE("/favorites/:mediaId", h.Unfavorite)
	r.POST("/favorites/page", h.ListFavorites)

	r.GET("/ratings/:mediaId", h.GetRating)
	r.PUT("/ratings/:mediaId", h.Rate)

	r.POST("/history", h.View)
	r.POST("/history/page", h.ListHistory)

	r.POST("/hidden", h.Hide)
	r.DELETE("/hidden/:mediaId", h.Unhide)
	r.POST("/hidden/page", h.ListHidden)
}

func (h *BehaviorAPI) Record(c *gin.Context) {
	var req model.BehaviorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}
	if err := h.svc.Record(req.MediaID, req.BehaviorType, req.BehaviorValue, req.BehaviorSource); err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, struct{}{})
}

func (h *BehaviorAPI) Statistics(c *gin.Context) {
	stats, err := h.svc.GetStatistics()
	if err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, stats)
}

func (h *BehaviorAPI) Favorite(c *gin.Context) {
	var req model.MediaIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}
	if err := h.svc.Favorite(req.MediaID); err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, struct{}{})
}

func (h *BehaviorAPI) Unfavorite(c *gin.Context) {
	mediaID, err := strconv.ParseInt(c.Param("mediaId"), 10, 64)
	if err != nil {
		response.Error(c, errorcode.ParameterInvalid)
		return
	}
	if err := h.svc.Unfavorite(mediaID); err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, struct{}{})
}

func (h *BehaviorAPI) ListFavorites(c *gin.Context) {
	var req model.PageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}
	resp, err := h.svc.ListFavoritesPage(req.Page, req.PageSize)
	if err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, resp)
}

func (h *BehaviorAPI) GetRating(c *gin.Context) {
	mediaID, err := strconv.ParseInt(c.Param("mediaId"), 10, 64)
	if err != nil {
		response.Error(c, errorcode.ParameterInvalid)
		return
	}
	rating, err := h.svc.GetRatingByMediaID(mediaID)
	if err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, gin.H{"rating": rating})
}

func (h *BehaviorAPI) Rate(c *gin.Context) {
	mediaID, err := strconv.ParseInt(c.Param("mediaId"), 10, 64)
	if err != nil {
		response.Error(c, errorcode.ParameterInvalid)
		return
	}
	var body struct {
		Rating float64 `json:"rating"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}
	if body.Rating < 0.5 || body.Rating > 5.0 {
		response.Error(c, errorcode.RatingInvalid)
		return
	}
	if err := h.svc.Rate(mediaID, body.Rating); err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, struct{}{})
}

func (h *BehaviorAPI) View(c *gin.Context) {
	var req model.MediaIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}
	if err := h.svc.View(req.MediaID); err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, struct{}{})
}

func (h *BehaviorAPI) ListHistory(c *gin.Context) {
	var req model.PageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}
	resp, err := h.svc.ListHistoryPage(req.Page, req.PageSize)
	if err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, resp)
}

func (h *BehaviorAPI) Hide(c *gin.Context) {
	var req model.MediaIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}
	if err := h.svc.Hide(req.MediaID); err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, struct{}{})
}

func (h *BehaviorAPI) Unhide(c *gin.Context) {
	mediaID, err := strconv.ParseInt(c.Param("mediaId"), 10, 64)
	if err != nil {
		response.Error(c, errorcode.ParameterInvalid)
		return
	}
	if err := h.svc.Unhide(mediaID); err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, struct{}{})
}

func (h *BehaviorAPI) ListHidden(c *gin.Context) {
	var req model.PageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}
	resp, err := h.svc.ListHiddenPage(req.Page, req.PageSize)
	if err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, resp)
}
