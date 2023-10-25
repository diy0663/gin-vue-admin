package pkgTest

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InnerDemoRouter struct {
}

// InitInnerDemoRouter 初始化 内部测试 路由信息
func (s *InnerDemoRouter) InitInnerDemoRouter(Router *gin.RouterGroup) {
	demoRouter := Router.Group("demo").Use(middleware.OperationRecord())
	demoRouterWithoutRecord := Router.Group("demo")
	var demoApi = v1.ApiGroupApp.PkgTestApiGroup.InnerDemoApi
	{
		demoRouter.POST("createInnerDemo", demoApi.CreateInnerDemo)             // 新建内部测试
		demoRouter.DELETE("deleteInnerDemo", demoApi.DeleteInnerDemo)           // 删除内部测试
		demoRouter.DELETE("deleteInnerDemoByIds", demoApi.DeleteInnerDemoByIds) // 批量删除内部测试
		demoRouter.PUT("updateInnerDemo", demoApi.UpdateInnerDemo)              // 更新内部测试
	}
	{
		demoRouterWithoutRecord.GET("findInnerDemo", demoApi.FindInnerDemo) // 根据ID获取内部测试
		// http://localhost:8888/demo/getInnerDemoList
		demoRouterWithoutRecord.GET("getInnerDemoList", demoApi.GetInnerDemoList) // 获取内部测试列表
	}
}
