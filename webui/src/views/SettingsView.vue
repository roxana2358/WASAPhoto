<script>
import SideMenu from '../components/SideMenu.vue'
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			token: null,
			oldUsername: localStorage.getItem("username"),
			newUsername: null,
		}
	},
	methods: {
		changeUsername: async function () {
			this.loading = true;
			this.errormsg = null;
			try {
				this.token = localStorage.getItem("token");
				await this.$axios.put(`/users/${this.token}`, { 
						username: this.newUsername
					},	{ 
						headers: {
							Authorization: this.token }
					});
				localStorage.setItem("username", this.newUsername);
                this.$router.push("/profile")
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		}
	},
	components: { SideMenu }
}
</script>

<template>
	<SideMenu></SideMenu>

	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Change your username</h1>
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
				This is your current username: {{ this.oldUsername }}.<br />
				Insert your new username:
				<input type="string" class="form-control" id="Username" v-model="newUsername">
			</p>
		</div>

		<div>
			<button style="width:200px" v-if="!loading" type="button" class="btn btn-primary" @click="changeUsername">
				OK!
			</button>
			<LoadingSpinner v-if="loading"></LoadingSpinner>
		</div>

	</div>

</template>

<style>
</style>
