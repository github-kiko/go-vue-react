<template>
  <div class="crudlist">
    <div class="crudlist-title">
      <div>
        <h1><span style="margin-right:10px" >go-vue-react</span>全栈小项目</h1>
      

      </div>
   
      <div>

      
      <p class="crudlist-title-red"># 用市场上最新的主流技术从0到1搭建一个全栈小项目</p>
      <p class="crudlist-title-red"># 后端使用Go+Mysql，前端做了两套，分别使用了Vue 3 和React 18 两种不同的技术栈</p>
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
      <el-table-column prop="name" label="姓名" />
      <el-table-column prop="age" label="年龄" />
      <el-table-column prop="school" label="学校" />
      <el-table-column prop="phone" label="手机号" />
      <el-table-column prop="address" label="地址" />
      <el-table-column fixed="right" label="操作" width="120">
        <template #default="scope">
          <el-button link type="primary" size="small" @click="handleClickEdit(scope.row)">编辑</el-button>
          <el-button link type="primary" size="small" @click="deleteform(scope.row)">删除</el-button>
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
          <el-input v-model="form.name" placeholder="请输入姓名" />
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
import { reactive, ref } from 'vue'
import { listadd,listdelete,listupdate,listquery} from "../../api/crudList";
import { ElMessage } from 'element-plus'
const AddDialog = defineAsyncComponent(() => import('./AddDialog.vue'))










// data
const name = ref('')

const tableData = ref()


const total = ref(0)
const page = ref(1)
const pageSize = ref(999)
const labelPosition = ref('right')

// interface form {
//   name: string;
//   id: number;
//   age:number;
//   school: string;
//   phone: string;
//   address: string;
// }

const form:any = ref({
  name:'',
  id:0,
  age:0,
  school:'',
  phone:'',
  address:'',

})



const centerDialogVisible = ref(false)
const title = ref('新增')

// method
// 新增编辑弹窗
const handleClickEdit = (row:any) => {
  centerDialogVisible.value = true
  form.value=JSON.parse(JSON.stringify(row))


  

}









// 新增
const addform=async()=>{
  const { code }= await listadd({ 
       ...form.value,
       age:parseInt(form.value.age)
       }
    );


    

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
const deleteform=async(row:any)=>{
  const { code } = await listdelete({ 
       id:row.id
       }
    );
    ElMessage({
        type: 'success',
        message: '删除成功',
        showClose: true,
      })
      search()
   


}


    // 编辑
 const updateform=async()=>{

  
  
  const { code } = await listupdate({ 
       ...form.value,
       age:parseInt(form.value.age)
       }
    );
    if (code==200) {
      centerDialogVisible.value = false
      ElMessage({
        type: 'success',
        message: '修改成功',
        showClose: true,
      })
      search()

      
    }
   

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
      tableData.value=data
      
    }
   

}

search()

// 新增、编辑
const summitform = () => {
  if (!form.value.id) {
    addform()
    
  }else{
      updateform()

  }


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