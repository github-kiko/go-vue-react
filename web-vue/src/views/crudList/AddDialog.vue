<template>
    <div>

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
        <el-button type="primary" @click="centerDialogVisible = false">
          确认
        </el-button>
      </span>
    </template>
  </el-dialog>

    </div>
</template>

<script setup lang="ts">

import { reactive, ref } from 'vue'

const labelPosition = ref('right')

const form = reactive({
  name: '',
  age: '',
  school: '',
  phone: '',
  address: '',
})

const centerDialogVisible = ref(true)
const title = ref('新增')



</script>

<style scoped>

.dialog-footer button:first-child {
  margin-right: 10px;
}

</style>