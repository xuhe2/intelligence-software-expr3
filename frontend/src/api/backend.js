import axios from 'axios';

// 创建一个 axios 实例
const backendAPI = axios.create({
    baseURL: 'http://127.0.0.1:3001',
    timeout: 1000,
});

// 添加请求拦截器
backendAPI.interceptors.request.use(
    function (config) {
        // 在发送请求之前做些什么
        const token = localStorage.getItem('jwt'); // 假设你的 JWT 存储在 localStorage 中
        if (token) {
            config.headers.Authorization = token;
        }
        return config;
    },
    function (error) {
        // 删除jwt
        localStorage.removeItem('jwt');
        // 对请求错误做些什么
        return Promise.reject(error);
    }
);

// 你可以根据需要添加更多的拦截器或方法

export default backendAPI;