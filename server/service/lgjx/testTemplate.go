package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/nonmigrate"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"io"
	"os"
)

type TestTemplateService struct {
}

func (testTemplateService *TestTemplateService) CreateTemplate(templateAndFile nonmigrate.TemplateAndFile) (err error) {
	//err = global.MustGetGlobalDBByDBName("lg-jx-test").Create(&template).Error
	basePath := "./tmp/"
	file, err := os.Open(basePath + *templateAndFile.FileName)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	templateAndFile.FileSteam = fileContent

	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		file := templateAndFile.File
		if err = tx.Create(&file).Error; err != nil {
			return err
		}
		template := templateAndFile.Template
		template.TemplateFileID = &file.ID
		if err = tx.Create(&template).Error; err != nil {
			return err
		}
		_ = os.Remove(basePath + *templateAndFile.FileName)
		return nil
	})

	return err
}

func (testTemplateService *TestTemplateService) DeleteTemplate(template lgjx.Template) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Select(clause.Associations).Delete(&template).Error
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

	err = db.Limit(limit).Preload(clause.Associations).Order("created_at desc").Offset(offset).Find(&templates).Error
	return templates, total, err
}
