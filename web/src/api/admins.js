import request from '@/utils/request'

export function requestList(query) {
  return request({
    url: '/api/admins/list',
    method: 'get',
    params: query
  })
}

export function requestDetail(id) {
  return request({
    url: '/api/admins/detail',
    method: 'get',
    params: { id }
  })
}

export function requestUpdate(data) {
  return request({
    url: '/api/admins/update',
    method: 'post',
    data
  })
}

export function requestCreate(data) {
  return request({
    url: '/api/admins/create',
    method: 'post',
    data
  })
}

export function requestDelete(data) {
  return request({
    url: '/api/admins/delete',
    method: 'post',
    data
  })
}

export function requestAdminsRoleIDList(adminsid) {
  return request({
    url: '/api/admins/adminsroleidlist',
    method: 'get',
    params: { adminsid }
  })
}

export function requestSetRole(adminsid, data) {
  return request({
    url: '/api/admins/setrole',
    method: 'post',
    params: { adminsid },
    data
  })
}
