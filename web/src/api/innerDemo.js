import service from '@/utils/request'

// @Tags InnerDemo
// @Summary 创建内部测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InnerDemo true "创建内部测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /demo/createInnerDemo [post]
export const createInnerDemo = (data) => {
  return service({
    url: '/demo/createInnerDemo',
    method: 'post',
    data
  })
}

// @Tags InnerDemo
// @Summary 删除内部测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InnerDemo true "删除内部测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /demo/deleteInnerDemo [delete]
export const deleteInnerDemo = (data) => {
  return service({
    url: '/demo/deleteInnerDemo',
    method: 'delete',
    data
  })
}

// @Tags InnerDemo
// @Summary 批量删除内部测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除内部测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /demo/deleteInnerDemo [delete]
export const deleteInnerDemoByIds = (data) => {
  return service({
    url: '/demo/deleteInnerDemoByIds',
    method: 'delete',
    data
  })
}

// @Tags InnerDemo
// @Summary 更新内部测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InnerDemo true "更新内部测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /demo/updateInnerDemo [put]
export const updateInnerDemo = (data) => {
  return service({
    url: '/demo/updateInnerDemo',
    method: 'put',
    data
  })
}

// @Tags InnerDemo
// @Summary 用id查询内部测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.InnerDemo true "用id查询内部测试"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /demo/findInnerDemo [get]
export const findInnerDemo = (params) => {
  return service({
    url: '/demo/findInnerDemo',
    method: 'get',
    params
  })
}

// @Tags InnerDemo
// @Summary 分页获取内部测试列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取内部测试列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /demo/getInnerDemoList [get]
export const getInnerDemoList = (params) => {
  return service({
    url: '/demo/getInnerDemoList',
    method: 'get',
    params
  })
}
