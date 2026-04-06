import Cookies from "js-cookie";
import { refreshJwtToken } from "./refresh-jwt-token";

export async function refreshJwtTokenIfNeeded() {

  const currentJwtToken = Cookies.get("jwt");
  if (currentJwtToken === undefined || currentJwtToken === "") {
    return null;
  }
  const currentJwtTokenPayload = JSON.parse(atob(currentJwtToken.split(".")[1]));
  const expirationDate = new Date(currentJwtTokenPayload.exp * 1000);

  if (expirationDate.getTime() - 1000 * 60 * 5 < new Date().getTime()) {

    await refreshJwtToken()
  }
}