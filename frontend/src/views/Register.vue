<template>
    <div class="register-container">
        <el-form class="register-form" :model="registerForm" :rules="registerRules" ref="registerFormRef">
            <h3 class="title">用户注册</h3>
            <el-form-item prop="username">
                <el-input v-model="registerForm.username" placeholder="请输入账号">
                    <template #prefix>
                        <el-icon class="el-input__icon">
                            <user />
                        </el-icon>
                    </template>
                </el-input>
            </el-form-item>
            <el-form-item prop="password">
                <el-input v-model="registerForm.password" placeholder="请输入密码" show-password>
                    <template #prefix>
                        <el-icon class="el-input__icon">
                            <lock />
                        </el-icon>
                    </template>
                </el-input>
            </el-form-item>
            <el-form-item prop="confirmPassword">
                <el-input v-model="registerForm.confirmPassword" placeholder="请确认密码" show-password>
                    <template #prefix>
                        <el-icon class="el-input__icon">
                            <lock />
                        </el-icon>
                    </template>
                </el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" style="width:100%;" @click="register">注册</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { User, Lock } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

import router from '../router';

import backendAPI from '../api/backend';
import { encryptPassword } from '../utils/crypt';

const validateConfirmPassword = (rule, value, callback) => {
    if (value !== registerForm.password) {
        callback(new Error('两次输入密码不一致!'))
    } else {
        callback()
    }
}

const registerRules = ref({
    username: [
        { required: true, trigger: 'blur', message: '请输入您的账号' },
        { min: 3, max: 15, message: '账号长度在 3 到 15 个字符', trigger: 'blur' }
    ],
    password: [
        { required: true, trigger: 'blur', message: '请输入您的密码' },
        { min: 6, message: '密码长度不能少于6位', trigger: 'blur' }
    ],
    confirmPassword: [
        { required: true, trigger: 'blur', message: '请再次输入您的密码' },
        { validator: validateConfirmPassword, trigger: 'blur' }
    ]
})

const registerForm = reactive({
    username: '',
    password: '',
    confirmPassword: ''
})

const register = () => {
    // encrypt password
    let body = {
        username: registerForm.username,
        password: encryptPassword(registerForm.password)
    }
    backendAPI.post('/register', body).then(res => {
        ElMessage({
            message: "注册成功，请登录",
            type: 'success',
            plain: true,
        })
        // 可以在这里进行页面跳转，例如跳转到登录页面
        router.push('/')
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
.register-container {
    width: 100%;
    height: 100vh;
    background-color: #2d3a4b;
    overflow: hidden;
}

.register-form {
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