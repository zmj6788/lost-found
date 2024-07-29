<template>
  <div class="home">
    <div class="tab">
      <a-card title="失物找回系统使用状况">
        <a-card-grid style="width: 25%; text-align: center">用户 <span> 20 </span></a-card-grid>
        <a-card-grid style="width: 25%; text-align: center">发布失物 <span> 20 </span></a-card-grid>
        <a-card-grid style="width: 25%; text-align: center">寻找失主 <span> 20 </span></a-card-grid>
        <a-card-grid style="width: 25%; text-align: center">成功帮助 <span> 20 </span>人</a-card-grid>
      </a-card>
    </div>
    <div class="echart">
      <div class="left">

        <div class="user_use">
          <div class="title_left">
            用户使用趋势
          </div>
          <div id="user_use">
            <UserCharts />
          </div>
        </div>
      </div>
      <div class="right">

        <div class="lost_found_status">
          <div class="title_right">
            物品状态
          </div>
          <div id="lost_found_status">
            <LostFoundCharts />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import UserCharts from '../../../components/charts/user_use.vue'
import LostFoundCharts from '../../../components/charts/lost_found_status.vue'

</script>

<style lang="scss">
.home {
  background-color: white;
  padding: 20px;

  .tab{
    height: auto;
    margin-bottom: 10px;
    span{
      padding: 0 5px;
      font-size: 20px;
      color: pink;
    }
  }

  .echart {
    display: flex;

    .left {
      width: 630px;
      margin-right: 40px;
      .user_use {

        padding: 10px 20px;
        border: 2px solid var(--bg);
        border-radius: 5px;
        background-color: #f5f5f5;

        .title_left {
          display: flex;
          justify-content: center;
          align-items: center;
          margin-top: 20px;
          margin-bottom: 40px;
        }
      }
    }

    .right {
      width: calc(100% - 780px);

      .lost_found_status {
        padding: 10px 20px;
        border: 2px solid var(--bg);
        border-radius: 5px;
        background-color: #f5f5f5;
      }

      .title_right {
        display: flex;
        justify-content: center;
        align-items: center;
        margin-top: 20px;
        margin-bottom: 40px;
      }

    }
  }

}
</style>