<template>
  <div class="crudlist">
    <div class="crudlist-title">
      <div>
        <h1><span class="crudlist-title-red">go-vue-react</span></h1>

      </div>
   
      <div>
      <p  class="crudlist-title-go"># Go + Gin + Gorm + Mysql</p>
      <p  class="crudlist-title-vue"># Vue 3 + TypeScript + Vite + ElementPlus</p>
      <p  class="crudlist-title-react"># React18 + TypeScript +Sws +Vite + Material-UI</p>
      <p  class="crudlist-title-greenyellow1"># 联系作者请添加微信：fangdongdong_25，备注：go-vue-react</p>
      </div>
     
      
    </div>
    <!-- 查询 -->
    <div class="crudlist-query">
      <el-input @blur="search" v-model="name" placeholder="请输入姓名查询" style="width:200px" />
      <el-button type="primary" @click="handleClickEdit">新增</el-button>
    </div>


    <!-- 列表 -->
    <el-table :data="tableData" border style="width: 100%">
      <el-table-column prop="Name" label="姓名" />
      <el-table-column prop="age" label="年龄" />
      <el-table-column prop="school" label="学校" />
      <el-table-column prop="phone" label="手机号" />
      <el-table-column prop="address" label="地址" />
      <el-table-column fixed="right" label="操作" width="120">
        <template #default>
          <el-button link type="primary" size="small" @click="handleClickEdit">编辑</el-button>
          <el-button link type="primary" size="small" @click="deleteform">删除</el-button>
        </template>
      </el-table-column>
    </el-table>


    <!-- 分页 -->
    <el-pagination v-if="total > 0" class="crudlist-pagination" layout="prev, pager, next" :total="total" />

    <!-- 添加和编辑 -->
    <!-- <AddDialog ref="AddDialogRef"></AddDialog> -->

    <el-dialog v-model="centerDialogVisible" :title="title" width="50%" center>
      
      <el-form
        :label-position="labelPosition"
        label-width="100px"
        :model="form"
        style="max-width: 360px"
      >
        <el-form-item label="姓名">
          <el-input v-model="form.Name" placeholder="请输入姓名" />
        </el-form-item>
        <el-form-item label="年龄">
          <el-input v-model="form.age" placeholder="请输入年龄" />
        </el-form-item>
        <el-form-item label="学校">
          <el-input v-model="form.school" placeholder="请输入学校" />
        </el-form-item>
        <el-form-item label="手机号">
          <el-input v-model="form.phone" placeholder="请输入手机号" />
        </el-form-item>
        <el-form-item label="地址">
          <el-input v-model="form.address" placeholder="请输入地址" />
        </el-form-item>
    
      </el-form>
    
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="centerDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="summitform">
              确认
            </el-button> 
          </span>
        </template>
      </el-dialog>

  </div>
</template>

<!-- name    string
age          uint8
school       string
phone        string
address       string -->

<script setup lang="ts">
// 引入

// import AddDialog from './AddDialog.vue'

// const AddDialogRef = ref(null)

// // 注册子组件
import { defineAsyncComponent } from 'vue'
const AddDialog = defineAsyncComponent(() => import('./AddDialog.vue'))

import { reactive, ref } from 'vue'
import { listadd,listdelete ,listupdate,listquery} from "../../api/crudList";
import { ElMessage } from 'element-plus'








// data
const name = ref('')

const tableData = ref([
  {
    Name: '',
  age: '',
  school: '',
  phone: '',
  address: '',
  }

])


const total = ref(0)
const page = ref(1)
const pageSize = ref(999)
const labelPosition = ref('right')

const form = reactive({
  name: '',
  age: '',
  school: '',
  phone: '',
  address: '',
})

const centerDialogVisible = ref(false)
const title = ref('新增')

// method
// 新增编辑弹窗
const handleClickEdit = () => {
  centerDialogVisible.value = true
}




// 新增
const addform=async()=>{
  const { code }= await listadd({ 
       ...form,
       age:parseInt(form.age)
       }
    );

    // console.log("res===",res);
    

    if (code==200) {
      centerDialogVisible.value = false
      ElMessage({
        type: 'success',
        message: '添加成功',
        showClose: true,
      })
      search()

      
    }
   


}

// 删除
const deleteform=async()=>{
  const { code } = await listdelete({ 
       ...form,
       age:parseInt(form.age)
       }
    );
   


}


    // 编辑
 const updateform=async()=>{
  
  const { code } = await listupdate({ 
      //  ...form,
       age:parseInt(form.age)
       }
    );
   

}




// 查询
const search=async()=>{
  const { code,data } = await listquery({ 
       name:name.value,
       page:page.value,
       pageSize :pageSize .value

       }
    );
    if (code==200) {
      console.log("data==",data, tableData.value );
      
      tableData.value ==data
      
    }
   

}

search()

// 新增、编辑
const summitform = () => {
  addform()
  // updateform()
}






</script>

<style scoped>
.crudlist {
  margin: 20px;
}

.crudlist-query {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.crudlist-pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 20px;
}
.crudlist-title{
  display: flex;
   flex-direction: column;

  justify-content: center;
  align-items: center;

}
.crudlist-title-red{
  color:red;
}
.crudlist-title-go{
  color:gold;

}
.crudlist-title-vue{
  color:green;

}
.crudlist-title-react{
  color:pink;

}
</style>