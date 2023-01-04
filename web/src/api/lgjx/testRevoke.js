import service from '@/utils/request'

export const createRevoke = (data) => {
  return service({
    url: '/testRevoke/createRevoke',
    method: 'post',
    data
  })
}

export const deleteRevoke = (data) => {
  return service({
    url: '/testRevoke/deleteRevoke',
    method: 'delete',
    data
  })
}

export const deleteRevokeByIds = (data) => {
  return service({
    url: '/testRevoke/deleteRevokeByIds',
    method: 'delete',
    data
  })
}

export const updateRevoke = (data) => {
  return service({
    url: '/testRevoke/updateRevoke',
    method: 'put',
    data
  })
}

export const findRevoke = (params) => {
  return service({
    url: '/testRevoke/findRevoke',
    method: 'get',
    params
  })
}

export const getRevokeList = (params) => {
  return service({
    url: '/testRevoke/getRevokeList',
    method: 'get',
    params
  })
}
