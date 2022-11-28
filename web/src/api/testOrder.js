import service from '@/utils/request'
import { ElMessage } from 'element-plus'

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

export const downloadExcel = (params) => {
  return service({
    url: '/testOrder/exportExcel',
    method: 'get',
    params: params,
    responseType: 'blob'
  }).then((res) => {
    handleExcelError(res)
  })
}

const handleExcelError = (res) => {
  if (typeof (res.data) !== 'undefined') {
    if (res.data.type === 'application/json') {
      const reader = new FileReader()
      reader.onload = function() {
        const message = JSON.parse(reader.result).msg
        ElMessage({
          showClose: true,
          message: message,
          type: 'error'
        })
      }
      reader.readAsText(new Blob([res.data]))
    }
  } else {
    const downloadUrl = window.URL.createObjectURL(new Blob([res]))
    const a = document.createElement('a')
    a.style.display = 'none'
    a.href = downloadUrl
    a.download = (new Date().getTime()).toString() + '.xlsx'
    a.click()
  }
}
