package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type ApplyService struct {
}

// CreateApply 创建Apply记录
// Author [piexlmax](https://github.com/piexlmax)
func (applyService *ApplyService) CreateApply(apply lgjx.Apply) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Create(&apply).Error
	return err
}

// DeleteApply 删除Apply记录
// Author [piexlmax](https://github.com/piexlmax)
func (applyService *ApplyService) DeleteApply(apply lgjx.Apply) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&apply).Error
	return err
}

// DeleteApplyByIds 批量删除Apply记录
// Author [piexlmax](https://github.com/piexlmax)
func (applyService *ApplyService) DeleteApplyByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&[]lgjx.Apply{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateApply 更新Apply记录
// Author [piexlmax](https://github.com/piexlmax)
func (applyService *ApplyService) UpdateApply(apply lgjx.Apply) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Save(&apply).Error
	return err
}

// GetApply 根据id获取Apply记录
// Author [piexlmax](https://github.com/piexlmax)
func (applyService *ApplyService) GetApply(id uint) (apply lgjx.Apply, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Where("id = ?", id).First(&apply).Error
	return
}

// GetApplyInfoList 分页获取Apply记录
// Author [piexlmax](https://github.com/piexlmax)
func (applyService *ApplyService) GetApplyInfoList(info lgjxReq.ApplySearch) (list []lgjx.Apply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx").Model(&lgjx.Apply{})
	var applys []lgjx.Apply
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&applys).Error
	return applys, total, err
}
