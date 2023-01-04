import service from '@/utils/request'

export const createTemplate = (data) => {
  return service({
    url: '/testTemplate/createTemplate',
    method: 'post',
    data
  })
}

export const deleteTemplate = (data) => {
  return service({
    url: '/testTemplate/deleteTemplate',
    method: 'delete',
    data
  })
}

export const deleteTemplateByIds = (data) => {
  return service({
    url: '/testTemplate/deleteTemplateByIds',
    method: 'delete',
    data
  })
}

export const updateTemplate = (data) => {
  return service({
    url: '/testTemplate/updateTemplate',
    method: 'put',
    data
  })
}

export const findTemplate = (params) => {
  return service({
    url: '/testTemplate/findTemplate',
    method: 'get',
    params
  })
}

export const getTemplateList = (params) => {
  return service({
    url: '/testTemplate/getTemplateList',
    method: 'get',
    params
  })
}
