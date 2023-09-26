import jwtDecode from "jwt-decode";

export const decodeToken = (token: string): API.CurrentUser => {
    return jwtDecode(token);
};
  