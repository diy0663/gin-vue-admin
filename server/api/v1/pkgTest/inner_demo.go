package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest"
	pkgTestReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type InnerDemoApi struct {
}

var demoService = service.ServiceGroupApp.PkgTestServiceGroup.InnerDemoService

// CreateInnerDemo 创建内部测试
// @Tags InnerDemo
// @Summary 创建内部测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest.InnerDemo true "创建内部测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /demo/createInnerDemo [post]
func (demoApi *InnerDemoApi) CreateInnerDemo(c *gin.Context) {
	var demo pkgTest.InnerDemo
	err := c.ShouldBindJSON(&demo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	demo.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(demo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := demoService.CreateInnerDemo(&demo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteInnerDemo 删除内部测试
// @Tags InnerDemo
// @Summary 删除内部测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest.InnerDemo true "删除内部测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /demo/deleteInnerDemo [delete]
func (demoApi *InnerDemoApi) DeleteInnerDemo(c *gin.Context) {
	var demo pkgTest.InnerDemo
	err := c.ShouldBindJSON(&demo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	demo.DeletedBy = utils.GetUserID(c)
	if err := demoService.DeleteInnerDemo(demo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteInnerDemoByIds 批量删除内部测试
// @Tags InnerDemo
// @Summary 批量删除内部测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除内部测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /demo/deleteInnerDemoByIds [delete]
func (demoApi *InnerDemoApi) DeleteInnerDemoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := demoService.DeleteInnerDemoByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateInnerDemo 更新内部测试
// @Tags InnerDemo
// @Summary 更新内部测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest.InnerDemo true "更新内部测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /demo/updateInnerDemo [put]
func (demoApi *InnerDemoApi) UpdateInnerDemo(c *gin.Context) {
	var demo pkgTest.InnerDemo
	err := c.ShouldBindJSON(&demo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	demo.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(demo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := demoService.UpdateInnerDemo(demo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindInnerDemo 用id查询内部测试
// @Tags InnerDemo
// @Summary 用id查询内部测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTest.InnerDemo true "用id查询内部测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /demo/findInnerDemo [get]
func (demoApi *InnerDemoApi) FindInnerDemo(c *gin.Context) {
	var demo pkgTest.InnerDemo
	err := c.ShouldBindQuery(&demo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redemo, err := demoService.GetInnerDemo(demo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redemo": redemo}, c)
	}
}

// GetInnerDemoList 分页获取内部测试列表
// @Tags InnerDemo
// @Summary 分页获取内部测试列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTestReq.InnerDemoSearch true "分页获取内部测试列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /demo/getInnerDemoList [get]
func (demoApi *InnerDemoApi) GetInnerDemoList(c *gin.Context) {
	response.Ok(c)
	return
	var pageInfo pkgTestReq.InnerDemoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := demoService.GetInnerDemoInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
