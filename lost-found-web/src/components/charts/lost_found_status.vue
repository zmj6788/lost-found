<template>
  <div id="lost_found" style="height: 280px"></div>
</template>

<script setup>
import { onMounted } from "vue";
import * as echarts from "echarts"

//需要传递的数据  data
function lost_found() {

  var chartDom = document.getElementById('lost_found');
  var myChart = echarts.init(chartDom);
  var option;

  option = {
    tooltip: {
      trigger: 'item'
    },
    legend: {
      top: '5%',
      left: 'center'
    },
    series: [
      {
        name: 'Access From',
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        padAngle: 5,
        itemStyle: {
          borderRadius: 10
        },
        label: {
          show: false,
          position: 'center'
        },
        emphasis: {
          label: {
            show: true,
            fontSize: 40,
            fontWeight: 'bold'
          }
        },
        labelLine: {
          show: false
        },
        data: [
          { value: 90, name: '丢失物品' },
          { value: 45, name: '已找回物品' },
          { value: 20, name: '寻找失主' },
          { value: 15, name: '已找到失主' },
        ]
      }
    ]
  };

  option && myChart.setOption(option);

}
onMounted(() => {
  lost_found()
})
</script>
