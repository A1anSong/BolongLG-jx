import service from '@/utils/request'
import { ElMessage } from 'element-plus'

export const createFile = (data) => {
  return service({
    url: '/testFile/createFile',
    method: 'post',
    data
  })
}

export const deleteFile = (data) => {
  return service({
    url: '/testFile/deleteFile',
    method: 'delete',
    data
  })
}

export const deleteFileByIds = (data) => {
  return service({
    url: '/testFile/deleteFileByIds',
    method: 'delete',
    data
  })
}

export const updateFile = (data) => {
  return service({
    url: '/testFile/updateFile',
    method: 'put',
    data
  })
}

export const findFile = (params) => {
  return service({
    url: '/testFile/findFile',
    method: 'get',
    params
  })
}

export const getFileList = (params) => {
  return service({
    url: '/testFile/getFileList',
    method: 'get',
    params
  })
}

export const downloadFile = (file) => {
  return service({
    url: '/testFile/download',
    method: 'get',
    params: file,
    responseType: 'blob'
  }).then((res) => {
    handleFileError(res, file.fileName)
  })
}

const handleFileError = (res, fileName) => {
  if (typeof (res.data) !== 'undefined') {
    if (res.data.type === 'application/json') {
      const reader = new FileReader()
      reader.onload = function() {
        const message = JSON.parse(reader.result).msg
        ElMessage({
          showClose: true,
          message: message,
          type: 'error'
        })
      }
      reader.readAsText(new Blob([res.data]))
    }
  } else {
    const downloadUrl = window.URL.createObjectURL(new Blob([res]))
    const a = document.createElement('a')
    a.style.display = 'none'
    a.href = downloadUrl
    a.download = fileName
    a.click()
  }
}
