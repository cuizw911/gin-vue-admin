package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTblCourseInfo
//@description: 创建TblCourseInfo记录
//@param: tblCourseInfo model.TblCourseInfo
//@return: err error

func CreateTblCourseInfo(tblCourseInfo model.TblCourseInfo) (err error) {
	err = global.GVA_DB.Create(&tblCourseInfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTblCourseInfo
//@description: 删除TblCourseInfo记录
//@param: tblCourseInfo model.TblCourseInfo
//@return: err error

func DeleteTblCourseInfo(tblCourseInfo model.TblCourseInfo) (err error) {
	err = global.GVA_DB.Delete(tblCourseInfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTblCourseInfoByIds
//@description: 批量删除TblCourseInfo记录
//@param: ids request.IdsReq
//@return: err error

func DeleteTblCourseInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.TblCourseInfo{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTblCourseInfo
//@description: 更新TblCourseInfo记录
//@param: tblCourseInfo *model.TblCourseInfo
//@return: err error

func UpdateTblCourseInfo(tblCourseInfo *model.TblCourseInfo) (err error) {
	err = global.GVA_DB.Save(tblCourseInfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTblCourseInfo
//@description: 根据id获取TblCourseInfo记录
//@param: id uint
//@return: err error, tblCourseInfo model.TblCourseInfo

func GetTblCourseInfo(id uint) (err error, tblCourseInfo model.TblCourseInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&tblCourseInfo).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTblCourseInfoInfoList
//@description: 分页获取TblCourseInfo记录
//@param: info request.TblCourseInfoSearch
//@return: err error, list interface{}, total int64

func GetTblCourseInfoInfoList(info request.TblCourseInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.TblCourseInfo{})
    var tblCourseInfos []model.TblCourseInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&tblCourseInfos).Error
	return err, tblCourseInfos, total
}