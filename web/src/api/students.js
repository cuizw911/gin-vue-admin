import service from '@/utils/request'

// @Tags TblStudents
// @Summary 创建TblStudents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblStudents true "创建TblStudents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /students/createTblStudents [post]
export const createTblStudents = (data) => {
     return service({
         url: "/students/createTblStudents",
         method: 'post',
         data
     })
 }


// @Tags TblStudents
// @Summary 删除TblStudents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblStudents true "删除TblStudents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /students/deleteTblStudents [delete]
 export const deleteTblStudents = (data) => {
     return service({
         url: "/students/deleteTblStudents",
         method: 'delete',
         data
     })
 }

// @Tags TblStudents
// @Summary 删除TblStudents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TblStudents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /students/deleteTblStudents [delete]
 export const deleteTblStudentsByIds = (data) => {
     return service({
         url: "/students/deleteTblStudentsByIds",
         method: 'delete',
         data
     })
 }

// @Tags TblStudents
// @Summary 更新TblStudents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblStudents true "更新TblStudents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /students/updateTblStudents [put]
 export const updateTblStudents = (data) => {
     return service({
         url: "/students/updateTblStudents",
         method: 'put',
         data
     })
 }


// @Tags TblStudents
// @Summary 用id查询TblStudents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblStudents true "用id查询TblStudents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /students/findTblStudents [get]
 export const findTblStudents = (params) => {
     return service({
         url: "/students/findTblStudents",
         method: 'get',
         params
     })
 }


// @Tags TblStudents
// @Summary 分页获取TblStudents列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TblStudents列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /students/getTblStudentsList [get]
 export const getTblStudentsList = (params) => {
     return service({
         url: "/students/getTblStudentsList",
         method: 'get',
         params
     })
 }