import service from '@/utils/request'

export const createInvoiceApply = (data) => {
  return service({
    url: '/testInvoiceApply/createInvoiceApply',
    method: 'post',
    data
  })
}

export const deleteInvoiceApply = (data) => {
  return service({
    url: '/testInvoiceApply/deleteInvoiceApply',
    method: 'delete',
    data
  })
}

export const deleteInvoiceApplyByIds = (data) => {
  return service({
    url: '/testInvoiceApply/deleteInvoiceApplyByIds',
    method: 'delete',
    data
  })
}

export const updateInvoiceApply = (data) => {
  return service({
    url: '/testInvoiceApply/updateInvoiceApply',
    method: 'put',
    data
  })
}

export const findInvoiceApply = (params) => {
  return service({
    url: '/testInvoiceApply/findInvoiceApply',
    method: 'get',
    params
  })
}

export const getInvoiceApplyList = (params) => {
  return service({
    url: '/testInvoiceApply/getInvoiceApplyList',
    method: 'get',
    params
  })
}

export const approveInvoiceApply = (data) => {
  return service({
    url: '/testInvoiceApply/approveInvoiceApply',
    method: 'put',
    data
  })
}

export const rejectInvoiceApply = (data) => {
  return service({
    url: '/testInvoiceApply/rejectInvoiceApply',
    method: 'put',
    data
  })
}