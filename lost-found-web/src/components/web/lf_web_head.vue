<template>
  <div class="logo">
    <div>失物找回系统</div>
    <div>lost-found</div>
  </div>
  <div class="left">
    <a-breadcrumb separator="" style="color: white; font-size: 15px;">
      <a-breadcrumb-item @click="home">首页</a-breadcrumb-item>
      <a-breadcrumb-item class="custom-style" @click="found">失物招领</a-breadcrumb-item>
    </a-breadcrumb>
  </div>
  <div class="right">
    <img :src="store.userInfo.avatar_id" alt="默认头像" title="默认头像">
    <LFWebHeadMenu />
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router';
import LFWebHeadMenu from '../../components/web/lf_web_head_menu.vue'
import { useStore } from '@/stores/store'

//默认头像显示
const store = useStore()
const router = useRouter();

//路由匹配模块
const home = () => {
  router.push('/web/index')
};

const found = () => {
  router.push('/web/found')
};


</script>

<style>

.logo {
  width: 40%;
  transform: scale(0.7);
  font-family: "宋体", sans-serif;
  margin-top: 10px;
  margin-bottom: 10px;
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
  font-size: 26px;
  margin-bottom: 5px;
  /* 可选，添加间距 */
}

.logo>div:last-child {
  font-size: 30px;
}

.left {
  width: 50%;
  display: flex;
  margin-left: 40px;

  span {
    margin-right: 30px;
  }

  .search i {
    cursor: pointer;
  }
  .custom-style{
    color: white !important;
    font-size: 15px;
  }
}

.right {
  width: 40%;
  margin-right: 100px;
  display: flex;
  justify-content: right;
  align-items: center;

  img {
    width: 40px;
    height: 40px;
    border-radius: 50%;
    margin-right: 20px;
  }
}
</style>