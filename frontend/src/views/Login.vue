<template>
    <div class="login-container">
        <el-form class="login-form" :model="loginForm" :rules="loginRules" ref="loginFormRef">
            <h3 class="title">舌像识别系统登录</h3>
            <el-form-item prop="username">
                <el-input v-model="loginForm.username" placeholder="账号">
                    <template #prefix>
                        <el-icon class="el-input__icon">
                            <user />
                        </el-icon>
                    </template>
                </el-input>
            </el-form-item>
            <el-form-item prop="password">
                <el-input v-model="loginForm.password" placeholder="密码" show-password>
                    <template #prefix>
                        <el-icon class="el-input__icon">
                            <lock />
                        </el-icon>
                    </template>
                </el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" style="width:100%;" @click="login">登录</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { User, Lock } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

import backendAPI from '../api/backend';

const loginRules = ref({
    username: [{ required: true, trigger: 'blur', message: '请输入您的账号' }],
    password: [{ required: true, trigger: 'blur', message: '请输入您的密码' }]
})

const loginForm = reactive({
    username: '',
    password: ''
})

const login = () => {
    backendAPI.post('/login', {
        username: loginForm.username,
        password: loginForm.password
    }).then(res => {
        ElMessage({
            message: "login success",
            type: 'success',
            plain: true,
        })
    }).catch(err => {
        ElMessage({
            message: err.response.data["error"],
            type: 'error',
            plain: true,
        })
    })
}

</script>

<style scoped>
.login-container {
    width: 100%;
    height: 100vh;
    background-color: #2d3a4b;
    overflow: hidden;
}

.login-form {
    position: relative;
    width: 520px;
    max-width: 100%;
    padding: 160px 35px 0;
    margin: 0 auto;
    overflow: hidden;
}

.title {
    font-size: 26px;
    color: #eee;
    margin: 0px auto 40px auto;
    text-align: center;
    font-weight: bold;
}
</style>