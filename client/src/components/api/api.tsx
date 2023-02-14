

/**
 * Aggregating request function
 * 
 * @param endpoint 
 * @param method 
 * @param body 
 * @returns 
 */
export async function request(endpoint: string, method: string, body?: {}) {
    // @todo https://vitejs.dev/guide/env-and-mode.html
    const url = 'http://dev.hau.se/api/'
    // const url = 'http://localhost/api/'

    let header = {}
    if (method != 'GET') {
        header = {
            method: method,
            body: body
        }
    }
    
    try {
        const res = await fetch(
            url + endpoint,
            header
        )
        // @todo look at making typed objects that are returned?
        return await res.json()
    } catch (err) {
        console.log(err)
        return {}
    }
}