package lgjx

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestProjectRouter struct {
}

func (s *TestProjectRouter) InitProjectRouter(Router *gin.RouterGroup) {
	testProjectRouter := Router.Group("testProject").Use(middleware.OperationRecord())
	testProjectRouterWithoutRecord := Router.Group("testProject")
	var testProjectApi = v1.ApiGroupApp.LgjxTestApiGroup.TestProjectApi
	{
		testProjectRouter.POST("createProject", testProjectApi.CreateProject)             // 新建Project
		testProjectRouter.DELETE("deleteProject", testProjectApi.DeleteProject)           // 删除Project
		testProjectRouter.DELETE("deleteProjectByIds", testProjectApi.DeleteProjectByIds) // 批量删除Project
		testProjectRouter.PUT("updateProject", testProjectApi.UpdateProject)              // 更新Project
	}
	{
		testProjectRouterWithoutRecord.GET("findProject", testProjectApi.FindProject)       // 根据ID获取Project
		testProjectRouterWithoutRecord.GET("getProjectList", testProjectApi.GetProjectList) // 获取Project列表
		testProjectRouterWithoutRecord.POST("bindProject", testProjectApi.BindProject)      // 绑定项目
		testProjectRouterWithoutRecord.POST("unbindProject", testProjectApi.UnbindProject)  // 解绑项目
	}
}
