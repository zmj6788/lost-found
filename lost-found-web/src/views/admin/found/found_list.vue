<template>
  <div class="lf_container">
    <a-modal v-model:open="data.modalVisible" :title="data.updateID === 0 ? '添加物品' : '编辑物品'" @ok="handleOk">
      <a-form :model="formState" name="basic" ref="formRef" :label-col="{ span: 4 }" :wrapper-col="{ span: 20 }"
        autocomplete="off">
        <a-form-item label="捡到物品名" name="fount_name" has-feedback
          :rules="[{ required: true, message: '请输入物品名', trigger: 'blur' }]">
          <a-input v-model:value="formState.fount_name" placeholder="物品名" />
        </a-form-item>
        <a-form-item label="捡到地点" name="fount_place" has-feedback
          :rules="[{ required: true, message: '请输入捡到地点', trigger: 'blur' }]">
          <a-input v-model:value="formState.fount_place" placeholder="捡到地点" />
        </a-form-item>


        <a-form-item label="捡到时间" name="fount_time" has-feedback
          :rules="[{ required: true, message: '请输入捡到时间', trigger: 'blur' }]">
          <a-input-password v-model:value="formState.fount_time" placeholder="捡到时间" />
        </a-form-item>

        <a-form-item label="物品图" name="fount_picture" has-feedback
          :rules="[{ required: true, message: '请输入物品图', trigger: 'blur' }]">
          <a-input-password v-model:value="formState.fount_picture" placeholder="物品图" />
        </a-form-item>

        <a-form-item label="信息描述" name="fount_describe"
          :rules="[{ required: true, message: '请输入信息描述', trigger: 'blur' }]">
          <a-input-password v-model:value="formState.fount_describe" placeholder="信息描述" />

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
      <a-button type="primary" @click="userAdd">添加物品</a-button>
      <a-button type="primary" @click="removeBatch" danger v-if="data.selectedRowKeys.length">批量删除</a-button>
    </div>
    <div class="lf_tables">
      <a-spin :spinning="data.spinning" tip="加载中..." :delay="300">
        <a-table :columns="data.columns" :pagination="false" :data-source="data.list"
          :row-selection="{ selectedRowKeys: data.selectedRowKeys, onChange: onSelectChange }" row-key="id">
          <template #bodyCell="{ column, record }">
            <template v-if="column.key === 'fount_picture'">
              <img class="lf_user_avatar" :src="record.fount_picture" alt="">
            </template>
            <template v-if="column.key === 'action'">
              <a-button type="primary" @click="userEdit(record.id)">编辑</a-button>
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


//用户权限
const statusOptions = [
  {
    value: 1,
    label: "未找到",
  },
  {
    value: 2,
    label: "已找回",
  }
]


//表格数据渲染
const data = reactive({
  // 标签
  columns: [
    { title: 'id', dataIndex: 'id', key: 'id' },
    { title: '捡到物品名', dataIndex: 'fount_name', key: 'fount_name' },
    {
      title: '捡到地点',
      dataIndex: 'fount_place',
      key: 'fount_place',
    },
    {
      title: '捡到时间',
      dataIndex: 'fount_time',
      key: 'fount_time',
    },
    {
      title: '物品图',
      dataIndex: 'fount_picture',
      key: 'fount_picture',
    },
    {
      title: '信息描述',
      key: 'fount_describe',
      dataIndex: 'fount_describe',
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
      fount_name: "苹果15",
      fount_place: '中原区',
      fount_time: "2024--06--18",
      fount_picture: "https://www.apple.com.cn/iphone-15/images/overview/portrait/portrait_mode__wzbojol1j2a6_large_2x.jpg",
      fount_describe: "捡到了一部iphone15",
      status: 1,
    }
  ],
  count: 20,
  selectedRowKeys: [],
  modalVisible: false,
  updateID: 0,
  spinning: true,
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

//表单数据
const formState = reactive({
  fount_name: "",
  fount_place: "",
  fount_time: "",
  fount_picture:"",
  fount_describe: "",
  status: 1,
})
//用于清空表单
const _formState = {
  fount_name: "",
  fount_place: "",
  fount_time: "",
  fount_picture: "",
  fount_describe: "",
  status: 1,
}
//表单验证
const formRef = ref(null)
// 校验密码和确认密码是否一致
let validateRePassword = async (_rule, value) => {
  if (value === '') {
    return Promise.reject('Please input the password again');
  } else if (value !== formState.password) {
    return Promise.reject("Two inputs don't match!");
  } else {
    return Promise.resolve();
  }
};

//按钮功能配置
// 选择复选框
function onSelectChange(selectedKeys) {
  data.selectedRowKeys = selectedKeys
  console.log(selectedKeys)
}

//增加用户按钮点击事件
async function handleOk() {
  //首先验证表单数据
  try {
    await formRef.value.validate()
    //表单验证正确

    console.log('添加用户按钮点击事件')
    console.log(formState)
    data.modalVisible = false
    //清空表单
    Object.assign(formState, _formState)
    //重新获取数据
    getData()
    if (data.updateID === 0) {
      //添加用户

      message.success("添加用户成功")
    } else {
      //修改用户
      message.success("修改用户成功")
    }
    // 清除验证规则
    formRef.value.clearValidate()
  } catch (e) {
  }
}

// 添加用户按钮点击事件
function userAdd() {
  data.updateID = 0
  data.modalVisible = true
  console.log('添加用户按钮点击事件')
}
// 批量删除按钮点击事件
function removeBatch() {
  console.log('批量删除按钮点击事件')
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
// 修改单个用户
function userEdit(user_id) {
  data.updateID = user_id
  data.modalVisible = true
  console.log(user_id)
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