package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/form_common/app/api/category"
)

func Router(group *gin.RouterGroup) {
	categoryRouters(group)
}

func categoryRouters(group *gin.RouterGroup) {
	categoryRouter := group.Group("category")
	{
		categoryRouter.POST("create", category.Create)
		categoryRouter.POST("edit", category.Edit)
		categoryRouter.POST("delete", category.Delete)
		categoryRouter.GET("list", category.List)
		categoryRouter.GET("detail", category.Detail)
	}
}
