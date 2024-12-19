<template>
    <div class="check">
        <el-upload class="avatar-uploader" action="http://127.0.0.1:3000/handle_picture/" :show-file-list="false"
            :on-success="handleAvatarSuccess" :before-upload="beforeAvatarUpload">
            <img v-if="imageUrl" :src="imageUrl" class="avatar" />
            <el-icon v-else class="avatar-uploader-icon">
                <Plus />
            </el-icon>
        </el-upload>
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import { ElNotification } from 'element-plus'

import type { UploadProps } from 'element-plus'

const imageUrl = ref('')

const handleAvatarSuccess: UploadProps['onSuccess'] = (
    response,
    uploadFile
) => {
    imageUrl.value = URL.createObjectURL(uploadFile.raw!)
    let kind = response.kind
    if (kind == 0) {
        ElNotification({
            title: 'health',
            message: '无转移症状',
            type: 'success',
        })
    } else if (kind == 1) {
        ElNotification({
            title: 'warning',
            message: '跨转移症状',
            type: 'warning',
        })
    } else {
        ElNotification({
            title: 'danger',
            message: '泛转移症状',
            type: 'error',
        })
    }
}

const beforeAvatarUpload: UploadProps['beforeUpload'] = (rawFile) => {
    if (rawFile.type !== 'image/jpeg' && rawFile.type !== 'image/png') {
        ElMessage.error('Avatar picture must be JPG format!')
        return false
    } else if (rawFile.size / 1024 / 1024 > 2) {
        ElMessage.error('Avatar picture size can not exceed 2MB!')
        return false
    }
    return true
}
</script>

<style scoped>
.avatar-uploader .avatar {
    width: 178px;
    height: 178px;
    display: block;
}
</style>

<style>
.avatar-uploader .el-upload {
    border: 1px dashed var(--el-border-color);
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
    border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    text-align: center;
}

/* 居中 */
.avatar-uploader {
    margin-left: 12vh;
    margin-top: 30vh;
}

/* 背景 */
.check {
    /* background-color: #8c939d; */
}
</style>