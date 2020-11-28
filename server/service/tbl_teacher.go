package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTblTeacherInfo
//@description: 创建TblTeacherInfo记录
//@param: tblTeacherInfo model.TblTeacherInfo
//@return: err error

func CreateTblTeacherInfo(tblTeacherInfo model.TblTeacherInfo) (err error) {
	err = global.GVA_DB.Create(&tblTeacherInfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTblTeacherInfo
//@description: 删除TblTeacherInfo记录
//@param: tblTeacherInfo model.TblTeacherInfo
//@return: err error

func DeleteTblTeacherInfo(tblTeacherInfo model.TblTeacherInfo) (err error) {
	err = global.GVA_DB.Delete(tblTeacherInfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTblTeacherInfoByIds
//@description: 批量删除TblTeacherInfo记录
//@param: ids request.IdsReq
//@return: err error

func DeleteTblTeacherInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.TblTeacherInfo{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTblTeacherInfo
//@description: 更新TblTeacherInfo记录
//@param: tblTeacherInfo *model.TblTeacherInfo
//@return: err error

func UpdateTblTeacherInfo(tblTeacherInfo *model.TblTeacherInfo) (err error) {
	err = global.GVA_DB.Save(tblTeacherInfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTblTeacherInfo
//@description: 根据id获取TblTeacherInfo记录
//@param: id uint
//@return: err error, tblTeacherInfo model.TblTeacherInfo

func GetTblTeacherInfo(id uint) (err error, tblTeacherInfo model.TblTeacherInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&tblTeacherInfo).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTblTeacherInfoInfoList
//@description: 分页获取TblTeacherInfo记录
//@param: info request.TblTeacherInfoSearch
//@return: err error, list interface{}, total int64

func GetTblTeacherInfoInfoList(info request.TblTeacherInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.TblTeacherInfo{})
    var tblTeacherInfos []model.TblTeacherInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if info.Grade != "" {
        db = db.Where("`grade` = ?",info.Grade)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&tblTeacherInfos).Error
	return err, tblTeacherInfos, total
}