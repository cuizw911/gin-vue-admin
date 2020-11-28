import service from '@/utils/request'

// @Tags TblCourseInfo
// @Summary 创建TblCourseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblCourseInfo true "创建TblCourseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tblCourseInfo/createTblCourseInfo [post]
export const createTblCourseInfo = (data) => {
     return service({
         url: "/tblCourseInfo/createTblCourseInfo",
         method: 'post',
         data
     })
 }


// @Tags TblCourseInfo
// @Summary 删除TblCourseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblCourseInfo true "删除TblCourseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tblCourseInfo/deleteTblCourseInfo [delete]
 export const deleteTblCourseInfo = (data) => {
     return service({
         url: "/tblCourseInfo/deleteTblCourseInfo",
         method: 'delete',
         data
     })
 }

// @Tags TblCourseInfo
// @Summary 删除TblCourseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TblCourseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tblCourseInfo/deleteTblCourseInfo [delete]
 export const deleteTblCourseInfoByIds = (data) => {
     return service({
         url: "/tblCourseInfo/deleteTblCourseInfoByIds",
         method: 'delete',
         data
     })
 }

// @Tags TblCourseInfo
// @Summary 更新TblCourseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblCourseInfo true "更新TblCourseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tblCourseInfo/updateTblCourseInfo [put]
 export const updateTblCourseInfo = (data) => {
     return service({
         url: "/tblCourseInfo/updateTblCourseInfo",
         method: 'put',
         data
     })
 }


// @Tags TblCourseInfo
// @Summary 用id查询TblCourseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblCourseInfo true "用id查询TblCourseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tblCourseInfo/findTblCourseInfo [get]
 export const findTblCourseInfo = (params) => {
     return service({
         url: "/tblCourseInfo/findTblCourseInfo",
         method: 'get',
         params
     })
 }


// @Tags TblCourseInfo
// @Summary 分页获取TblCourseInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TblCourseInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tblCourseInfo/getTblCourseInfoList [get]
 export const getTblCourseInfoList = (params) => {
     return service({
         url: "/tblCourseInfo/getTblCourseInfoList",
         method: 'get',
         params
     })
 }