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
	behaviorSvc *service.BehaviorService
}

func NewBehaviorAPI(behaviorSvc *service.BehaviorService) *BehaviorAPI {
	return &BehaviorAPI{behaviorSvc: behaviorSvc}
}

func (h *BehaviorAPI) Register(r *gin.RouterGroup) {
	r.POST("/behavior", h.Record)
	r.GET("/behavior/statistics", h.Statistics)

	r.GET("/favorites", h.ListFavoritesGet)
	r.POST("/favorites", h.AddFavorite)
	r.DELETE("/favorites/:mediaId", h.RemoveFavorite)
	r.POST("/favorites/page", h.ListFavoritesPage)

	r.GET("/ratings/:mediaId", h.GetRatingForMedia)
	r.PUT("/ratings/:mediaId", h.RateMedia)
	r.POST("/ratings", h.Rate)

	r.GET("/history", h.ListHistory)
	r.POST("/history", h.MarkViewed)
	r.POST("/viewed", h.AddViewed)
	r.POST("/viewed/page", h.ListViewedPage)

	r.GET("/hidden", h.ListHiddenGet)
	r.POST("/hidden", h.AddHidden)
	r.DELETE("/hidden/:mediaId", h.RemoveHidden)
}

func (h *BehaviorAPI) Record(c *gin.Context) {
	var req model.BehaviorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}
	if err := h.behaviorSvc.Record(req.MediaID, req.BehaviorType, req.Score); err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, struct{}{})
}

func (h *BehaviorAPI) Statistics(c *gin.Context) {
	stats, err := h.behaviorSvc.GetStatistics()
	if err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, stats)
}

func (h *BehaviorAPI) ListFavoritesGet(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	resp, err := h.behaviorSvc.ListFavoritesAsMedia(page, pageSize)
	if err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, resp)
}

func (h *BehaviorAPI) AddFavorite(c *gin.Context) {
	var req model.MediaIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}
	if err := h.behaviorSvc.Record(req.MediaID, "favorite", 1); err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, struct{}{})
}

func (h *BehaviorAPI) RemoveFavorite(c *gin.Context) {
	mediaID, err := strconv.ParseInt(c.Param("mediaId"), 10, 64)
	if err != nil {
		response.Error(c, errorcode.ParameterInvalid)
		return
	}
	if err := h.behaviorSvc.RemoveBehavior(mediaID, "favorite"); err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, struct{}{})
}

func (h *BehaviorAPI) ListFavoritesPage(c *gin.Context) {
	var req model.FavoritePageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}
	resp, err := h.behaviorSvc.ListFavoritesAsMedia(req.Page, req.PageSize)
	if err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, resp)
}

func (h *BehaviorAPI) GetRatingForMedia(c *gin.Context) {
	mediaID, err := strconv.ParseInt(c.Param("mediaId"), 10, 64)
	if err != nil {
		response.Error(c, errorcode.ParameterInvalid)
		return
	}
	rating, err := h.behaviorSvc.GetRating(mediaID)
	if err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, gin.H{"rating": rating})
}

func (h *BehaviorAPI) RateMedia(c *gin.Context) {
	mediaID, err := strconv.ParseInt(c.Param("mediaId"), 10, 64)
	if err != nil {
		response.Error(c, errorcode.ParameterInvalid)
		return
	}
	var body struct {
		Rating int `json:"rating"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}
	if body.Rating < 1 || body.Rating > 5 {
		response.Error(c, errorcode.RatingInvalid)
		return
	}
	if err := h.behaviorSvc.Record(mediaID, "rating", float64(body.Rating)); err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, struct{}{})
}

func (h *BehaviorAPI) Rate(c *gin.Context) {
	var req model.RateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}
	if req.Rating < 1 || req.Rating > 5 {
		response.Error(c, errorcode.RatingInvalid)
		return
	}
	if err := h.behaviorSvc.Record(req.MediaID, "rating", float64(req.Rating)); err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, struct{}{})
}

func (h *BehaviorAPI) ListHistory(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	resp, err := h.behaviorSvc.ListViewedAsMedia(page, pageSize)
	if err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, resp)
}

func (h *BehaviorAPI) MarkViewed(c *gin.Context) {
	var req model.MediaIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}
	if err := h.behaviorSvc.Record(req.MediaID, "view", 1); err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, struct{}{})
}

func (h *BehaviorAPI) AddViewed(c *gin.Context) {
	var req model.MediaIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}
	if err := h.behaviorSvc.Record(req.MediaID, "view", 1); err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, struct{}{})
}

func (h *BehaviorAPI) ListViewedPage(c *gin.Context) {
	var req model.PageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}
	resp, err := h.behaviorSvc.ListViewedAsMedia(req.Page, req.PageSize)
	if err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, resp)
}

func (h *BehaviorAPI) ListHiddenGet(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	resp, err := h.behaviorSvc.ListHiddenAsMedia(page, pageSize)
	if err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, resp)
}

func (h *BehaviorAPI) AddHidden(c *gin.Context) {
	var req model.MediaIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}
	if err := h.behaviorSvc.Record(req.MediaID, "hidden", 1); err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, struct{}{})
}

func (h *BehaviorAPI) RemoveHidden(c *gin.Context) {
	mediaID, err := strconv.ParseInt(c.Param("mediaId"), 10, 64)
	if err != nil {
		response.Error(c, errorcode.ParameterInvalid)
		return
	}
	if err := h.behaviorSvc.RemoveBehavior(mediaID, "hidden"); err != nil {
		response.Error(c, errorcode.InternalError)
		return
	}
	response.Success(c, struct{}{})
}
