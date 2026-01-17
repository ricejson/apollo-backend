package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ricejson/apollo-backend/service/toggle"
	"net/http"
)

type ToggleController struct {
	toggleService toggle.ToggleService
}

func NewToggleController(toggleService toggle.ToggleService) *ToggleController {
	return &ToggleController{
		toggleService: toggleService,
	}
}

func (t *ToggleController) FindAll(ctx *gin.Context) {
	toggles, err := t.toggleService.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"Errno":  -1,
			"ErrMsg": "获取开关数据失败！",
		})
		// TODO 记录日志、埋点、监控
		return
	}
	ctx.JSON(http.StatusOK, map[string]any{
		"Errno":  0,
		"ErrMsg": "",
		"Data":   toggles,
	})
}

func (t *ToggleController) RegisterServices(server *gin.Engine) {
	group := server.Group("/toggles")
	group.GET("/all", t.FindAll)
}
