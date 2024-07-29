
import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import Antd from 'ant-design-vue';
import './assets/css/public.css'
import 'ant-design-vue/dist/reset.css';
import "font-awesome/css/font-awesome.min.css";
import "./assets/css/iconfont.css";
import "./assets/css/theme.css";
import Echarts from "vue-echarts";
import * as echarts from "echarts";


const app = createApp(App)

app.component("v-chart", Echarts)

app.config.globalProperties.$echarts = echarts

app.use(createPinia())
app.use(router)
app.use(Antd)

app.mount('#app')
