// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** 退出登录接口 POST /api/user/logout */
export async function logout(options?: { [key: string]: any }) {
  return request<Record<string, any>>('/api/user/logout', {
    method: 'POST',
    ...(options || {}),
  });
}

/** 登录接口 POST /api/login/account */
export async function login(body: API.LoginParams, options?: { [key: string]: any }) {
  return request<API.LoginResult>('/api/user/login', {
    method: 'POST',
    data: body,
    ...(options || {}),
  });
}

/** 此处后端没有提供注释 GET /api/notices */
export async function getNotices(options?: { [key: string]: any }) {
  return request<API.NoticeIconList>('/api/notices', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 获取规则列表 GET /api/rule */
export async function rule(
  params: {
    // query
    /** 当前的页码 */
    current?: number;
    /** 页面的容量 */
    pageSize?: number;
  },
  options?: { [key: string]: any },
) {
  return request<API.RuleList>('/api/rule', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 新建规则 PUT /api/rule */
export async function updateRule(options?: { [key: string]: any }) {
  return request<API.RuleListItem>('/api/rule', {
    method: 'PUT',
    ...(options || {}),
  });
}

/** 新建规则 POST /api/rule */
export async function addRule(options?: { [key: string]: any }) {
  return request<API.RuleListItem>('/api/rule', {
    method: 'POST',
    ...(options || {}),
  });
}

/** 删除规则 DELETE /api/rule */
export async function removeRule(options?: { [key: string]: any }) {
  return request<Record<string, any>>('/api/rule', {
    method: 'DELETE',
    ...(options || {}),
  });
}

/** Register Account POST /api/user/signup */
export async function registerAccount(body: API.RegisterParams, options?: { [key: string]: any }) {
  return request<Record<string, any>>('/api/user/sign-up', {
    method: 'POST',
    data: body,
    ...(options || {}),
  });
}

/** Verify Account POST /api/user/verify-account */
export async function verifyAccount(body: API.VerifyAccountParams, options?: { [key: string]: any }) {
  return request<Record<string, any>>('/api/user/verify-account', {
    method: 'POST',
    data: body,
    ...(options || {}),
  });
}

/** Get Code POST /api/user/send-mail */
export async function GetCode(options?: { [key: string]: any }) {
  return request<Record<string, any>>('/api/user/send-mail', {
    method: 'POST',
    ...(options || {}),
  });
}

export async function CreateStationMergeRequest(body: API.CreateStationMergeRequest, options?: { [key: string]: any }) {
  return request<API.CreateStationMergeRequestReply>('/api/station/merge-request', {
    method: 'POST',
    data: body,
    ...(options || {}),
  });
}

export async function GetStationMergeRequest(params: number, options?: { [key: string]: any }) {
  return request<any>(`/api/station/merge-request/${params}`, {
    method: 'GET',
    ...(options || {}),
  })
}