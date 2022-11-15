import service from '@/utils/request'

export const createProject = (data) => {
  return service({
    url: '/testProject/createProject',
    method: 'post',
    data
  })
}

export const deleteProject = (data) => {
  return service({
    url: '/testProject/deleteProject',
    method: 'delete',
    data
  })
}

export const deleteProjectByIds = (data) => {
  return service({
    url: '/testProject/deleteProjectByIds',
    method: 'delete',
    data
  })
}

export const updateProject = (data) => {
  return service({
    url: '/testProject/updateProject',
    method: 'put',
    data
  })
}

export const findProject = (params) => {
  return service({
    url: '/testProject/findProject',
    method: 'get',
    params
  })
}

export const getProjectList = (params) => {
  return service({
    url: '/testProject/getProjectList',
    method: 'get',
    params
  })
}
