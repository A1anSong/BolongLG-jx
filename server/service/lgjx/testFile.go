package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type TestFileService struct {
}

func (testFileService *TestFileService) CreateFile(file lgjx.File) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Create(&file).Error
	return err
}

func (testFileService *TestFileService) DeleteFile(file lgjx.File) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&file).Error
	return err
}

func (testFileService *TestFileService) DeleteFileByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&[]lgjx.File{}, "id in ?", ids.Ids).Error
	return err
}

func (testFileService *TestFileService) UpdateFile(file lgjx.File) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Save(&file).Error
	return err
}

func (testFileService *TestFileService) GetFile(id uint) (file lgjx.File, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Where("id = ?", id).First(&file).Error
	return
}

func (testFileService *TestFileService) GetFileInfoList(info lgjxReq.FileSearch) (list []lgjx.File, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx-test").Model(&lgjx.File{})
	var files []lgjx.File
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&files).Error
	return files, total, err
}

func (testFileService *TestFileService) UploadFile(file *multipart.FileHeader) (fileName string, err error) {
	basePath := "./tmp/"
	fileSuffix := path.Ext(file.Filename)
	fileNameOnly := strings.TrimSuffix(file.Filename, fileSuffix)
	year, month, day := time.Now().Date()
	fileName = fileNameOnly + strconv.Itoa(year) + strconv.Itoa(int(month)) + strconv.Itoa(day) + fileSuffix
	out, err := os.Create(basePath + fileName)
	if err != nil {
		return "", err
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)
	fileOut, err := file.Open()
	_, err = io.Copy(out, fileOut)
	if err != nil {
		return "", err
	}
	return fileName, nil
}
