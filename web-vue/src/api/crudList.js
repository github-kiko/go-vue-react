import http from "@/utils/http"
// 新增
export  const listadd=params=>http.post('/list/add',params)
// 删除
export  const listdelete=params=>http.get(`/list/delete/${params.id}`,params)
// 修改
export  const listupdate=params=>http.get(`/list/update/${params.id}`,params)
// 删除
export  const listquery=params=>http.get('/list/query',params)


