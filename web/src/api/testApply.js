import service from '@/utils/request'

export const createApply = (data) => {
  return service({
    url: '/testApply/createApply',
    method: 'post',
    data
  })
}

export const deleteApply = (data) => {
  return service({
    url: '/testApply/deleteApply',
    method: 'delete',
    data
  })
}

export const deleteApplyByIds = (data) => {
  return service({
    url: '/testApply/deleteApplyByIds',
    method: 'delete',
    data
  })
}

export const updateApply = (data) => {
  return service({
    url: '/testApply/updateApply',
    method: 'put',
    data
  })
}

export const findApply = (params) => {
  return service({
    url: '/testApply/findApply',
    method: 'get',
    params
  })
}

export const getApplyList = (params) => {
  return service({
    url: '/testApply/getApplyList',
    method: 'get',
    params
  })
}
