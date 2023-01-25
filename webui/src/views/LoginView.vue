<script>
export default {
	data: function() {
		return {
			errormsg: null,
			username: null,
		}
	},
	methods: {
		doLogin: async function () {
			this.errormsg = null;
			try {
				if (this.username == null || this.username == "") {
					this.errormsg = "Your credentials are not valid!"
					return;
				}
				let res = await this.$axios.post("/session", {
					username: this.username
				});
				localStorage.setItem('username', this.username);
				localStorage.setItem('token', res.data.id);
				this.$axios.interceptors.request.use(function (config) {
					config.headers['Authorization'] = "Bearer "+localStorage.getItem('token');
					return config;
				}, function (error) {
					console.log(error);
					return Promise.reject(error);
				});
				this.$router.replace("/home");
			} catch (e) {
				this.errormsg = e.toString();
			}
		}
	}
}
</script>

<template>
	<div style="margin: 200px; text-align: center;">
		<h1 
			style="color:dodgerblue; font-weight:bold ;font-family:'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif">
			LOGIN PAGE
		</h1>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<div class="card text-center" style="width: 25rem; margin:auto;">
  			<div class="card-body">
    			<h5 class="card-title">Insert your username</h5>
				<input type="string" class="form-control" v-model="username" placeholder="Username">
				<p></p>
    			<button type="button" class="btn btn-primary" @click="doLogin"> Login</button>
  			</div>
		</div>
	</div>
</template>

<style>
</style>
