import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import SideMenu from './components/SideMenu.vue'
import UserPost from './components/UserPost.vue'
import UserProfile from './components/UserProfile.vue'

import './assets/dashboard.css'
import './assets/main.css'

const app = createApp(App)
app.config.globalProperties.$axios = axios;
axios.interceptors.request.use(function (config) {
        config.headers['Authorization'] = "Bearer "+localStorage.getItem('token');
        return config;
    }, function (error) {
        console.log(error);
        return Promise.reject(error);
    }
)
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("SideMenu", SideMenu)
app.component("UserPost", UserPost)
app.component("UserProfile", UserProfile)
app.use(router)
app.mount('#app')
