package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type TestLetterService struct {
}

func (testLetterService *TestLetterService) CreateLetter(letter lgjx.Letter) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Create(&letter).Error
	return err
}

func (testLetterService *TestLetterService) DeleteLetter(letter lgjx.Letter) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&letter).Error
	return err
}

func (testLetterService *TestLetterService) DeleteLetterByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&[]lgjx.Letter{}, "id in ?", ids.Ids).Error
	return err
}

func (testLetterService *TestLetterService) UpdateLetter(letter lgjx.Letter) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Save(&letter).Error
	return err
}

func (testLetterService *TestLetterService) GetLetter(id uint) (letter lgjx.Letter, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Where("id = ?", id).First(&letter).Error
	return
}

func (testLetterService *TestLetterService) GetLetterInfoList(info lgjxReq.LetterSearch) (list []lgjx.Letter, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx-test").Model(&lgjx.Letter{})
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
