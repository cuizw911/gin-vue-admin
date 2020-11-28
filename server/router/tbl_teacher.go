package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTblTeacherInfoRouter(Router *gin.RouterGroup) {
	TblTeacherInfoRouter := Router.Group("tblTeacherInfo").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		TblTeacherInfoRouter.POST("createTblTeacherInfo", v1.CreateTblTeacherInfo)   // 新建TblTeacherInfo
		TblTeacherInfoRouter.DELETE("deleteTblTeacherInfo", v1.DeleteTblTeacherInfo) // 删除TblTeacherInfo
		TblTeacherInfoRouter.DELETE("deleteTblTeacherInfoByIds", v1.DeleteTblTeacherInfoByIds) // 批量删除TblTeacherInfo
		TblTeacherInfoRouter.PUT("updateTblTeacherInfo", v1.UpdateTblTeacherInfo)    // 更新TblTeacherInfo
		TblTeacherInfoRouter.GET("findTblTeacherInfo", v1.FindTblTeacherInfo)        // 根据ID获取TblTeacherInfo
		TblTeacherInfoRouter.GET("getTblTeacherInfoList", v1.GetTblTeacherInfoList)  // 获取TblTeacherInfo列表
	}
}
