import { Service } from "@/services/services"


//路径参数传递
//获取物品列表
export function getArticlesInfoApi(userid,page) {
  return Service.get(`api/articles/${userid}`, {
    params: {
      page: page.page,
      key: page.key ,
      limit: page.limit,
      sort: page.sort,
      id: page.id
    }}
  )
}

//增加物品
export function addArticlesInfoApi(userid, data) {
  return Service.post(`api/articles/${userid}`, data)
}

//删除物品，data需要是一个数组
export function deleteArticlesInfoApi(userid, data) {
  return Service.delete(`api/articles/${userid}`, {data})
}
//修改物品
export function updateArticlesInfoApi(userid, data) {
  return Service.put(`api/articles/${userid}`, data)
}