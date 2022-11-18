import service from '@/utils/request'

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
