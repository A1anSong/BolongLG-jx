package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type ProjectService struct {
}

// CreateProject 创建Project记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectService) CreateProject(project lgjx.Project) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Create(&project).Error
	return err
}

// DeleteProject 删除Project记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectService) DeleteProject(project lgjx.Project) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&project).Error
	return err
}

// DeleteProjectByIds 批量删除Project记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectService) DeleteProjectByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&[]lgjx.Project{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateProject 更新Project记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectService) UpdateProject(project lgjx.Project) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Save(&project).Error
	return err
}

// GetProject 根据id获取Project记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectService) GetProject(id uint) (project lgjx.Project, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Where("id = ?", id).First(&project).Error
	return
}

// GetProjectInfoList 分页获取Project记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectService) GetProjectInfoList(info lgjxReq.ProjectSearch) (list []lgjx.Project, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx").Model(&lgjx.Project{})
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
