import request from "@/utils/request"
const http={
    get(url,params){
        return request({
            url:url,
            method:'get',
            params:params,
        })
    },
    post(url,params){
        return request({
            url:url,
            method:'post',
            data:params
        })

    },
    put(url,params){
        return request({
            url:url,
            method:'put',
            data:params


        })
    },
    delete(url,params){
        return request({
            url:url,
            method:'delete',
            data:params
        })
    }

}

export default http