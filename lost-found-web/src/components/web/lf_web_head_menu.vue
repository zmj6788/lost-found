<template>
  <a-dropdown>
    <a class="ant-dropdown-link" @click.prevent>
      {{ store.userInfo.nick_name }}
      <i class="fa fa-chevron-down"></i>
    </a>
    <template #overlay>
      <a-menu @click="actions" style="background-color: #a4b0be;">
        <a-menu-item key="article">
          <a href="javascript:;">发布物品</a>
        </a-menu-item>
        <a-menu-item key="user">
          <a href="javascript:;">个人信息</a>
        </a-menu-item>
        <a-menu-item key="logout">
          <a href="javascript:;">退出登录</a>
        </a-menu-item>
      </a-menu>
    </template>
  </a-dropdown>

</template>

<script setup>
import { useRouter } from 'vue-router';
import { useStore } from '@/stores/store'

//默认用户名显示
const store = useStore()
//功能匹配
const router = useRouter();



const actions = (key) => {
  if (key.key === 'logout') {
    //退出登录
    //清空userid,清空后则不能够进行业务逻辑
    localStorage.clear("user")
    router.push('/')

    return
  }
 
  router.push(`/web/${key.key}`)
  // console.log(key.key);
};
</script>

<style></style>