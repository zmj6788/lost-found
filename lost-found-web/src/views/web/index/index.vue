<template>
  
  <div class="home">
    <div class="slogan">
      <VueTyped v-if="data.abstract.length" :strings="data.abstract" :loop="true" :startDelay="300" :typeSpeed="100"
        :backSpeed="50">
        <span class="typing"></span>
      </VueTyped>
    </div>
      <a-carousel autoplay :autoplaySpeed="data.banner_time * 1000" class="lunbotu">
        <div class="dynamic_banner" v-for="item in data.banners" :key="item.id">
          <img :src="item.path" alt="">
        </div>
      </a-carousel>

  </div>
</template>

<script setup>
import VueTyped from 'vue3typed/libs/typed/index.vue';`   `
import { reactive } from "vue";
import img1 from '@/assets/images/a1.webp';
import img2 from '@/assets/images/a2.webp';
import img3 from '@/assets/images/a3.webp';
import { useStore } from '@/stores/store'
//保存userinfo
const store = useStore()
store.loadUserInfo()
const data = reactive({
  abstract: ['lost-found.com', '你值得信赖的网站！！！'],
  abstract_time: 6,
  banner_time: 3,
  banners: [
    {
      id: 1,
      path: img1
    },
    {
      id: 2,
      path: img2
    },
    {
      id: 3,
      path: img3
    }
  ],
})
</script>

<style lang="scss">
.home {
  display: flex;
  border: 5px solid rgba(0, 0, 0, 0.1);
  border-radius: 10px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    /* 将.home变为Flex容器 */
    flex-direction: column;
  width: 80%;
  height: 100%;
  background-color: white;
  margin: 0 auto;

  .slogan {
    font-size: 80px;
    text-align: center;
  }

  .lunbotu {
      height: 100%;
      .dynamic_banner {
       
          
    
        height: 510px;
        img {
          height:100%;
          object-fit: cover;
          display: block;
          width: 100%;
        }
      }
  }

}
</style>