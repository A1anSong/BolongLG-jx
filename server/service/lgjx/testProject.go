package lgjx

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
	"gorm.io/gorm"
)

type TestProjectService struct {
}

func (testProjectService *TestProjectService) CreateProject(project lgjx.Project) (err error) {
	fmt.Println(project.ID)
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Create(&project).Error
	fmt.Println(project.ID)
	return err
}

func (testProjectService *TestProjectService) DeleteProject(project lgjx.Project) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&project).Error
	return err
}

func (testProjectService *TestProjectService) DeleteProjectByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&[]lgjx.Project{}, "id in ?", ids.Ids).Error
	return err
}

func (testProjectService *TestProjectService) UpdateProject(project lgjx.Project) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Save(&project).Error
	return err
}

func (testProjectService *TestProjectService) GetProject(id uint) (project lgjx.Project, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Where("id = ?", id).First(&project).Error
	return
}

func (testProjectService *TestProjectService) GetProjectInfoList(info lgjxReq.ProjectSearch) (list []lgjx.Project, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx-test").Model(&lgjx.Project{})
	var projects []lgjx.Project
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&projects).Error
	return projects, total, err
}

func (testProjectService *TestProjectService) BindProject(project lgjx.Project) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		var orders []lgjx.Order
		err = tx.Model(&lgjx.Order{}).Joins("Apply").Where("Apply.project_no = ?", project.ProjectNo).Find(&orders).Error
		if err != nil {
			return err
		}
		for i := range orders {
			orders[i].ProjectID = &project.ID
		}
		err = tx.Save(&orders).Error
		if err != nil {
			return err
		}
		err = tx.Save(&project).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (testProjectService *TestProjectService) UnbindProject(project lgjx.Project) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		var orders []lgjx.Order
		err = tx.Model(&lgjx.Order{}).Joins("Apply").Where("Apply.project_no = ?", project.ProjectNo).Find(&orders).Error
		if err != nil {
			return err
		}
		for i := range orders {
			orders[i].ProjectID = nil
		}
		err = tx.Save(&orders).Error
		if err != nil {
			return err
		}
		err = tx.Save(&project).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
