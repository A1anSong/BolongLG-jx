import service from '@/utils/request'

export const createDelay = (data) => {
  return service({
    url: '/testDelay/createDelay',
    method: 'post',
    data
  })
}

export const deleteDelay = (data) => {
  return service({
    url: '/testDelay/deleteDelay',
    method: 'delete',
    data
  })
}

export const deleteDelayByIds = (data) => {
  return service({
    url: '/testDelay/deleteDelayByIds',
    method: 'delete',
    data
  })
}

export const updateDelay = (data) => {
  return service({
    url: '/testDelay/updateDelay',
    method: 'put',
    data
  })
}

export const findDelay = (params) => {
  return service({
    url: '/testDelay/findDelay',
    method: 'get',
    params
  })
}

export const getDelayList = (params) => {
  return service({
    url: '/testDelay/getDelayList',
    method: 'get',
    params
  })
}
