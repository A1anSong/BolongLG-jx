package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type TemplateService struct {
}

// CreateTemplate 创建Template记录
// Author [piexlmax](https://github.com/piexlmax)
func (templateService *TemplateService) CreateTemplate(template lgjx.Template) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Create(&template).Error
	return err
}

// DeleteTemplate 删除Template记录
// Author [piexlmax](https://github.com/piexlmax)
func (templateService *TemplateService) DeleteTemplate(template lgjx.Template) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&template).Error
	return err
}

// DeleteTemplateByIds 批量删除Template记录
// Author [piexlmax](https://github.com/piexlmax)
func (templateService *TemplateService) DeleteTemplateByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&[]lgjx.Template{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTemplate 更新Template记录
// Author [piexlmax](https://github.com/piexlmax)
func (templateService *TemplateService) UpdateTemplate(template lgjx.Template) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Save(&template).Error
	return err
}

// GetTemplate 根据id获取Template记录
// Author [piexlmax](https://github.com/piexlmax)
func (templateService *TemplateService) GetTemplate(id uint) (template lgjx.Template, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Where("id = ?", id).First(&template).Error
	return
}

// GetTemplateInfoList 分页获取Template记录
// Author [piexlmax](https://github.com/piexlmax)
func (templateService *TemplateService) GetTemplateInfoList(info lgjxReq.TemplateSearch) (list []lgjx.Template, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx").Model(&lgjx.Template{})
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
