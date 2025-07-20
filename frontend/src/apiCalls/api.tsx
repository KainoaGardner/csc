import API_URL from "../env.tsx"
import { useApp, useLogoutHandler } from "../appContext/useApp.tsx"

export function useFetchWithAuth() {
  const { accessToken, setAccessToken } = useApp();
  const { handleLogout } = useLogoutHandler()

  async function doFetch(url: string, options: RequestInit = {}, token: string | null) {
    return fetch(url, {
      ...options,
      headers: {
        ...options.headers,
        Authorization: `Bearer ${token}`,
      }
    })
  }

  async function fetchWithAuth(url: string, options: RequestInit = {}) {
    let token = accessToken
    let response = await doFetch(url, options, token)

    if (response.status == 401) {
      const refreshedToken = await tryRefreshToken()
      if (!refreshedToken) {
        handleLogout()
        throw new Error("Session expired")
      }

      token = refreshedToken
      setAccessToken(token)

      response = await doFetch(url, options, token)
    }

    return response;
  }

  return fetchWithAuth
}

export async function tryRefreshToken(): Promise<string | null> {
  const response = await fetch(API_URL + "auth/refresh", {
    method: "POST",
    credentials: "include",
  })

  if (!response.ok) {
    return null
  }

  const data = await response.json()
  return data.accessToken
}




