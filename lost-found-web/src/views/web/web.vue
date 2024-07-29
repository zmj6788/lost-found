<template>

  <div class="lf_web">
    <div class="main">

      <header>
        <LFWebHead />
      </header>

      <main>
        <RouterView />
      </main>

    </div>
  </div>

</template>
<script setup>
import LFWebHead from '../../components/web/lf_web_head.vue'









</script>

<style lang="scss">
.lf_web {
  width: 100%;
  height: 100vh;
  
  
  .main {
    height: 100vh;
    display: flex;
    flex-direction: column;
    
    
    header {
      width: 100%;
      height: 90px;
      background-color: #57606f;
      color: white;
      display: flex;
      justify-content: space-between;
      align-items: center;
      
      
    }
    
    main {
      flex-grow: 1;
      width: 100%;
      height: calc(100% - 90px);
      background-color: white;
    }

  }
}
</style>