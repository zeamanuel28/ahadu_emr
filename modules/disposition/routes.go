package disposition

import (
	"saas-api/core"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(parentGroup *gin.RouterGroup, ctrl *DispositionController) {
	group := parentGroup.Group("/dispositions")

	baseRoute := core.NewBaseRoute[Disposition, CreateDispositionDTO, UpdateDispositionDTO](ctrl)
	baseRoute.Register(group, core.RouteOptions{
		Tag:            "Dispositions",
		PathPrefix:     "/dispositions",
		CreateSchema:   "disposition.CreateDispositionDTO",
		UpdateSchema:   "disposition.UpdateDispositionDTO",
		ResponseSchema: "disposition.Disposition",
	})
}
