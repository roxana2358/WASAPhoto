<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			username: null,
		}
	},
	methods: {
		clearData: async function() {
			localStorage.clear
		},
		doLogin: async function () {
			this.loading = true;
			this.errormsg = null;
			try {
				let res = await this.$axios.post("/session", {
					username: this.username
				});
				localStorage.clear
				localStorage.setItem('token', res.data.ID)
				this.$router.replace("/home");
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		}
	},
	mounted() {
		this.clearData()
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Login page</h1>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<div class="mb-3">
			<label for="description" class="form-label">Insert your username</label>
			<input type="string" class="form-control" id="Username" v-model="username">
		</div>

		<div>
			<button v-if="!loading" type="button" class="btn btn-primary" @click="doLogin">
				Login
			</button>
			<LoadingSpinner v-if="loading"></LoadingSpinner>
		</div>
	</div>
</template>

<style>
</style>
