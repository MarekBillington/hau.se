import { $ } from "@builder.io/qwik";
import { refreshReq, request } from "../api/api";

import type Auth from "../interfaces/auth";
import type User from "../interfaces/user";
import type UserSession from "../interfaces/user-session";

export const login = $(async (email: string, password: string) => {
  // @todo hash password for sending
  const res = await fetch("http://dev.hau.se/api/auth/login", {
    method: "POST",
    body: JSON.stringify({
      email: email,
      password: password,
    }),
  })
    .then((response) => response.json())
    .then((obj) => {
      return obj;
    });
  return res;
});

export const signup = $(async (reg: any) => {
  // @todo hash password for sending
  const res = await fetch("http://dev.hau.se/api/auth/register", {
    method: "POST",
    body: JSON.stringify({
      email: reg.email,
      password: reg.password,
      firstName: reg.firstName,
      lastName: reg.lastName
    }),
  })
    .then((response) => response.json())
    .then((obj) => {
      return obj;
    });
  return res;
});

export const getTokenOrRefresh = $(async (auth: Auth) => {
    if (auth.expiry <= Math.floor(Date.now() / 1000)) {
        await refreshToken(auth)
    }
    return auth.token
})

export const refreshToken = $(async (auth: Auth) => {
  const res = await refreshReq();
  
  if (Object.keys(res).length != 0) {
    auth.token = res.token
    auth.expiry = res.expiry
  }
});

export const getUserInfo = $(async (auth: Auth, sess: UserSession) => {  
  return await request("user/setup", auth, "GET")
    .then((u: User) => {
      sess.user = u
    })
})

// Just unset all the properties for the stateful objects
export const destroyUser = $((auth: Auth, sess: UserSession) => {
  auth.token = ""
  auth.expiry = 0

  sess.user = {
    id: 0,
    email: "",
    firstName: "",
    lastName: "",
  } as User
  sess.portfolio = 0
})

export const logout = async (auth: Auth, sess: UserSession) => {
  await request("auth/logout", auth, "POST")
  
  destroyUser(auth, sess)
}
