package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type TestTemplateService struct {
}

func (testTemplateService *TestTemplateService) CreateTemplate(template lgjx.Template) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Create(&template).Error
	return err
}

func (testTemplateService *TestTemplateService) DeleteTemplate(template lgjx.Template) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&template).Error
	return err
}

func (testTemplateService *TestTemplateService) DeleteTemplateByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&[]lgjx.Template{}, "id in ?", ids.Ids).Error
	return err
}

func (testTemplateService *TestTemplateService) UpdateTemplate(template lgjx.Template) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Save(&template).Error
	return err
}

func (testTemplateService *TestTemplateService) GetTemplate(id uint) (template lgjx.Template, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Where("id = ?", id).First(&template).Error
	return
}

func (testTemplateService *TestTemplateService) GetTemplateInfoList(info lgjxReq.TemplateSearch) (list []lgjx.Template, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx-test").Model(&lgjx.Template{})
	var templates []lgjx.Template
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&templates).Error
	return templates, total, err
}
