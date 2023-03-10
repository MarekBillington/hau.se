import { $ } from "@builder.io/qwik";
import { refreshReq, request } from "../api/api";

import type Auth from "./interface/auth";
import type User from "../../interfaces/user";
import type UserSession from "./interface/user-session";
import type Portfolio from "../../interfaces/portfolio";

// export const url = 'http://localhost:8002/api/'
export const url = "http://dev.hau.se/api/";

export const login = $(async (email: string, password: string) => {
  // @todo hash password for sending
  try {
  const res = await fetch(url + "auth/login", {
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
  } catch (err) {
    console.log(err);
    return {};
  }
});

export const signup = $(async (reg: any) => {
  // @todo hash password for sending
  try {
    const res = await fetch(url + "auth/register", {
      method: "POST",
      body: JSON.stringify({
        email: reg.email,
        password: reg.password,
        first_name: reg.first_name,
        last_name: reg.last_name
      }),
    })
      .then((response) => response.json())
      .then((obj) => {
        return obj;
      });
    return res;
  } catch (err) {
    console.log(err);
    return {};
  }
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
    .then((res) => {
      sess.user = res.user
      sess.portfolio = res.portfolio
    })
})

// Just unset all the properties for the stateful objects
export const destroyUser = $((auth: Auth, sess: UserSession) => {
  auth.token = ""
  auth.expiry = 0

  sess.user = {
    id: 0,
    email: "",
    first_name: "",
    last_name: "",
  } as User
  sess.portfolio = {
    id: 0,
    name: "",
  } as Portfolio
})

export const logout = async (auth: Auth, sess: UserSession) => {
  await request("auth/logout", auth, "POST")
  
  destroyUser(auth, sess)
}
