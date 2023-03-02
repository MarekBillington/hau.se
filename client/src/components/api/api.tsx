import { getTokenOrRefresh } from "../auth/auth";
import type Auth from "../interfaces/auth";

// @todo https://vitejs.dev/guide/env-and-mode.html
// const url = 'http://localhost:8002/api/'
const url = "http://dev.hau.se/api/";

/**
 * Aggregating request function
 *
 * @param endpoint
 * @param method
 * @param body
 * @returns
 */
export async function request(endpoint: string, auth: Auth, method: string, body?: {}) {
  const token = await getTokenOrRefresh(auth)

  let options: any = {}
  options = {
    method: method,
    headers: {
      Authorization: "Bearer " + token,
    },
  };

  // @ts-ignore shutup typescript
  if (typeof body !== "undefined" && Object.keys(body).length != 0) {
    options.body = body
  }

  try {
    // add token to header

    // @ts-ignore No overload matches this call. ts(2769)
    const res = fetch(url + endpoint, options)
      .then((response) => response.json())
      .then((obj) => {
        return obj;
      });

    // @todo look at making typed objects that are returned?
    return res;
  } catch (err) {
    console.log(err);
    return {};
  }
}

export async function refreshReq() {
  try {
    const res = await fetch(url + "auth/refresh", { method: "POST" })
      .then((response) => response.json())
      .then((obj) => {
        return obj;
      });

    return res;
  } catch (err) {
    console.log(err);
    return {};
  }
}
