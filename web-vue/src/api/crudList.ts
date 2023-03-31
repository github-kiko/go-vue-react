import { AxiosResponse } from 'axios';
import {ListResponse,ResponseData,listaddRequset, listdeleteRequset,listupdateRequset,listqueryRequset,listqueryResponse} from "../model/crudList";
import http from "../utils/http"

// 新
  export async function listadd(params: listaddRequset): Promise<ListResponse<listqueryResponse>> {
    const response: AxiosResponse<ResponseData<listqueryResponse>> = await http.post('/list/add',params);
    return response.data;
  }


 // 删除
  export async function listdelete(params: listdeleteRequset): Promise<ListResponse<listqueryResponse>> {
    const response: AxiosResponse<ResponseData<listqueryResponse>> = await http.delete(`/list/delete/${params.id}`,params);
    return response.data;
}

// 修改
 export async function listupdate(params: listupdateRequset): Promise<ListResponse<listqueryResponse>> {
    const response: AxiosResponse<ResponseData<listqueryResponse>> = await http.put(`/list/update/${params.id}`,params);
    return response.data;
}

// 查询
export async function listquery(params: listqueryRequset): Promise<ListResponse<listqueryResponse>> {
    const response: AxiosResponse<ResponseData<listqueryResponse>> = await http.get('/list/query',params);
    return response.data;
}




 








