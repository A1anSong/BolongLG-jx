import service from '@/utils/request'

export const createPay = (data) => {
  return service({
    url: '/testPay/createPay',
    method: 'post',
    data
  })
}

export const deletePay = (data) => {
  return service({
    url: '/testPay/deletePay',
    method: 'delete',
    data
  })
}

export const deletePayByIds = (data) => {
  return service({
    url: '/testPay/deletePayByIds',
    method: 'delete',
    data
  })
}

export const updatePay = (data) => {
  return service({
    url: '/testPay/updatePay',
    method: 'put',
    data
  })
}

export const findPay = (params) => {
  return service({
    url: '/testPay/findPay',
    method: 'get',
    params
  })
}

export const getPayList = (params) => {
  return service({
    url: '/testPay/getPayList',
    method: 'get',
    params
  })
}
