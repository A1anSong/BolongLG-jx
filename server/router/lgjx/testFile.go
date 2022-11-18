package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestFileRouter struct {
}

func (s *TestFileRouter) InitFileRouter(Router *gin.RouterGroup) {
	testFileRouter := Router.Group("testFile").Use(middleware.OperationRecord())
	testFileRouterWithoutRecord := Router.Group("testFile")
	var testFileApi = v1.ApiGroupApp.LgjxTestApiGroup.TestFileApi
	{
		testFileRouter.POST("createFile", testFileApi.CreateFile)             // 新建File
		testFileRouter.DELETE("deleteFile", testFileApi.DeleteFile)           // 删除File
		testFileRouter.DELETE("deleteFileByIds", testFileApi.DeleteFileByIds) // 批量删除File
		testFileRouter.PUT("updateFile", testFileApi.UpdateFile)              // 更新File
	}
	{
		testFileRouterWithoutRecord.GET("findFile", testFileApi.FindFile)       // 根据ID获取File
		testFileRouterWithoutRecord.GET("getFileList", testFileApi.GetFileList) // 获取File列表
		testFileRouterWithoutRecord.POST("upload", testFileApi.UploadFile)      // 上传文件
	}
}
