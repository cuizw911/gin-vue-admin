package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTblStudentsRouter(Router *gin.RouterGroup) {
	TblStudentsRouter := Router.Group("students").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		TblStudentsRouter.POST("createTblStudents", v1.CreateTblStudents)             // 新建TblStudents
		TblStudentsRouter.DELETE("deleteTblStudents", v1.DeleteTblStudents)           // 删除TblStudents
		TblStudentsRouter.DELETE("deleteTblStudentsByIds", v1.DeleteTblStudentsByIds) // 批量删除TblStudents
		TblStudentsRouter.PUT("updateTblStudents", v1.UpdateTblStudents)              // 更新TblStudents
		TblStudentsRouter.GET("findTblStudents", v1.FindTblStudents)                  // 根据ID获取TblStudents
		TblStudentsRouter.GET("getTblStudentsList", v1.GetTblStudentsList)            // 获取TblStudents列表
	}
}
