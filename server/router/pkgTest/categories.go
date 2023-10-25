package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CategoriesRouter struct {
}

// InitCategoriesRouter 初始化 categories表 路由信息
func (s *CategoriesRouter) InitCategoriesRouter(Router *gin.RouterGroup) {
	categoriesRouter := Router.Group("categories").Use(middleware.OperationRecord())
	categoriesRouterWithoutRecord := Router.Group("categories")
	var categoriesApi = v1.ApiGroupApp.PkgTestApiGroup.CategoriesApi
	{
		categoriesRouter.POST("createCategories", categoriesApi.CreateCategories)   // 新建categories表
		categoriesRouter.DELETE("deleteCategories", categoriesApi.DeleteCategories) // 删除categories表
		categoriesRouter.DELETE("deleteCategoriesByIds", categoriesApi.DeleteCategoriesByIds) // 批量删除categories表
		categoriesRouter.PUT("updateCategories", categoriesApi.UpdateCategories)    // 更新categories表
	}
	{
		categoriesRouterWithoutRecord.GET("findCategories", categoriesApi.FindCategories)        // 根据ID获取categories表
		categoriesRouterWithoutRecord.GET("getCategoriesList", categoriesApi.GetCategoriesList)  // 获取categories表列表
	}
}
