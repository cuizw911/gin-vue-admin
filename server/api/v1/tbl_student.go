package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags TblStudents
// @Summary 创建TblStudents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblStudents true "创建TblStudents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /students/createTblStudents [post]
func CreateTblStudents(c *gin.Context) {
	var students model.TblStudents
	_ = c.ShouldBindJSON(&students)
	if err := service.CreateTblStudents(students); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TblStudents
// @Summary 删除TblStudents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblStudents true "删除TblStudents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /students/deleteTblStudents [delete]
func DeleteTblStudents(c *gin.Context) {
	var students model.TblStudents
	_ = c.ShouldBindJSON(&students)
	if err := service.DeleteTblStudents(students); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TblStudents
// @Summary 批量删除TblStudents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TblStudents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /students/deleteTblStudentsByIds [delete]
func DeleteTblStudentsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteTblStudentsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags TblStudents
// @Summary 更新TblStudents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblStudents true "更新TblStudents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /students/updateTblStudents [put]
func UpdateTblStudents(c *gin.Context) {
	var students model.TblStudents
	_ = c.ShouldBindJSON(&students)
	if err := service.UpdateTblStudents(&students); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TblStudents
// @Summary 用id查询TblStudents
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblStudents true "用id查询TblStudents"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /students/findTblStudents [get]
func FindTblStudents(c *gin.Context) {
	var students model.TblStudents
	_ = c.ShouldBindQuery(&students)
	if err, restudents := service.GetTblStudents(students.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restudents": restudents}, c)
	}
}

// @Tags TblStudents
// @Summary 分页获取TblStudents列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TblStudentsSearch true "分页获取TblStudents列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /students/getTblStudentsList [get]
func GetTblStudentsList(c *gin.Context) {
	var pageInfo request.TblStudentsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTblStudentsInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
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
