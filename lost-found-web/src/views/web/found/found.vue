<template>
  <div class="lost_card">
    <a-card title="失物招领物品栏" style="background-color: white;">
      <a-card-grid v-for="(item, index) in cards" @click="showDetails(item.id)" :key="index"
        style="width: 25%; text-align: center;">
        <a-card hoverable style="width: 240px; height: 240px;">
          <template #cover>
            <img :alt="item.name" :src="item.picture" style="width: 240px; height: auto;" />
          </template>
          <a-card-meta :title="'物品：' + item.name" style="padding: 10px; margin-top: 20px">
            <template #description>地点：{{ item.place }}</template>
            <p>类型:{{ item.types }}</p>
          </a-card-meta>
        </a-card>
      </a-card-grid>
    </a-card>



  </div>
</template>

<script setup>

import { reactive } from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from '@/stores/store';
import { getArticlesInfoApi } from '@/api/article_api.js';
//加载userinfo
const store = useStore()
store.loadUserInfo()

const router = useRouter();
// 定义卡片数据数组
const cards = reactive([
  {
    id: 1,
    picture: 'https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png',
    name: '物品名1',
    place: '地点1',
  }
]);
const page = {
  page: 0,
  limit: 0,
  key: "",
  sort: "created_at asc"
}
const userid = store.userInfo.id
//获取数据
const getArticlesInfo = async () => {

  const res = await getArticlesInfoApi(userid, page)
console.log(res);
  cards.splice(0, cards.length, ...res.data.list.map(item => ({
    id: item.id,
    picture: item.picture,
    name: item.name,
    place: item.place,
  })));

}
getArticlesInfo()
//卡片点击事件


// 定义方法在卡片被点击时展示详情
const showDetails = (id) => {

  router.push({ name: 'show', params: { id} });
};



</script>

<style>
.lost_card {
  border: 2px solid rgba(0, 0, 0, 0.1);
  border-radius: 10px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  width: 80%;
  background-color: #a4b0be;
  margin: 0 auto;
}
</style>