import service from '@/utils/request'

// @Tags Categories
// @Summary 创建categories表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Categories true "创建categories表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /categories/createCategories [post]
export const createCategories = (data) => {
  return service({
    url: '/categories/createCategories',
    method: 'post',
    data
  })
}

// @Tags Categories
// @Summary 删除categories表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Categories true "删除categories表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /categories/deleteCategories [delete]
export const deleteCategories = (data) => {
  return service({
    url: '/categories/deleteCategories',
    method: 'delete',
    data
  })
}

// @Tags Categories
// @Summary 批量删除categories表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除categories表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /categories/deleteCategories [delete]
export const deleteCategoriesByIds = (data) => {
  return service({
    url: '/categories/deleteCategoriesByIds',
    method: 'delete',
    data
  })
}

// @Tags Categories
// @Summary 更新categories表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Categories true "更新categories表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /categories/updateCategories [put]
export const updateCategories = (data) => {
  return service({
    url: '/categories/updateCategories',
    method: 'put',
    data
  })
}

// @Tags Categories
// @Summary 用id查询categories表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Categories true "用id查询categories表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /categories/findCategories [get]
export const findCategories = (params) => {
  return service({
    url: '/categories/findCategories',
    method: 'get',
    params
  })
}

// @Tags Categories
// @Summary 分页获取categories表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取categories表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /categories/getCategoriesList [get]
export const getCategoriesList = (params) => {
  return service({
    url: '/categories/getCategoriesList',
    method: 'get',
    params
  })
}
