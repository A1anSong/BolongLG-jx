import service from '@/utils/request'

export const createRefund = (data) => {
  return service({
    url: '/testRefund/createRefund',
    method: 'post',
    data
  })
}

export const deleteRefund = (data) => {
  return service({
    url: '/testRefund/deleteRefund',
    method: 'delete',
    data
  })
}

export const deleteRefundByIds = (data) => {
  return service({
    url: '/testRefund/deleteRefundByIds',
    method: 'delete',
    data
  })
}

export const updateRefund = (data) => {
  return service({
    url: '/testRefund/updateRefund',
    method: 'put',
    data
  })
}

export const findRefund = (params) => {
  return service({
    url: '/testRefund/findRefund',
    method: 'get',
    params
  })
}

export const getRefundList = (params) => {
  return service({
    url: '/testRefund/getRefundList',
    method: 'get',
    params
  })
}
