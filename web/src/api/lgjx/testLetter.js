import service from '@/utils/request'

export const createLetter = (data) => {
  return service({
    url: '/testLetter/createLetter',
    method: 'post',
    data
  })
}

export const deleteLetter = (data) => {
  return service({
    url: '/testLetter/deleteLetter',
    method: 'delete',
    data
  })
}

export const deleteLetterByIds = (data) => {
  return service({
    url: '/testLetter/deleteLetterByIds',
    method: 'delete',
    data
  })
}

export const updateLetter = (data) => {
  return service({
    url: '/testLetter/updateLetter',
    method: 'put',
    data
  })
}

export const findLetter = (params) => {
  return service({
    url: '/testLetter/findLetter',
    method: 'get',
    params
  })
}

export const getLetterList = (params) => {
  return service({
    url: '/testLetter/getLetterList',
    method: 'get',
    params
  })
}
