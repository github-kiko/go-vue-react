import axios from 'axios'


// 创建axios实例
const conf = {
    
     baseURL:"http://127.0.0.1:3000",

    // baseURL: '/api', // 根路径

   
    timeout: 90000, // 请求超时时间
    // withCredentials:true ,
  };

  const service = axios.create(conf);
  

//请求拦截器        
service.interceptors.request.use(
    config=>{
        // console.log("config:",config);
        return config
    },
    error=>{
        return   Promise.reject(error);
    },
    
    )

// 相应拦截器
service.interceptors.response.use(
    response=>{
        // console.log("response:",response);
     const   res =response.data
        return res
        
    },
   error=>{
       return Promise.reject(error);
   }
)


export   default service;
export { conf };