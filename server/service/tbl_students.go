package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTblStudents
//@description: 创建TblStudents记录
//@param: students model.TblStudents
//@return: err error

func CreateTblStudents(students model.TblStudents) (err error) {
	err = global.GVA_DB.Create(&students).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTblStudents
//@description: 删除TblStudents记录
//@param: students model.TblStudents
//@return: err error

func DeleteTblStudents(students model.TblStudents) (err error) {
	err = global.GVA_DB.Delete(students).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTblStudentsByIds
//@description: 批量删除TblStudents记录
//@param: ids request.IdsReq
//@return: err error

func DeleteTblStudentsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.TblStudents{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTblStudents
//@description: 更新TblStudents记录
//@param: students *model.TblStudents
//@return: err error

func UpdateTblStudents(students *model.TblStudents) (err error) {
	err = global.GVA_DB.Save(students).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTblStudents
//@description: 根据id获取TblStudents记录
//@param: id uint
//@return: err error, students model.TblStudents

func GetTblStudents(id uint) (err error, students model.TblStudents) {
	err = global.GVA_DB.Where("id = ?", id).First(&students).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTblStudentsInfoList
//@description: 分页获取TblStudents记录
//@param: info request.TblStudentsSearch
//@return: err error, list interface{}, total int64

func GetTblStudentsInfoList(info request.TblStudentsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.TblStudents{})
	var studentss []model.TblStudents
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&studentss).Error
	return err, studentss, total
}
