<template>
  <div class="lf_container">
    <a-modal v-model:open="data.modalVisible" :title="data.updateID === 0 ? '发布物品' : '编辑物品'" @ok="handleOk">
      <a-form :model="formState" name="basic" ref="formRef" :label-col="{ span: 4 }" :wrapper-col="{ span: 20 }"
        autocomplete="off">
        <a-form-item label="物品名" name="name" has-feedback
          :rules="[{ required: true, message: '请输入物品名', trigger: 'blur' }]">
          <a-input v-model:value="formState.name" placeholder="物品名" />
        </a-form-item>
        <a-form-item label="地点" name="place" has-feedback
          :rules="[{ required: true, message: '请输入地点', trigger: 'blur' }]">
          <a-input v-model:value="formState.place" placeholder="丢失地点" />
        </a-form-item>


        <a-form-item label="时间" name="time" has-feedback
          :rules="[{ required: true, message: '请输入时间', trigger: 'blur' }]">
          <a-input v-model:value="formState.time" placeholder="时间" />
        </a-form-item>


        <a-form-item label="物品图" name="picture" has-feedback
          >

          <a-upload v-model:file-list="data.fileList" :action="uploadAction" list-type="picture-card" name="images"
            :multiple="false" style="margin-top: 40px;">
            <i class="fa fa-cloud-upload"></i>
            <div style="margin-left: 5px">上传图片</div>
          </a-upload>
        </a-form-item>

        <a-form-item label="信息描述" name="describe" :rules="[{ required: true, message: '请输入信息描述', trigger: 'blur' }]">
          <a-input v-model:value="formState.describe" placeholder="信息描述" />

        </a-form-item>
        <a-form-item label="类型" name="types" :rules="[{ required: true, message: '请输入物品类型', trigger: 'blur' }]">
          <a-select v-model:value="formState.types" style="width: 200px" :options="typesOptions"
            placeholder="类型"></a-select>
        </a-form-item>
        <a-form-item label="状态" name="status" :rules="[{ required: true, message: '请输入物品状态', trigger: 'blur' }]">
          <a-select v-model:value="formState.status" style="width: 200px" :options="statusOptions"
            placeholder="状态"></a-select>
        </a-form-item>
      </a-form>
    </a-modal>




    <div class="lf_search">
      <a-input-search v-model:value="page.key" placeholder="请输入物品名" style="width: 200px" @search="onSearch" />
      <div class="lf_refresh">
        <a-button title="刷新本页" @click="refresh"><i class="fa fa-refresh"></i></a-button>
      </div>
    </div>

    <div class="lf_actions">
      <a-button type="primary" @click=" articleAdd">添加物品</a-button>
      <a-button type="primary" @click="removeBatch" danger v-if="data.selectedRowKeys.length">批量删除</a-button>
    </div>
    <div class="lf_tables">
      <a-spin :spinning="data.spinning" tip="加载中..." :delay="300">
        <a-table :columns="data.columns" :pagination="false" :data-source="data.list"
          :row-selection="{ selectedRowKeys: data.selectedRowKeys, onChange: onSelectChange }" row-key="id">
          <template #bodyCell="{ column, record }">
            <template v-if="column.key === 'picture'">
              <img class="lf_user_avatar" :src="record.picture" alt="">
            </template>
            <template v-if="column.key === 'types'">
              <!-- <i>{{ typesOptions[record.types - 1].label}}</i> -->
            </template>
            <template v-if="column.key === 'status'">
              <!-- <i>{{ statusOptions[record.types - 1].label}}</i> -->
            </template>
            <template v-if="column.key === 'action'">
              <a-button type="primary" @click=" articleEdit(record)">编辑</a-button>
              <a-popconfirm title="是否确定删除?" ok-text="删除" cancel-text="取消" @confirm=" articleRemove(record.id)">
                <a-button type="primary" danger>删除</a-button>
              </a-popconfirm>
            </template>

          </template>
        </a-table>
      </a-spin>
    </div>
    <div class="lf_pages">
      <a-pagination v-model:current="page.page" v-model:page-size="page.limit" :total="data.count"
        :show-total="total => `共 ${total} 条`" :showSizeChange="false" show-less-items @change="pageChange" />
      <br />
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue';
import { message } from 'ant-design-vue';
import { getArticlesInfoApi, addArticlesInfoApi, deleteArticlesInfoApi,updateArticlesInfoApi } from "@/api/article_api"
import { useStore } from '@/stores/store'

const store = useStore()
store.loadUserInfo()

// 动态设置action
const uploadAction = ref('/api/images/' + store.userInfo.id);

//物品状态
const statusOptions = [
  {
    value: 1,
    label: "未找到",
  },
  {
    value: 2,
    label: "已找到",
  }
]
//物品类型
const typesOptions = [
  {
    value: 1,
    label: "寻找失物",
  },
  {
    value: 2,
    label: "寻找失主",
  }
]


