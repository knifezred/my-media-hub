package api

import (
	"github.com/gin-gonic/gin"
	"my-media-hub/backend/internal/errorcode"
	"my-media-hub/backend/internal/response"
	"my-media-hub/backend/internal/service"
)

type ScannerAPI struct {
	svc *service.ScannerService
}

func NewScannerAPI(svc *service.ScannerService) *ScannerAPI {
	return &ScannerAPI{svc: svc}
}

func (h *ScannerAPI) Register(r *gin.RouterGroup) {
	r.POST("/scanner/start", h.Start)
	r.GET("/scanner/status", h.Status)
}

type scanStartRequest struct {
	Directories []string `json:"directories"`
}

func (h *ScannerAPI) Start(c *gin.Context) {
	var req scanStartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, errorcode.ValidationError)
		return
	}

	if len(req.Directories) == 0 {
		response.Error(c, errorcode.ValidationError)
		return
	}

	if err := h.svc.Start(req.Directories); err != nil {
		response.Error(c, errorcode.ScanInProgress)
		return
	}

	response.Success(c, struct{}{})
}

func (h *ScannerAPI) Status(c *gin.Context) {
	status := h.svc.Status()
	response.Success(c, status)
}
