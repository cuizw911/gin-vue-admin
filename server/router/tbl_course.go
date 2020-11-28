package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTblCourseInfoRouter(Router *gin.RouterGroup) {
	TblCourseInfoRouter := Router.Group("tblCourseInfo").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		TblCourseInfoRouter.POST("createTblCourseInfo", v1.CreateTblCourseInfo)   // 新建TblCourseInfo
		TblCourseInfoRouter.DELETE("deleteTblCourseInfo", v1.DeleteTblCourseInfo) // 删除TblCourseInfo
		TblCourseInfoRouter.DELETE("deleteTblCourseInfoByIds", v1.DeleteTblCourseInfoByIds) // 批量删除TblCourseInfo
		TblCourseInfoRouter.PUT("updateTblCourseInfo", v1.UpdateTblCourseInfo)    // 更新TblCourseInfo
		TblCourseInfoRouter.GET("findTblCourseInfo", v1.FindTblCourseInfo)        // 根据ID获取TblCourseInfo
		TblCourseInfoRouter.GET("getTblCourseInfoList", v1.GetTblCourseInfoList)  // 获取TblCourseInfo列表
	}
}
