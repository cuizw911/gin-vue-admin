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

// @Tags TblTeacherInfo
// @Summary 创建TblTeacherInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblTeacherInfo true "创建TblTeacherInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tblTeacherInfo/createTblTeacherInfo [post]
func CreateTblTeacherInfo(c *gin.Context) {
	var tblTeacherInfo model.TblTeacherInfo
	_ = c.ShouldBindJSON(&tblTeacherInfo)
	if err := service.CreateTblTeacherInfo(tblTeacherInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TblTeacherInfo
// @Summary 删除TblTeacherInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblTeacherInfo true "删除TblTeacherInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tblTeacherInfo/deleteTblTeacherInfo [delete]
func DeleteTblTeacherInfo(c *gin.Context) {
	var tblTeacherInfo model.TblTeacherInfo
	_ = c.ShouldBindJSON(&tblTeacherInfo)
	if err := service.DeleteTblTeacherInfo(tblTeacherInfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TblTeacherInfo
// @Summary 批量删除TblTeacherInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TblTeacherInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tblTeacherInfo/deleteTblTeacherInfoByIds [delete]
func DeleteTblTeacherInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteTblTeacherInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags TblTeacherInfo
// @Summary 更新TblTeacherInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblTeacherInfo true "更新TblTeacherInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tblTeacherInfo/updateTblTeacherInfo [put]
func UpdateTblTeacherInfo(c *gin.Context) {
	var tblTeacherInfo model.TblTeacherInfo
	_ = c.ShouldBindJSON(&tblTeacherInfo)
	if err := service.UpdateTblTeacherInfo(&tblTeacherInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TblTeacherInfo
// @Summary 用id查询TblTeacherInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TblTeacherInfo true "用id查询TblTeacherInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tblTeacherInfo/findTblTeacherInfo [get]
func FindTblTeacherInfo(c *gin.Context) {
	var tblTeacherInfo model.TblTeacherInfo
	_ = c.ShouldBindQuery(&tblTeacherInfo)
	if err, retblTeacherInfo := service.GetTblTeacherInfo(tblTeacherInfo.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retblTeacherInfo": retblTeacherInfo}, c)
	}
}

// @Tags TblTeacherInfo
// @Summary 分页获取TblTeacherInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TblTeacherInfoSearch true "分页获取TblTeacherInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tblTeacherInfo/getTblTeacherInfoList [get]
func GetTblTeacherInfoList(c *gin.Context) {
	var pageInfo request.TblTeacherInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTblTeacherInfoInfoList(pageInfo); err != nil {
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
