import service from '@/utils/request'

export const createClaim = (data) => {
  return service({
    url: '/testClaim/createClaim',
    method: 'post',
    data
  })
}

export const deleteClaim = (data) => {
  return service({
    url: '/testClaim/deleteClaim',
    method: 'delete',
    data
  })
}

export const deleteClaimByIds = (data) => {
  return service({
    url: '/testClaim/deleteClaimByIds',
    method: 'delete',
    data
  })
}

export const updateClaim = (data) => {
  return service({
    url: '/testClaim/updateClaim',
    method: 'put',
    data
  })
}

export const findClaim = (params) => {
  return service({
    url: '/testClaim/findClaim',
    method: 'get',
    params
  })
}

export const getClaimList = (params) => {
  return service({
    url: '/testClaim/getClaimList',
    method: 'get',
    params
  })
}
