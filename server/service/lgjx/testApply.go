package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type TestApplyService struct {
}

func (testApplyService *TestApplyService) CreateApply(apply lgjx.Apply) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Create(&apply).Error
	return err
}

func (testApplyService *TestApplyService) DeleteApply(apply lgjx.Apply) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&apply).Error
	return err
}

func (testApplyService *TestApplyService) DeleteApplyByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&[]lgjx.Apply{}, "id in ?", ids.Ids).Error
	return err
}

func (testApplyService *TestApplyService) UpdateApply(apply lgjx.Apply) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Save(&apply).Error
	return err
}

func (testApplyService *TestApplyService) GetApply(id uint) (apply lgjx.Apply, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Where("id = ?", id).First(&apply).Error
	return
}

func (testApplyService *TestApplyService) GetApplyInfoList(info lgjxReq.ApplySearch) (list []lgjx.Apply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx-test").Model(&lgjx.Apply{})
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
