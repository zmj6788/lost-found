import { Service } from "@/services/services"

export function loginApi(data) {
  return Service.post("api/login", data)
}

export function registerApi(data) {
  return Service.post("api/register", data)
}

//路径参数传递
export function getUserInfoApi(userid) {
  return Service.get(`api/user/${userid}`)
}

//
export function setUserInfoApi(userid, data) {
  return Service.put(`api/user/${userid}`,data)
}


// //使用方法

// import { loginApi } from "@/api/user_api"

// //将login绑定到按钮

// //请求体
// const data = {
//   username: "admin",
//   password: "123456"
// }
// async function login() {
//   const res = await loginApi(data)
//   //响应体
//   console.log(res)
// }