//表格数据渲染
const data = reactive({
  // 标签
  columns: [
    { title: 'id', dataIndex: 'id', key: 'id' },
    { title: '物品名', dataIndex: 'name', key: 'name' },
    {
      title: '地点',
      dataIndex: 'place',
      key: 'place',
    },
    {
      title: '时间',
      dataIndex: 'time',
      key: 'time',
    },
    {
      title: '物品图',
      dataIndex: 'picture',
      key: 'picture',
    },
    {
      title: '信息描述',
      key: 'describe',
      dataIndex: 'describe',
    },
    {
      title: '类型',
      key: 'types',
      dataIndex: 'types',
    },
    {
      title: '状态',
      key: 'status',
      dataIndex: 'status',
    },
    {
      title: '操作',
      key: 'action',
    },
  ],
  //数据
  list: [
    {
      id: 1,
      name: "苹果15",
      place: '中原区',
      time: "2024--06--18",
      picture: "https://www.apple.com.cn/iphone-15/images/overview/portrait/portrait_mode__wzbojol1j2a6_large_2x.jpg",
      describe: "奶奶滴真倒霉",
      types: 1,
      status: 1,
    }
  ],
  count: null,
  selectedRowKeys: [],
  modalVisible: false,
  updateID: 0,
  spinning: true,
  fileList:[]
})
const page = reactive({
  page: 1,
  limit: 5,
  key: "",
  sort: "created_at asc"
})


// 获取列表页数据
async function getData() {
  data.spinning = true
  const res = await getArticlesInfoApi(store.userInfo.id,page)
  data.count = res.data.count
  data.list = res.data.list
  data.spinning = false
}
//初始化数据
getData()
// 搜索
function onSearch() {
  //根据关键字查询
  page.key = page.key.trim()
  //
  // page.page = 1
  getData()
  message.success("搜索成功")
}
//刷新
function refresh() {
  message.success("刷新成功")
  data.spinning = true
  store.loadUserInfo()
  getData()
}
// 分页
function pageChange(_page, limit) {
  //页面切换后重新获取数据
  getData()
}

//表单数据
const formState = reactive({
  name: "",
  place: "",
  time: "",
  picture: "",
  describe: "",
  types: 1,
  status: 1,
})
//用于清空表单
const _formState = {
  name: "",
  place: "",
  time: "",
  picture: "",
  describe: "",
  types: 1,
  status: 1,
}
//表单验证
const formRef = ref(null)


// 选择复选框
//按钮功能配置用于批量删除
function onSelectChange(selectedKeys) {
  data.selectedRowKeys = selectedKeys
  console.log(selectedKeys)
}

//表单完成回调函数
async function handleOk() {
  //首先验证表单数据
  try {
    await formRef.value.validate()
    //表单验证正确

    //判断事件类型进行对应操作
    if (data.updateID === 0) {
      //添加物品
    
      formState.picture = data.fileList[0].response.data[0].file_name
      const res =  await addArticlesInfoApi(store.userInfo.id, formState)
      if (res.code !== 200) {
        message.error("添加物品失败")
        return
      }
      message.success("添加物品成功")
    } else {
      formState.picture = data.fileList[0].response.data[0].file_name
      let reqData = {
        id: data.updateID,
        name: formState.name,
        place: formState.place,
        time: formState.time,
        picture: formState.picture,
        describe: formState.describe,
        status: formState.status,
        types: formState.types,
      }
      //修改用户
      await updateArticlesInfoApi(store.userInfo.id, reqData)
      message.success("修改物品成功")
    }
    //关闭表单
    data.modalVisible = false
    //清空表单
    Object.assign(formState, _formState)
    //重新获取数据
    getData()
    
    // 清除验证规则
    formRef.value.clearValidate()
  } catch (e) {
  }
}


// 添加用户按钮点击事件
function  articleAdd() {
  //打开表单
  data.updateID = 0
  //清空表单
  Object.assign(formState, _formState)
  data.fileList = []
  data.modalVisible = true
}
// 批量删除按钮点击事件
async function removeBatch() {
  
    const ids = JSON.stringify({
      id_list: data.selectedRowKeys
    });

    const res = await deleteArticlesInfoApi(store.userInfo.id, ids)
    if (res.code == 200) {
      console.log(res);
      getData()
      message.success("删除物品成功")
    } else {
      console.log(res);
      message.error("删除物品失败")
    }
    location.reload()
}
// 删除单个用户
async function  articleRemove(articleid) {
  try {
    
    const ids = JSON.stringify({
      id_list: [articleid]
    });
      
    
    const res = await deleteArticlesInfoApi(store.userInfo.id, ids)
    if (res.code == 200) {
      console.log(res);
      getData()
      message.success("删除物品成功")
    }else{
      console.log(res);
      message.error("删除物品失败")
    }
  } catch (error) {
    console.error("删除文章时出错:", error);
  }
  location.reload()
}
// 修改单个用户
function  articleEdit(item) {
  data.updateID = item.id
  formState.name = item.name
  formState.place = item.place
  formState.time = item.time
  data.fileList = [{url: item.picture}]
  formState.describe = item.describe
  if (item.types === "寻找失物") {
    formState.types = 1
    
  } else {
    formState.types = 2
    
  }
  if (item.status === "未找到") {
    formState.status = 1
    
  } else {
    formState.status = 2
    
  }
  data.modalVisible = true
}

</script>

<style lang="scss">
.lf_container {

  background-color: white;

  padding: 20px;
  border: 2px solid rgba(0, 0, 0, 0.1);
  border-radius: 10px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);


  .lf_search {
    padding: 10px;
    border-bottom: 2px solid var(--bg);
  }

  .lf_refresh {
    position: absolute;
    right: 150px;
    top: 120px;

    i {
      color: var(--slide)
    }
  }

  .lf_actions {
    padding: 10px;

    .ant-btn {
      margin-right: 10px;
    }
  }

  .lf_tables {
    padding: 10px;

    .ant-btn {
      margin-right: 10px;
    }
  }

  .lf_pages {
    display: flex;
    justify-content: center;
    padding: 10px;
    margin-bottom: 10px;
  }

  .lf_user_avatar {
    width: 40px;
    height: 40px;
    border-radius: 50%;
  }
}
</style>