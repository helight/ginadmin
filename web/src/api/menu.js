import request from '@/utils/request'

export function requestList(query) {
  return request({
    url: '/api/menu/list',
    method: 'get',
    params: query
  })
}

export function requestAll() {
  return request({
    url: '/api/menu/allmenu',
    method: 'get'
  })
}

export function requestDetail(id) {
  return request({
    url: '/api/menu/detail',
    method: 'get',
    params: { id }
  })
}

export function requestUpdate(data) {
  return request({
    url: '/api/menu/update',
    method: 'post',
    data
  })
}

export function requestCreate(data) {
  return request({
    url: '/api/menu/create',
    method: 'post',
    data
  })
}

export function requestDelete(data) {
  return request({
    url: '/api/menu/delete',
    method: 'post',
    data
  })
}

export function requestMenuButton(menucode) {
  return request({
    url: '/api/menu/menubuttonlist',
    method: 'get',
    params: { menucode }
  })
}
