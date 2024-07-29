<template>
  <aside class="lf_aside">
    <div class="lf_aside_header flex">
      <div class="lf_aside_logo flex">
        <img src="https://img0.baidu.com/it/u=4054616240,2486283051&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500" alt="">
      </div>
      <div class="lf_aside_title">
        失物找回系统后台
      </div>
    </div>
    <div class="lf_aside_body flex">
      <div>
        <LFAsidemenu></LFAsidemenu>
      </div>
    </div>
  </aside>

</template>

<script setup>
import LFAsidemenu from '../../components/admin/if_aside_menu.vue'

</script>

<style lang="scss">
.lf_aside {
  width: 240px;
  height: 100vh;
  background-color: var(--slide);
  

  .lf_aside_header {
    width: 100%;
    height: 180px;
    flex-direction: column;
  }

  .lf_aside_logo {
    padding: 10px;
    img{
      margin-top: 40px;
      width: 80px;
      height: 80px;
      border-radius: 10%;
    }
  }
  .lf_aside_title {
    margin-top: 20px;
    font-size: 20px;
    color:white;
    
    
  }
    .lf_aside_body{
      width: 100%;
      margin-top: 50px;
    }

}
</style>