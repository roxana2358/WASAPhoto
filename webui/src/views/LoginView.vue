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
		doLogin: async function () {
			this.loading = true;
			this.errormsg = null;
			try {
				if (this.username == null || this.username == "") {
					this.errormsg = "Your credentials are not valid!"
					return
				}
				let res = await this.$axios.post("/session", {
					username: this.username
				});
				localStorage.setItem('username', this.username)
				localStorage.setItem('token', res.data.id)
				this.$router.replace("/home");
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;

		}
	}
}
</script>

<template>
	<div style="position:absolute; left:720px; top: 150px; width:400px">
		<h1 style="text-align: center" class="h2">LOGIN PAGE</h1>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<LoadingSpinner v-if="loading"></LoadingSpinner>

		<div class="card text-center" style="width: 25rem;">
  			<div class="card-body">
    			<h5 class="card-title">Insert your username</h5>
				<input type="string" class="form-control" id="Username" v-model="username" placeholder="Username">
    			<button type="button" class="btn btn-primary" @click="doLogin"> Login</button>
  			</div>
		</div>
	</div>
</template>

<style>
</style>
