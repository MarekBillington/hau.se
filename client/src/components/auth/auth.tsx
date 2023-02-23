import { $ } from "@builder.io/qwik"
import { request } from "../api/api"


export const checkAuth = $(() => {

    // @todo https://github.com/BuilderIO/qwik/issues/2451

    // const token = document.cookie
    //         .split("; ")
    //         .find((row) => row.startsWith("tkn="))
    //         ?.split("="[1])

    // const expiry = document.cookie
    //     .split("; ")
    //     .find((row) => row.startsWith("exp="))
    //     ?.split("="[1])

    // console.log(expiry)
    // console.log(token)

    console.log(document.cookie)
    
    //const now = Math.floor(Date.now())
    // let val = true

    return true
})


export const login = $(async (email: string, password: string) => {
    // @todo hash password for sending
    await request("auth/login", "POST", JSON.stringify({
        "email": email,
        "password": password
    }))
})

export const getToken = $(() => {
    const token = document.cookie
            .split("; ")
            .find((row) => row.startsWith("tkn="))
            ?.split("="[1])
    
    return token
})