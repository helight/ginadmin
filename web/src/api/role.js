import request from '@/utils/request'

export function requestList(query) {
  return request({
    url: '/api/role/list',
    method: 'get',
    params: query
  })
}

export function requestDetail(id) {
  return request({
    url: '/api/role/detail',
    method: 'get',
    params: { id }
  })
}

export function requestUpdate(data) {
  return request({
    url: '/api/role/update',
    method: 'post',
    data
  })
}

export function requestCreate(data) {
  return request({
    url: '/api/role/create',
    method: 'post',
    data
  })
}

export function requestDelete(data) {
  return request({
    url: '/api/role/delete',
    method: 'post',
    data
  })
}

export function requestRoleMenuIDList(roleid) {
  return request({
    url: '/api/role/rolemenuidlist',
    method: 'get',
    params: { roleid }
  })
}

export function requestSetRole(roleid, data) {
  return request({
    url: '/api/role/setrole',
    method: 'post',
    params: { roleid },
    data
  })
}

export function requestAll() {
  return request({
    url: '/api/role/allrole',
    method: 'get'
  })
}
