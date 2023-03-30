

// 后端返回
export interface ListResponse{
    code: number;
    data: any;
    total:number;
  }

  export interface ResponseData{
    code: number;
    data: any;
    total:number;
  }

  // 前端请求
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