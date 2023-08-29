import Cookies from "js-cookie"

function calculateTimeRemainingInSeconds(targetTimestamp: number) {
    const currentTimestamp = Math.floor(Date.now() / 1000); // Current Unix timestamp in seconds
    const timeRemainingInSeconds = targetTimestamp - currentTimestamp;

    if (timeRemainingInSeconds < 0 ){
        return 0
    }
    return timeRemainingInSeconds;
}

const tokenKey = "token-key"
export const setToken = (token: string, tokenExpiryTime: number) => {
    const tokenRemainingTime = calculateTimeRemainingInSeconds(tokenExpiryTime) // return remaining time that unit is seconds
    // 86400 is entire seconds of one day    
    Cookies.set(tokenKey, token, {expires: tokenRemainingTime / 86400})
}
export const getToken = () => Cookies.get(tokenKey)
export const removeToken = () => Cookies.remove(tokenKey)
