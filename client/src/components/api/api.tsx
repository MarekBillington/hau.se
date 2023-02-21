
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
    // const url = 'http://localhost:8002/api/'

    let header = {}
    if (method != 'GET') {
        header = {
            method: method,
            body: body
        }
        console.log(header);
    }
    
    try {
        const res = await fetch(
            url + endpoint,
            header
        ).then((response) => response.json())
        .then((obj) => {
            return obj;
        });
        
        // @todo look at making typed objects that are returned?
        return res;
    } catch (err) {
        console.log(err)
        return {};
    }
}