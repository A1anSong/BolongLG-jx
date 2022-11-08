package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type LetterService struct {
}

// CreateLetter 创建Letter记录
// Author [piexlmax](https://github.com/piexlmax)
func (letterService *LetterService) CreateLetter(letter lgjx.Letter) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Create(&letter).Error
	return err
}

// DeleteLetter 删除Letter记录
// Author [piexlmax](https://github.com/piexlmax)
func (letterService *LetterService) DeleteLetter(letter lgjx.Letter) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&letter).Error
	return err
}

// DeleteLetterByIds 批量删除Letter记录
// Author [piexlmax](https://github.com/piexlmax)
func (letterService *LetterService) DeleteLetterByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&[]lgjx.Letter{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateLetter 更新Letter记录
// Author [piexlmax](https://github.com/piexlmax)
func (letterService *LetterService) UpdateLetter(letter lgjx.Letter) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Save(&letter).Error
	return err
}

// GetLetter 根据id获取Letter记录
// Author [piexlmax](https://github.com/piexlmax)
func (letterService *LetterService) GetLetter(id uint) (letter lgjx.Letter, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Where("id = ?", id).First(&letter).Error
	return
}

// GetLetterInfoList 分页获取Letter记录
// Author [piexlmax](https://github.com/piexlmax)
func (letterService *LetterService) GetLetterInfoList(info lgjxReq.LetterSearch) (list []lgjx.Letter, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx").Model(&lgjx.Letter{})
	var letters []lgjx.Letter
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&letters).Error
	return letters, total, err
}
