/* eslint-disable space-before-function-paren */
import Cookies from 'js-cookie'
// App
const sidebarStatusKey = 'sidebar_status'
export const getSidebarStatus = () => Cookies.get(sidebarStatusKey)
export const setSidebarStatus = (sidebarStatus: string) => Cookies.set(sidebarStatusKey, sidebarStatus)

// User
const usernameKey = 'username'
export const setUserID = (username: string) => Cookies.set(usernameKey, username)
export const getUserID = () => Cookies.get(usernameKey)

const tokenKey = 'token'
export const getToken = () => Cookies.get(tokenKey)
export const setToken = (token: string) => Cookies.set(tokenKey, token)
export const removeToken = () => Cookies.remove(tokenKey)

const languageKey = 'language'
export const getLanguage = () => Cookies.get(languageKey)
export const setLanguage = (language: string) => Cookies.set(languageKey, language)

const rolesKey = 'roles'
export const getRoles = function (): number[] {
  const c = Cookies.get(rolesKey)
  if (c) {
    return JSON.parse(c)
  }
  return []
}

const selectStationsInfoKey = 'selectStationsInfo'
export const getSelectStationsInfo = function (): string[] {
  const c = Cookies.get(selectStationsInfoKey)
  if (c) {
    return JSON.parse(c)
  }
  return []
}
export const setSelectStationsInfo = function (selectStationsInfo: string[]) {
  Cookies.set(selectStationsInfoKey, selectStationsInfo)
}
export const removeSelectStationsInfo = () => Cookies.remove(selectStationsInfoKey)

const authorizedDepartmentsKey = 'authorizedDepartments'
const loginTypeKey = 'loginType'
const groupKey = 'group'
const workDateKey = 'workDate'

export const setUserInfo = function (ID: string, roles: number[], authorizedDepartments: string, loginType: string, group: string, workDate: string) {
  Cookies.set(usernameKey, ID)
  Cookies.set(rolesKey, JSON.stringify(roles))
  Cookies.set(authorizedDepartmentsKey, authorizedDepartments)
  Cookies.set(loginTypeKey, loginType)
  Cookies.set(groupKey, group)
  Cookies.set(workDateKey, workDate)
}
export const getAuthorizedDepartments = () => Cookies.get(authorizedDepartmentsKey)
export const getLoginType = () => Cookies.get(loginTypeKey)
export const getGroup = () => Cookies.get(groupKey)
export const getWorkDate = () => Cookies.get(workDateKey)

export const removeUserInfo = function () {
  Cookies.remove(rolesKey)
  Cookies.remove(authorizedDepartmentsKey)
  Cookies.remove(loginTypeKey)
  Cookies.remove(groupKey)
  Cookies.remove(workDateKey)
}

// info
const stationKey = 'station'
export const setStation = function (station: string) {
  Cookies.set(stationKey, station)
}
export const getStation = () => Cookies.get(stationKey)

const workOrderInfoKey = 'workOrderInfo'
export const setWorkOrderInfo = function (workOrder: any) {
  Cookies.set(workOrderInfoKey, workOrder)
}
export const getWorkOrderInfo = () => Cookies.get(workOrderInfoKey)
export const removeWorkOrderInfo = () => Cookies.remove(workOrderInfoKey)

const stationSiteInfoKey = 'stationSiteInfo'
export const setStationSiteInfo = function (stationSite: any) {
  Cookies.set(stationSiteInfoKey, stationSite)
}
export const getStationSiteInfo = () => Cookies.get(stationSiteInfoKey)
export const removeStationSiteInfo = () => Cookies.remove(stationSiteInfoKey)

const feedAndCollectModeKey = 'feedAndCollectMode'
export const setFeedAndCollectMode = function (feedAndCollectMode: any) {
  Cookies.set(feedAndCollectModeKey, feedAndCollectMode)
}
export const getFeedAndCollectMode = () => Cookies.get(feedAndCollectModeKey)
export const removeFeedAndCollectMode = () => Cookies.remove(feedAndCollectModeKey)

const mountResourcesInfoKey = 'mountResourcesInfo'
export const setMountResourcesInfo = function (mountResources: any) {
  Cookies.set(mountResourcesInfoKey, mountResources)
}
export const getMountResourcesInfo = () => Cookies.get(mountResourcesInfoKey)
export const removeMountResourcesInfo = () => Cookies.remove(mountResourcesInfoKey)
