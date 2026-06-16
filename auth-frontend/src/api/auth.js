import api_url from "./api"

async function FetchReq(method, reqdata) {

    return await fetch(`${api_url}/${method}`, reqdata)
}

export async function authRegister(data) {


    const res = await FetchReq("register", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    })

    return res

}


export async function authLogin(data) {


    const res = await FetchReq("login", {
        method: "POST",
        credentials:"include",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)

    })


    return res

}

export async function authProfile() {


    const res=await FetchReq("profile",{
        credentials:"include"
    })

    return res

}

export async function authLogout() {


    return FetchReq("logout",{
        method: "POST",
        credentials:"include"

    })


}