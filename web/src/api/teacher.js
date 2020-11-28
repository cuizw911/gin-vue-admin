import service from '@/utils/request'

// @Tags TblTeacherInfo
// @Summary 创建TblTeacherInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblTeacherInfo true "创建TblTeacherInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tblTeacherInfo/createTblTeacherInfo [post]
export const createTblTeacherInfo = (data) => {
     return service({
         url: "/tblTeacherInfo/createTblTeacherInfo",
         method: 'post',
         data
     })
 }


// @Tags TblTeacherInfo
// @Summary 删除TblTeacherInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblTeacherInfo true "删除TblTeacherInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tblTeacherInfo/deleteTblTeacherInfo [delete]
 export const deleteTblTeacherInfo = (data) => {
     return service({
         url: "/tblTeacherInfo/deleteTblTeacherInfo",
         method: 'delete',
         data
     })
 }

// @Tags TblTeacherInfo
// @Summary 删除TblTeacherInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TblTeacherInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tblTeacherInfo/deleteTblTeacherInfo [delete]
 export const deleteTblTeacherInfoByIds = (data) => {
     return service({
         url: "/tblTeacherInfo/deleteTblTeacherInfoByIds",
         method: 'delete',
         data
     })
 }

// @Tags TblTeacherInfo
// @Summary 更新TblTeacherInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblTeacherInfo true "更新TblTeacherInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tblTeacherInfo/updateTblTeacherInfo [put]
 export const updateTblTeacherInfo = (data) => {
     return service({
         url: "/tblTeacherInfo/updateTblTeacherInfo",
         method: 'put',
         data
     })
 }


// @Tags TblTeacherInfo
// @Summary 用id查询TblTeacherInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblTeacherInfo true "用id查询TblTeacherInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tblTeacherInfo/findTblTeacherInfo [get]
 export const findTblTeacherInfo = (params) => {
     return service({
         url: "/tblTeacherInfo/findTblTeacherInfo",
         method: 'get',
         params
     })
 }


// @Tags TblTeacherInfo
// @Summary 分页获取TblTeacherInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TblTeacherInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tblTeacherInfo/getTblTeacherInfoList [get]
 export const getTblTeacherInfoList = (params) => {
     return service({
         url: "/tblTeacherInfo/getTblTeacherInfoList",
         method: 'get',
         params
     })
 }