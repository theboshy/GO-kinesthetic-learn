package serviceDitpatcher

import (
	"github.com/gin-gonic/gin"
	"../serviceRouters"
)

func MapRouterGroup(router *gin.Engine)  {
	//primer grupo de empaquetado (mappings para mcs/principal)
	mappingsUrl := router.Group("mcs/principal");{
		mappingsUrl.GET("/instructions", serviceRouters.GetInstructions)
		mappingsUrl.GET("/delivery", serviceRouters.GetDelivery)
		mappingsUrl.POST("/captureRecords", serviceRouters.PostConsoleParams)
	}


	//segundo grupo de empaquetado (mappings para mcs/secondary)
	mappingsUrl2 := router.Group("mcs/secondary");{
		mappingsUrl2.GET("/instructions2", serviceRouters.GetInstructions2)
		mappingsUrl2.GET("/delivery2", serviceRouters.GetDelivery2)
	}
}
