<script>
import SideMenu from '../components/SideMenu.vue'
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			user: null,
		}
	},
	methods: {
		changeUsernamePage: async function() {
			this.$router.push("/settings");
		},
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let res = await this.$axios.get(`/users/${localStorage.getItem("token")}`, {
					headers: {
						Authorization: localStorage.getItem("token") 
					}
				});
				this.user = res.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		}
	},
	mounted() {
		this.refresh();
	},
	components: { SideMenu }
}
</script>

<template>
	<SideMenu></SideMenu>

	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Personal profile</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="changeUsernamePage">
						Change username
					</button>
				</div>
			</div>

		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<LoadingSpinner v-if="loading"></LoadingSpinner>

		<div class="mb-3">
			<p class="card-text">
				Hello there!<br />
			</p>
		</div>

		<div class="card" v-if="!loading && user!=null">
			<div class="card-header">
				Username: {{ this.user.username }}
			</div>
			<div class="card-body">
				<p class="card-text">
					Number of photos: {{ this.user.numberOfPhotos }}<br />
					Followers: {{ this.user.followers }}<br />
					Following: {{ this.user.following }}
				</p>
			</div>
		</div>

	</div>

</template>

<style>
</style>
