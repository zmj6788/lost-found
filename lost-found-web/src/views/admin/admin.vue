<template>
  <div class="lf_admin">
    <LFAside></LFAside>

    <div class="main">
      <header>
        <div class="left">
          <div class="logo">
            <div>失物找回系统</div>
            <div>lost-found</div>
          </div>
        </div>
        <div class="rignt">
          <div class="actions">
            <div class="icon_actions">
              <i class="fa fa-home"></i>
              <LFTheme></LFTheme>

              <LFquanping></LFquanping>
            </div>
            <div class="avatar">
              <img src="https://img0.baidu.com/it/u=4054616240,2486283051&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500"
                alt="">
            </div>
            <div class="drop_menu">
              <a-dropdown>
                <a class="ant-dropdown-link" @click.prevent>
                  管理员
                  <i class="fa fa-angle-down"></i>
                </a>
                <template #overlay>
                  <a-menu @click="menuDrop">
                    <a-menu-item key="logout">
                      <a href="javascript:;">注销退出</a>
                    </a-menu-item>
                  </a-menu>
                </template>
              </a-dropdown>
            </div>
          </div>

        </div>
      </header>
      <div class="tabs"></div>
      <main>
        <router-view/>
      </main>
    </div>
  </div>
</template>

<script setup>
import LFTheme from '../../components/lf_theme.vue'
import LFAside from '../../components/admin/lf_aside.vue'
import LFquanping from '../../components/admin/lf_quanping.vue'

import { useRouter } from 'vue-router'
//下拉菜单点击事件
const router = useRouter()
const menuDrop = key => {
  if (key.key === 'logout') {
    router.push('/')
    localStorage.clear("user")
  }
}





</script>

<style lang="scss">
.lf_admin {
  width: 100%;
  display: flex;
}


.main {
  width: calc(100% - 240px);
  height:100vh;
  overflow-y: auto;

  header {
    height: 60px;
    background-color: var(--hd);
    padding: 0 20px;
    display: flex;
    justify-content: space-between;
    align-items: center;
      
     .logo {
    width: 200%;
    transform: scale(0.7);
    font-family: "宋体", sans-serif;
    margin-top: 10px;
    margin-bottom: 10px;
    display: flex; /* 使用Flexbox布局 */
    flex-direction: column; /* 设置为垂直堆叠 */
    align-items: center; /* 水平居中 */
    justify-content: center; /* 垂直居中，假设.logo容器本身需要垂直居中于其父容器，否则可以省略这条 */
}

.logo > div:first-child {
    font-size: 26px;
    margin-bottom: 5px; /* 可选，添加间距 */
}

.logo > div:last-child {
    font-size: 30px;
}
    
    .right {}

    .actions {
      align-items: center;
      display: flex;

    }

    .icon_actions {
      margin-right: 20px;

      i {
        margin-right: 10px;
        font-size: 20px;
        color: var(--text);
      }

      i:hover {
        color: var(--active);
      }
    }

    .avatar {
      img {
        width: 40px;
        height: 40px;
        border-radius: 50%;
      }
    }

    .drop_menu {
      margin-left: 10px;
    }
  }

  .tabs {
    height: 30px;
    border: 1px solid #f0eeee;
    background-color: var(--order);
  }

  main {
    background-color: var(--bg);
    height: calc(100vh - 90px);
  }
}
</style>