<template>
  <div class="crudlist">
    <!-- 查询 -->
    <div class="crudlist-query">
      <el-input v-model="input" placeholder="Please input" style="width:200px" />
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
        <template #default>
          <el-button link type="primary" size="small" @click="handleClickEdit">编辑</el-button>
        </template>
      </el-table-column>
    </el-table>


    <!-- 分页 -->
    <el-pagination v-if="total > 0" class="crudlist-pagination" layout="prev, pager, next" :total="total" />

    <!-- 添加和编辑 -->
    <AddDialog ref="AddDialogRef"></AddDialog>

  </div>
</template>

<!-- name    string
age          uint8
school       string
phone        string
address       string -->

<script setup lang="ts">
// 引入
import { ref } from 'vue'
import AddDialog from './AddDialog.vue'

// data
const input = ref('')
const total = 100

const tableData = [
  {
    date: '2016-05-03',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-02',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-04',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-01',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
]

// method
const handleClickEdit = () => {
  console.log('click', AddDialogRef)
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
</style>