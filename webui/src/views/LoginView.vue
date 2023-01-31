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
				if (this.username == null || this.username.localeCompare("")==0) {
					this.errormsg = "Your credentials are not valid!"
					return;
				}
				let res = await this.$axios.post("/session", {
					username: this.username
				});
				localStorage.setItem('username', this.username);
				localStorage.setItem('token', res.data.id);
				this.$root.logIn();
				this.$router.push("/home");
			} catch (e) {
				this.errormsg = e.toString();
			}
		}
	}
}
</script>

<template>
	<div id="loginPage">
		<h1 id="loginTitle">
			LOGIN PAGE
		</h1>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<div class="formCard">
			<h5 class="card-title">Insert your username:</h5>
			<input type="string" class="form-control" v-model="username" placeholder="Username">
			<button type="button" class="btn" @click="doLogin">Login</button>
		</div>
	</div>
</template>

<style>
</style>