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

type CategoriesApi struct {
}

var categoriesService = service.ServiceGroupApp.PkgTestServiceGroup.CategoriesService


// CreateCategories 创建categories表
// @Tags Categories
// @Summary 创建categories表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest.Categories true "创建categories表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /categories/createCategories [post]
func (categoriesApi *CategoriesApi) CreateCategories(c *gin.Context) {
	var categories pkgTest.Categories
	err := c.ShouldBindJSON(&categories)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    categories.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
    }
	if err := utils.Verify(categories, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := categoriesService.CreateCategories(&categories); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCategories 删除categories表
// @Tags Categories
// @Summary 删除categories表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest.Categories true "删除categories表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /categories/deleteCategories [delete]
func (categoriesApi *CategoriesApi) DeleteCategories(c *gin.Context) {
	var categories pkgTest.Categories
	err := c.ShouldBindJSON(&categories)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    categories.DeletedBy = utils.GetUserID(c)
	if err := categoriesService.DeleteCategories(categories); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCategoriesByIds 批量删除categories表
// @Tags Categories
// @Summary 批量删除categories表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除categories表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /categories/deleteCategoriesByIds [delete]
func (categoriesApi *CategoriesApi) DeleteCategoriesByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := categoriesService.DeleteCategoriesByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCategories 更新categories表
// @Tags Categories
// @Summary 更新categories表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest.Categories true "更新categories表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /categories/updateCategories [put]
func (categoriesApi *CategoriesApi) UpdateCategories(c *gin.Context) {
	var categories pkgTest.Categories
	err := c.ShouldBindJSON(&categories)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    categories.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
      }
    if err := utils.Verify(categories, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := categoriesService.UpdateCategories(categories); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCategories 用id查询categories表
// @Tags Categories
// @Summary 用id查询categories表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTest.Categories true "用id查询categories表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /categories/findCategories [get]
func (categoriesApi *CategoriesApi) FindCategories(c *gin.Context) {
	var categories pkgTest.Categories
	err := c.ShouldBindQuery(&categories)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recategories, err := categoriesService.GetCategories(categories.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recategories": recategories}, c)
	}
}

// GetCategoriesList 分页获取categories表列表
// @Tags Categories
// @Summary 分页获取categories表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTestReq.CategoriesSearch true "分页获取categories表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /categories/getCategoriesList [get]
func (categoriesApi *CategoriesApi) GetCategoriesList(c *gin.Context) {
	var pageInfo pkgTestReq.CategoriesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := categoriesService.GetCategoriesInfoList(pageInfo); err != nil {
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
