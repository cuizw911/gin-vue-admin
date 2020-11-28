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

// @Tags TblCourseInfo
// @Summary 创建TblCourseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblCourseInfo true "创建TblCourseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tblCourseInfo/createTblCourseInfo [post]
func CreateTblCourseInfo(c *gin.Context) {
	var tblCourseInfo model.TblCourseInfo
	_ = c.ShouldBindJSON(&tblCourseInfo)
	if err := service.CreateTblCourseInfo(tblCourseInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TblCourseInfo
// @Summary 删除TblCourseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblCourseInfo true "删除TblCourseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tblCourseInfo/deleteTblCourseInfo [delete]
func DeleteTblCourseInfo(c *gin.Context) {
	var tblCourseInfo model.TblCourseInfo
	_ = c.ShouldBindJSON(&tblCourseInfo)
	if err := service.DeleteTblCourseInfo(tblCourseInfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TblCourseInfo
// @Summary 批量删除TblCourseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TblCourseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tblCourseInfo/deleteTblCourseInfoByIds [delete]
func DeleteTblCourseInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteTblCourseInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags TblCourseInfo
// @Summary 更新TblCourseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblCourseInfo true "更新TblCourseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tblCourseInfo/updateTblCourseInfo [put]
func UpdateTblCourseInfo(c *gin.Context) {
	var tblCourseInfo model.TblCourseInfo
	_ = c.ShouldBindJSON(&tblCourseInfo)
	if err := service.UpdateTblCourseInfo(&tblCourseInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TblCourseInfo
// @Summary 用id查询TblCourseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblCourseInfo true "用id查询TblCourseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tblCourseInfo/findTblCourseInfo [get]
func FindTblCourseInfo(c *gin.Context) {
	var tblCourseInfo model.TblCourseInfo
	_ = c.ShouldBindQuery(&tblCourseInfo)
	if err, retblCourseInfo := service.GetTblCourseInfo(tblCourseInfo.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retblCourseInfo": retblCourseInfo}, c)
	}
}

// @Tags TblCourseInfo
// @Summary 分页获取TblCourseInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TblCourseInfoSearch true "分页获取TblCourseInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tblCourseInfo/getTblCourseInfoList [get]
func GetTblCourseInfoList(c *gin.Context) {
	var pageInfo request.TblCourseInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTblCourseInfoInfoList(pageInfo); err != nil {
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
