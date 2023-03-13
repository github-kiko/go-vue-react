<template>
    <div>

        <el-dialog v-model="centerDialogVisible" title="Warning" width="30%" center>
    <span>
      It should be noted that the content will not be aligned in center by
      default
    </span>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="centerDialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="centerDialogVisible = false">
          Confirm
        </el-button>
      </span>
    </template>
  </el-dialog>

    </div>
</template>

<script setup lang="ts">

import { ref } from 'vue'

const centerDialogVisible = ref(false)

</script>

<style scoped>

.dialog-footer button:first-child {
  margin-right: 10px;
}

</style>