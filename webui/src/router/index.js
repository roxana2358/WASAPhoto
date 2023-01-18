import {createRouter, createWebHashHistory} from 'vue-router'
import LogopageView from '../views/LogopageView.vue'
import LoginView from '../views/LoginView.vue'
import HomepageView from '../views/HomepageView.vue'
import SearchuserView from '../views/SearchuserView.vue'
import PersonalprofileView from '../views/PersonalprofileView.vue'
import SettingsView from '../views/SettingsView.vue'
import UploadphotoView from '../views/UploadphotoView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LogopageView},
		{path: '/login', component: LoginView},
		{path: '/home', component: HomepageView},
		{path: '/search', component: SearchuserView},
		{path: '/profile', component: PersonalprofileView},
		{path: '/settings', component: SettingsView},
		{path: '/uploadPhoto', component: UploadphotoView}
	]
})

export default router
