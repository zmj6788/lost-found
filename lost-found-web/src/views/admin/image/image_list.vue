<template>
  <div class="lf_container">
    <a-modal title="上传图片" v-model:open="data.visible" @ok="uploadComplete">
      <a-upload v-model:file-list="data.fileList" action="/api/images" list-type="picture-card" name="images" multiple
        :headers="{ token: store.userInfo.token }" style="margin-top: 40px;">
        <i class="fa fa-cloud-upload"></i>
        <div style="margin-left: 5px">上传图片</div>
      </a-upload>
    </a-modal>

    <a-modal title="编辑图片" v-model:open="data.modalVisible" @ok="update(updateState.id)">
      <a-form :model="updateState" name="basic" ref="formRef" :label-col="{ span: 4 }" :wrapper-col="{ span: 20 }"
        autocomplete="off">
        <a-form-item label="文件名称" name="name" has-feedback
          :rules="[{ required: true, message: '请输入文件名', trigger: 'blur' }]">
          <a-input v-model:value="updateState.name" placeholder="文件名" />
        </a-form-item>
        <a-form-item label="图片预览">
          <img :src="updateState.path" height="80" style="border-radius: 5px" alt="">
        </a-form-item>
      </a-form>
    </a-modal>

    <div class="lf_search">
      <a-input-search v-model:value="page.key" placeholder="请输入图片名" style="width: 200px" @search="onSearch" />
      <div class="lf_refresh">
        <a-button title="刷新本页" @click="refresh"><i class="fa fa-refresh"></i></a-button>
      </div>
    </div>

    <div class="lf_actions">
      <a-button type="primary" @click="data.visible = true">上传图片</a-button>
      <a-button type="primary" @click="removeBatch" danger v-if="data.selectedRowKeys.length">批量删除</a-button>
    </div>
    <div class="lf_tables">
      <a-spin :spinning="data.spinning" tip="加载中..." :delay="300">
        <a-table :columns="data.columns" :pagination="false" :data-source="data.list"
          :row-selection="{selectedRowKeys: data.selectedRowKeys,onChange: onSelectChange }" row-key="id">
          <template #bodyCell="{ column, record }">
            <template v-if="column.key === 'img_show'">
              <img class="lf_user_avatar" :src="record.img_show" alt="">
            </template>

            <template v-if="column.key === 'action'">
              <a-button type="primary" @click="data.modalVisible = true">编辑</a-button>
              <a-popconfirm title="是否确定删除?" ok-text="删除" cancel-text="取消" @confirm="userRemove(record.id)">
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
import { useStore } from "@/stores/store";

const store = useStore()
const formRef = ref(null)
//用户权限



//表格数据渲染
const data = reactive({
  // 标签
  columns: [
    { title: 'id', dataIndex: 'id', key: 'id' },
    { title: '图片名', dataIndex: 'img_name', key: 'img_name' },
    {
      title: '图片路径',
      dataIndex: 'img_path',
      key: 'img_path',
    },
    {
      title: '图片预览',
      dataIndex: 'img_show',
      key: 'img_show',
    },
    {
      title: '操作',
      key: 'action',
    },
  ],
  //表格渲染数据
  list: [
    {
      id: 0,
      img_name: 'a.png',
      img_path: '/upload/image/a.png',
      img_show: "https://img0.baidu.com/it/u=4054616240,2486283051&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500",
    },
    {
      id: 1,
      img_name: 'a.png',
      img_path: '/upload/image/a.png',
      img_show: "https://img0.baidu.com/it/u=4054616240,2486283051&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500",
    }
  ],
  count: 20,
  selectedRowKeys: [],
  modalVisible: false,
  visible: false,
  spinning: true,
  fileList: []
})
//表单输入数据
const updateState = reactive({
  id: 0,
  name: "",
  path: data.list[0].img_show
})
const page = reactive({
  page: 1,
  limit: 5,
  key: "",
})

// 获取列表页数据
async function getData() {
  data.spinning = true
  // let res = await userListApi(page)
  // data.list = res.data.list
  // data.count = res.data.count
  data.spinning = false
}
getData()
// 搜索
function onSearch() {
  page.key = page.key.trim()
  //
  page.page = 1
  console.log(page.key);
  getData()
  message.success("搜索成功")
}
//刷新
function refresh() {
  message.success("刷新成功")
  data.spinning = true
  
  // location.reload()
  getData()
}
// 分页
function pageChange(_page, limit) {
  getData()
}


//按钮功能配置
// 选择复选框
function onSelectChange(selectedKeys) {
  data.selectedRowKeys = selectedKeys
  console.log(selectedKeys)
}

//上传图片按钮点击事件
function uploadComplete() {
  //首先验证表单数据
  try {
    // await formRef.value.validate()
    //表单验证正确

    console.log('添加用户按钮点击事件')
    
    
    //重新获取数据
    
    message.success("添加用户成功")
    // 清除验证规则
    data.visible = false
    // formRef.value.clearValidate()
  } catch (e) {
  }
}


// 批量删除按钮点击事件
function removeBatch() {
  console.log('批量删除按钮点击事件')
  console.log(data.selectedRowKeys)
  getData()
  message.success("移除多个用户成功")
}
// 删除单个用户
function userRemove(user_id) {
  console.log('删除单个用户按钮点击事件')
  console.log(user_id)
  getData()
  message.success("移除单个用户成功")
}
// 修改图片名
function update(id) {
  
  data.modalVisible = false
  console.log(id)
}


</script>

<style lang="scss">
.lf_container {
  background-color: white;
  padding: 20px;
  

  .lf_search {
    padding: 10px;
    border-bottom: 2px solid var(--bg);
  }
  .lf_refresh{
    position: absolute;
    right: 150px;
    top: 120px;
    i{
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
  .lf_user_avatar{
    width: 40px;
    height: 40px;
    border-radius: 50%;
  }
}
</style>