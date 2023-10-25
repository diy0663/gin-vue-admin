package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    pkgTestReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type LinksApi struct {
}

var linksService = service.ServiceGroupApp.PkgTestServiceGroup.LinksService


// CreateLinks 创建links表
// @Tags Links
// @Summary 创建links表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest.Links true "创建links表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /links/createLinks [post]
func (linksApi *LinksApi) CreateLinks(c *gin.Context) {
	var links pkgTest.Links
	err := c.ShouldBindJSON(&links)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    links.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Url":{utils.NotEmpty()},
    }
	if err := utils.Verify(links, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := linksService.CreateLinks(&links); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLinks 删除links表
// @Tags Links
// @Summary 删除links表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest.Links true "删除links表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /links/deleteLinks [delete]
func (linksApi *LinksApi) DeleteLinks(c *gin.Context) {
	var links pkgTest.Links
	err := c.ShouldBindJSON(&links)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    links.DeletedBy = utils.GetUserID(c)
	if err := linksService.DeleteLinks(links); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLinksByIds 批量删除links表
// @Tags Links
// @Summary 批量删除links表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除links表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /links/deleteLinksByIds [delete]
func (linksApi *LinksApi) DeleteLinksByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := linksService.DeleteLinksByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLinks 更新links表
// @Tags Links
// @Summary 更新links表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest.Links true "更新links表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /links/updateLinks [put]
func (linksApi *LinksApi) UpdateLinks(c *gin.Context) {
	var links pkgTest.Links
	err := c.ShouldBindJSON(&links)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    links.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Url":{utils.NotEmpty()},
      }
    if err := utils.Verify(links, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := linksService.UpdateLinks(links); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLinks 用id查询links表
// @Tags Links
// @Summary 用id查询links表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTest.Links true "用id查询links表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /links/findLinks [get]
func (linksApi *LinksApi) FindLinks(c *gin.Context) {
	var links pkgTest.Links
	err := c.ShouldBindQuery(&links)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if relinks, err := linksService.GetLinks(links.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relinks": relinks}, c)
	}
}

// GetLinksList 分页获取links表列表
// @Tags Links
// @Summary 分页获取links表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTestReq.LinksSearch true "分页获取links表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /links/getLinksList [get]
func (linksApi *LinksApi) GetLinksList(c *gin.Context) {
	var pageInfo pkgTestReq.LinksSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := linksService.GetLinksInfoList(pageInfo); err != nil {
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
