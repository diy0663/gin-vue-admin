package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LinksRouter struct {
}

// InitLinksRouter 初始化 links表 路由信息
func (s *LinksRouter) InitLinksRouter(Router *gin.RouterGroup) {
	linksRouter := Router.Group("links").Use(middleware.OperationRecord())
	linksRouterWithoutRecord := Router.Group("links")
	var linksApi = v1.ApiGroupApp.PkgTestApiGroup.LinksApi
	{
		linksRouter.POST("createLinks", linksApi.CreateLinks)   // 新建links表
		linksRouter.DELETE("deleteLinks", linksApi.DeleteLinks) // 删除links表
		linksRouter.DELETE("deleteLinksByIds", linksApi.DeleteLinksByIds) // 批量删除links表
		linksRouter.PUT("updateLinks", linksApi.UpdateLinks)    // 更新links表
	}
	{
		linksRouterWithoutRecord.GET("findLinks", linksApi.FindLinks)        // 根据ID获取links表
		linksRouterWithoutRecord.GET("getLinksList", linksApi.GetLinksList)  // 获取links表列表
	}
}
