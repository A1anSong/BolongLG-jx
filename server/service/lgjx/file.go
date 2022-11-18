package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type FileService struct {
}

// CreateFile 创建File记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileService *FileService) CreateFile(file lgjx.File) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Create(&file).Error
	return err
}

// DeleteFile 删除File记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileService *FileService) DeleteFile(file lgjx.File) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&file).Error
	return err
}

// DeleteFileByIds 批量删除File记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileService *FileService) DeleteFileByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&[]lgjx.File{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateFile 更新File记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileService *FileService) UpdateFile(file lgjx.File) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Save(&file).Error
	return err
}

// GetFile 根据id获取File记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileService *FileService) GetFile(id uint) (file lgjx.File, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Where("id = ?", id).First(&file).Error
	return
}

// GetFileInfoList 分页获取File记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileService *FileService) GetFileInfoList(info lgjxReq.FileSearch) (list []lgjx.File, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx").Model(&lgjx.File{})
	var files []lgjx.File
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&files).Error
	return files, total, err
}
