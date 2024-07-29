<template>
  <div class="login">
    <aside>
      <img src="E:\lost-found\src\assets\images\show.webp" alt="">
    </aside>
    <div class="main">
      <div class="logo">
        <div>失物找回系统</div>
        <div>lost-found</div>
      </div>
      <div class="form">
        <a-card :title="formState.status === false ? '用户注册' : '用户登录'" style=" opacity: 0.7; color: white;">
          <a-form :model="formState" name="normal_login" class="login-form" @finish="onFinish"
            @finishFailed="onFinishFailed">
            <a-form-item label="账户" name="username" :rules="[{ required: true, message: '请输入你的账户' }]">
              <a-input v-model:value="formState.username">
                <template #prefix>
                  <UserOutlined class="site-form-item-icon" />
                </template>
              </a-input>
            </a-form-item>

            <a-form-item label="密码" name="password" :rules="[{ required: true, message: '请输入你的密码' }]">
              <a-input-password v-model:value="formState.password">
                <template #prefix>
                  <LockOutlined class="site-form-item-icon" />
                </template>
              </a-input-password>
            </a-form-item>

            <a-form-item>
              <a-button :disabled="disabled" type="primary" html-type="submit" class="login-form-button"
                v-if="formState.status">
                登录
              </a-button>
              <a-button :disabled="disabled" type="primary" html-type="submit" class="login-form-button" v-else>
                注册
              </a-button>
              或
              <a @click="actions" v-if="formState.status">去注册</a>
              <a @click="actions" v-else>去登录</a>
            </a-form-item>
          </a-form>
        </a-card>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive, computed } from 'vue';
import { useRouter } from 'vue-router';
import { loginApi,registerApi,getUserInfoApi } from "@/api/user_api"
import { message } from 'ant-design-vue';
import { useStore } from '@/stores/store'
//保存userinfo
const store = useStore()
const router = useRouter();
const formState = reactive({
  username: '',
  password: '',
  remember: true,
  status: true,
});
//通过formState.status判断还是注册
const actions = () => {
  formState.status = !formState.status;
  // console.log(formState.status);
  
};

const onFinish = async () => {
  var data = {
    user_name: formState.username,
    password: formState.password,
  }
  //首先判断是登陆事件还是注册事件
  // console.log(formState.status);
  if(formState.status){
    const res = await loginApi(data)
    //   //响应体
    if(res.msg === '登录成功'){
      
      //登陆成功保存userinfo到全局变量
      // console.log(res)
      // console.log(res.data.userid)
      const userinfo = await getUserInfoApi(res.data.userid)
      localStorage.setItem('user', JSON.stringify(userinfo.data));
      store.loadUserInfo()
      message.success(res.msg)
      router.push('/web')
    } else if (res.msg === '管理员登录'){
      const userinfo = await getUserInfoApi(res.data.userid)
      localStorage.setItem('user', JSON.stringify(userinfo.data));
      store.loadUserInfo()
      message.success(res.msg)
      router.push('/admin')

    }else {
      message.error(res.msg)

    }
  }else{
    const res = await registerApi(data)
    // console.log(res)
    //   //响应体
    if (res.msg === '注册成功'){
      message.success(res.msg)
    }else{
      message.error(res.msg)
    }
  }

};
const onFinishFailed = errorInfo => {
  console.log('Failed:', errorInfo);
};
const disabled = computed(() => {
  return !(formState.username && formState.password);
});


</script>

<style lang="scss">
.login {
  display: flex;
  flex-direction: row;
  aside {
    width: 1900px;
    height: 100vh;
    background-color: white;
    img{
      width: 1000px;
      height: 100vh;
       border: 2px solid rgba(0, 0, 0, 0.1);
  border-radius: 10px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    }
  }

  .main {
    width: calc(100% - 240px);
    height: 100vh;
    background-color: #87c0ca;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
     border: 2px solid rgba(0, 0, 0, 0.1);
  border-radius: 10px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    .logo {
        width: 90%;
        transform: scale(0.7);
        font-family: "宋体", sans-serif;
        margin-bottom: 50px;
        display: flex;
        /* 使用Flexbox布局 */
        flex-direction: column;
        /* 设置为垂直堆叠 */
        align-items: center;
        /* 水平居中 */
        justify-content: center;
        /* 垂直居中，假设.logo容器本身需要垂直居中于其父容器，否则可以省略这条 */
      }
    
      .logo>div:first-child {
        font-size: 50px;
        margin-bottom: 5px;
        /* 可选，添加间距 */
      }
    
      .logo>div:last-child {
        font-size: 50px;
      }
    .form{
      
    }
    
  }
}
</style>