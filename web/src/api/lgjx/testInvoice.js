import service from '@/utils/request'

export const createInvoice = (data) => {
  return service({
    url: '/testInvoice/createInvoice',
    method: 'post',
    data
  })
}

export const deleteInvoice = (data) => {
  return service({
    url: '/testInvoice/deleteInvoice',
    method: 'delete',
    data
  })
}

export const deleteInvoiceByIds = (data) => {
  return service({
    url: '/testInvoice/deleteInvoiceByIds',
    method: 'delete',
    data
  })
}

export const updateInvoice = (data) => {
  return service({
    url: '/testInvoice/updateInvoice',
    method: 'put',
    data
  })
}

export const findInvoice = (params) => {
  return service({
    url: '/testInvoice/findInvoice',
    method: 'get',
    params
  })
}

export const getInvoiceList = (params) => {
  return service({
    url: '/testInvoice/getInvoiceList',
    method: 'get',
    params
  })
}
