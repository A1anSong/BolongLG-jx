import service from '@/utils/request'

export const createOrder = (data) => {
  return service({
    url: '/testOrder/createOrder',
    method: 'post',
    data
  })
}

export const deleteOrder = (data) => {
  return service({
    url: '/testOrder/deleteOrder',
    method: 'delete',
    data
  })
}

export const deleteOrderByIds = (data) => {
  return service({
    url: '/testOrder/deleteOrderByIds',
    method: 'delete',
    data
  })
}

export const updateOrder = (data) => {
  return service({
    url: '/testOrder/updateOrder',
    method: 'put',
    data
  })
}

export const findOrder = (params) => {
  return service({
    url: '/testOrder/findOrder',
    method: 'get',
    params
  })
}

export const getOrderList = (params) => {
  return service({
    url: '/testOrder/getOrderList',
    method: 'get',
    params
  })
}

export const getOrderStatisticData = (params) => {
  return service({
    url: '/testOrder/getOrderStatisticData',
    method: 'get',
    params
  })
}