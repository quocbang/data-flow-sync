import { request } from "@umijs/max";

/** check whether user exist /rest/user/:userID/check */
export async function checkIsUserExisted(params: string, options?: { [key: string]: any }) {
  return request<Graph.CheckUserResult>(`/api/rest/user/${params}/check`, {
    method: 'GET',
    ...(options || {}),
  });
}

export async function checkIsEmailExisted(params: string, options?: { [key: string]: any }) {
  return request<Graph.CheckEmailResult>(`/api/rest/email/${params}/check`, {
    method: "GET",
    ...(options || {}),
  })
}