<template>
  <div class="show">

    <a-card title="物品详情" style="width: 1000px; height: 600px; background-color: white; margin-top: 20px;">
      <a-card-grid style="width: 40%; text-align: center">
        <a-card hoverable style="width: 240px">
          <img :alt="data.name" :src="data.picture" />
          <a-card-meta :title="'物品: ' + data.name" style="padding: 10px;">
            <template #description>丢失地点: <span> {{ data.place }}</span></template>
          </a-card-meta>
        </a-card>
      </a-card-grid>
      <a-card-grid style="width: 60%; text-align: center" :hoverable="false">
        <a-card title="物品信息描述" :bordered="false" style="width: 300px; background-color: white; padding: 30px">
          <p>{{ data.describe }}</p>
        </a-card>
      </a-card-grid>
      <a-card-grid style="width: 1; text-align: center">联系人<span>{{ data.user}}</span></a-card-grid>
      <a-card-grid style="width: 1; text-align: center">电话号码<span>{{ data.tel }}</span></a-card-grid>
      <a-card-grid style="width: 1; text-align: center">时间<span>{{data.time}}</span></a-card-grid>
    </a-card>

  </div>
</template>

<script setup>
import { reactive } from 'vue';
import { useRoute } from 'vue-router';
import { useStore } from '@/stores/store';
import { getArticlesInfoApi } from '@/api/article_api.js';
import { getUserInfoApi } from '@/api/user_api';
//保存userinfo
const store = useStore()
store.loadUserInfo()
const route = useRoute();
const id = parseInt(route.params.id);

const data = reactive({
  name: 'iphone15',
  place: '操场',
  time: '2024-06-18',
  picture: 'https://b0.bdstatic.com/ugc/8LmC1A5bDjGqJhbZZiOaSw3f3642cd4e68de16093bd7f387aa8b01.jpg@h_1280',
  describe: '我怎么这么倒霉呀，今天在操场溜达，中午去吃饭，发现手机没了，打饭阿姨，看着我的那个尴尬呀，蓝瘦香菇',
  status: 1,
  tel:"",
  user:""
})

//根据id去查询对应的数据
const userid = store.userInfo.id
const page = {
  page: 0,
  limit: 0,
  key: "",
  sort: "created_at asc",
  id: id
}
const getArticledata = async () => {
  console.log(userid, page);
  const res = await getArticlesInfoApi(userid, page)
  console.log(res.data.list[0]);
  data.name = res.data.list[0].name
  data.place = res.data.list[0].place
  data.time = res.data.list[0].time
  data.picture = res.data.list[0].picture
  data.describe = res.data.list[0].describe
  data.status = res.data.list[0].status
 let  id = res.data.list[0].user_id
  const userdata = await getUserInfoApi(id)
  console.log(userdata);
  data.tel = userdata.data.tel
  data.user = userdata.data.nick_name

}
getArticledata()
</script>

<style lang="scss">
.show {
  flex-grow: 1;
  width: 80%;
  height: 100%;
  background-color: white;
  border: 2px solid rgba(0, 0, 0, 0.1);
  border-radius: 10px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  margin: 0 auto;
  display: flex;
  justify-content: center;
  align-items: center;

  img {
    width: 200px;
    height: 300px;
  }

  span {
    color: pink;
    font-size: 20px;
    margin-left: 15px;
  }

  p {
    margin-top: 60px;
    // color: white;
    font-size: 20px;
  }
}
</style>
