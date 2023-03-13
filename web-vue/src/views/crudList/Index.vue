<template>
    <div class="crudlist">
      <div class="crudlist-query">          
      <el-input v-model="input" placeholder="Please input" />
      <el-button type="primary" >-22333-新增11</el-button>
    </div>
 
       

        <el-table :data="tableData" border style="width: 100%">
    <el-table-column prop="name" label="姓名"  />
    <el-table-column prop="age" label="年龄"  />
    <el-table-column prop="school" label="学校"  />
    <el-table-column prop="phone" label="手机号"  />
    <el-table-column prop="address" label="地址" />
  </el-table>

  <el-pagination layout="prev, pager, next" :total="total" />

    </div>
</template>

<!-- name    string
age          uint8
school       string
phone        string
address       string -->

<script setup lang="ts">
import { ref } from 'vue'
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

</script>

<style scoped>

</style>