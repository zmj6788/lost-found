<template>
  <div class="user">
    <a-card title="个人信息配置" :bordered="false" style="width: 700px;  background-color: #ecf0f1;">
      <a-form :model="formState" v-bind="layout" name="nest-messages" :validate-messages="validateMessages"
        @finish="onFinish">
        <a-form-item :name="['user', 'nick_name']" label="昵称" :rules="[{ required: true }]">
          <a-input v-model:value="formState.user.nick_name" />
        </a-form-item>
        <a-form-item :name="['user', 'user_name']" label="用户名" :rules="[{ required: true }]">
          <a-input v-model:value="formState.user.user_name" />
        </a-form-item>
        <a-form-item :name="['user', 'password']" label="密码" :rules="[{ required: true}]">
          <a-input-password v-model:value="formState.user.password" />
        </a-form-item>
        <a-form-item :name="['user', 'avatar_id']" label="头像地址" :rules="[{ required: true}]">
          <a-input v-model:value="formState.user.avatar_id" />
        </a-form-item>
        <a-form-item :name="['user', 'tel']" label="联系电话" :rules="[{ required: true}]">
          <a-input v-model:value="formState.user.tel" />
        </a-form-item>
        <a-form-item :name="['user', 'addr']" label="地址" :rules="[{ required: true}]">
          <a-input v-model:value="formState.user.addr" />
        </a-form-item>
        <a-form-item :wrapper-col="{ ...layout.wrapperCol, offset: 8 }">
          <a-button type="primary" html-type="submit">修改个人信息</a-button>
        </a-form-item>
      </a-form>
    </a-card>

  </div>
</template>

<script setup>
import { reactive } from 'vue';
import { useStore } from '@/stores/store'
import { setUserInfoApi } from "@/api/user_api"
import { message } from 'ant-design-vue';
//使用userid获取用户信息
const store = useStore()
store.loadUserInfo()

const layout = {
  labelCol: {
    span: 8,
  },
  wrapperCol: {
    span: 16,
  },
};
const validateMessages = {
  required: '${label} 是必需的!',
};
const formState = reactive({
  user: {
    nick_name: '  ',
    user_name: '',
    password: '',
    avatar_id: '',
    tel: '',
    addr: '',
  },
});
formState.user.nick_name = store.userInfo.nick_name;
formState.user.user_name = store.userInfo.user_name;
formState.user.password = store.userInfo.password;
formState.user.avatar_id = store.userInfo.avatar_id;
formState.user.tel = store.userInfo.tel;
formState.user.addr = store.userInfo.addr;
//表单提交事件
const onFinish = async () => {
  const userid = store.userInfo.id
  console.log(userid);
  const res = await setUserInfoApi(userid,formState.user)
  console.log(res);

  if(res.code == 200){
    
    store.setUserInfo(formState.user)
    message.success("修改个人信息成功")
  }else{
    message.error("修改个人信息失败")
  }
  
};
</script>

<style lang="scss">
.user{
  width: 80%;
  height: 100%;
   border: 2px solid rgba(0, 0, 0, 0.1);
  border-radius: 10px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  background-color: white;
  margin: 0 auto;
  display: flex;
  justify-content: center;
  align-items: center;
}

</style>