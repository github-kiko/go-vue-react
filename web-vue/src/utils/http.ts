import request from "./request"
const http={
    get(url:string,params:any){
        return request({
            url:url,
            method:'get',
            params:params,
        })
    },
    post(url:string,params:any){
        return request({
            url:url,
            method:'post',
            data:params
        })

    },
    put(url:string,params:any){
        return request({
            url:url,
            method:'put',
            data:params


        })
    },
    delete(url:string,params:any){
        return request({
            url:url,
            method:'delete',
            data:params
        })
    }

}

export default http