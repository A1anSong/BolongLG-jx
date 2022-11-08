import service from '@/utils/request'

export const createApply = (data) => {
  return service({
    url: '/testAppy/createApply',
    method: 'post',
    data
  })
}

export const deleteApply = (data) => {
  return service({
    url: '/testAppy/deleteApply',
    method: 'delete',
    data
  })
}

export const deleteApplyByIds = (data) => {
  return service({
    url: '/testAppy/deleteApplyByIds',
    method: 'delete',
    data
  })
}

export const updateApply = (data) => {
  return service({
    url: '/testAppy/updateApply',
    method: 'put',
    data
  })
}

export const findApply = (params) => {
  return service({
    url: '/testAppy/findApply',
    method: 'get',
    params
  })
}

export const getApplyList = (params) => {
  return service({
    url: '/testAppy/getApplyList',
    method: 'get',
    params
  })
}
