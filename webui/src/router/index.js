import {createRouter, createWebHashHistory} from 'vue-router'
import LogopageView from '../views/LogopageView.vue'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import HomepageView from '../views/HomepageView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LogopageView},
		{path: '/login', component: LoginView},
		{path: '/home', component: HomepageView},

		//{path: '/', component: HomeView},
		{path: '/link1', component: HomeView},
		{path: '/link2', component: HomeView},
		{path: '/some/:id/link', component: HomeView},
	]
})

export default router
