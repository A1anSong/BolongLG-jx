import service from '@/utils/request'

export const createLogout = (data) => {
  return service({
    url: '/testLogout/createLogout',
    method: 'post',
    data
  })
}

export const deleteLogout = (data) => {
  return service({
    url: '/testLogout/deleteLogout',
    method: 'delete',
    data
  })
}

export const deleteLogoutByIds = (data) => {
  return service({
    url: '/testLogout/deleteLogoutByIds',
    method: 'delete',
    data
  })
}

export const updateLogout = (data) => {
  return service({
    url: '/testLogout/updateLogout',
    method: 'put',
    data
  })
}

export const findLogout = (params) => {
  return service({
    url: '/testLogout/findLogout',
    method: 'get',
    params
  })
}

export const getLogoutList = (params) => {
  return service({
    url: '/testLogout/getLogoutList',
    method: 'get',
    params
  })
}
