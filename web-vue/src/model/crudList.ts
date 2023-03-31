// 后端返回结构体
export interface ListResponse<T>{
    code: number;
    data: T;
    total:number;
  }

  export interface ResponseData<T>{
    code: number;
    data: T;
    total:number;
  }


  export interface listqueryResponse{
    name: string;
    age: number;
    school: string;
    phone: string;
    address: string;
  }


  // 前端请求结构体
  export interface listaddRequset{
    name: string;
    age: number;
    school: string;
    phone: string;
    address: string;
  }

  export interface listdeleteRequset{
    id: number;
   
  }

  export interface listupdateRequset{
    name: string;
    id: number;
    age: number;
    school: string;
    phone: string;
    address: string;
  }

  export interface listqueryRequset{
    name: string;
    page: number;
    pageSize:number;
  }