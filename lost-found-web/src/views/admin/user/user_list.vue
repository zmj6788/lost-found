<template>
  <div class="lf_container">
    <a-modal v-model:open="data.modalVisible" :title="data.updateID === 0 ? '添加用户' : '编辑用户'" @ok="handleOk">
      <a-form :model="formState" name="basic" ref="formRef" :label-col="{ span: 4 }" :wrapper-col="{ span: 20 }"
        autocomplete="off">
        <a-form-item label="电话号码" name="phone" has-feedback
          :rules="[{ required: true, message: '请输入电话号码', trigger: 'blur' }]">
          <a-input v-model:value="formState.phone" placeholder="电话号码" />
        </a-form-item>
        <a-form-item label="用户名" name="user_name" has-feedback
          :rules="[{ required: true, message: '请输入用户名', trigger: 'blur' }]">
          <a-input v-model:value="formState.user_name" placeholder="用户名" />
        </a-form-item>


        <a-form-item label="密码" name="password" has-feedback
          :rules="[{ required: true, message: '请输入密码', trigger: 'blur' }]">
          <a-input-password v-model:value="formState.password" placeholder="密码" />
        </a-form-item>

        <a-form-item label="确认密码" name="re_password" has-feedback
          :rules="[{ required: true, message: '请输入确认密码' }, { validator: validateRePassword, message: '两次密码不一致', trigger: 'blur' }]">
          <a-input-password v-model:value="formState.re_password" placeholder="确认密码" />
        </a-form-item>

        <a-form-item label="权限" name="role" :rules="[{ required: true, message: '请选择用户权限', trigger: 'blur' }]">
          <a-select v-model:value="formState.role" style="width: 200px" :options="roleOptions"
            placeholder="选择权限"></a-select>
        </a-form-item>
      </a-form>
    </a-modal>




    <div class="lf_search">
      <a-input-search v-model:value="page.key" placeholder="请输入用户名" style="width: 200px" @search="onSearch" />
      <div class="lf_refresh">
        <a-button title="刷新本页" @click="refresh"><i class="fa fa-refresh"></i></a-button>
      </div>
    </div>

    <div class="lf_actions">
      <a-button type="primary" @click="userAdd">添加用户</a-button>
      <a-button type="primary" @click="removeBatch" danger v-if="data.selectedRowKeys.length">批量删除</a-button>
    </div>
    <div class="lf_tables">
      <a-spin :spinning="data.spinning" tip="加载中..." :delay="300">
        <a-table :columns="data.columns" :pagination="false" :data-source="data.list"
          :row-selection="{selectedRowKeys: data.selectedRowKeys,onChange: onSelectChange }" row-key="id">
          <template #bodyCell="{ column, record }">
            <template v-if="column.key === 'avatar'">
              <img class="lf_user_avatar" :src="record.avatar" alt="">
            </template>
            
            <template v-if="column.key === 'action'">
              <a-button type="primary" @click="userEdit(record)">编辑</a-button>
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
const roleOptions =  [
  {
    value: 1,
    label: "普通用户",
  },
   { 
    value: 2,
    label: "管理员",
  },
   { 
    value: 3,
    label: "游客",
  }

  
]


//表格数据渲染
const data = reactive({
  // 标签
  columns: [
    { title: 'id', dataIndex: 'id', key: 'id' },
    { title: '头像', dataIndex: 'avatar', key: 'avatar' },
    {
      title: '用户名',
      dataIndex: 'user_name',
      key: 'user_name',
    },
    {
      title: '密码',
      dataIndex: 'password',
      key: 'password',
    },
    {
      title: '权限',
      dataIndex: 'role',
      key: 'role',
    },
    {
      title: '电话号码',
      key: 'phone',
      dataIndex: 'phone',
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
      avatar: "https://img0.baidu.com/it/u=4054616240,2486283051&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500",
      user_name: 'John Brown',
      password: 32,
      role: 1,
      phone: 110
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
  setTimeout(() => {

    data.spinning = false                                      
  }, 100)
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
  user_name: "",
  phone: "",
  password: "",
  re_password: "",
  role: "",
})
//用于清空表单
const _formState = {
  user_name: "",
  phone: "",
  password: "",
  re_password: "",
  role: 2,
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
async  function handleOk() {
  //首先验证表单数据
  try {
    await formRef.value.validate()
    //表单验证正确
    //添加前清空编辑表单
    //清空表单
    Object.assign(formState, _formState)
    console.log('添加用户按钮点击事件')
    console.log(formState)
    data.modalVisible = false
    //添加前清空添加表单
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
function userEdit(record) {
  //填充表单用于编辑
  formState.user_name = record.user_name
  formState.phone = record.phone
  formState.password = record.password
  formState.re_password = record.password
  formState.role = record.role
  
  data.updateID = record.id
  data.modalVisible = true
  
  console.log(record.id)
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