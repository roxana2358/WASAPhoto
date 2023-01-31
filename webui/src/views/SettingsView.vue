<script>
export default {
	data: function() {
		return {
			errormsg: null,
			token: null,
			oldUsername: localStorage.getItem("username"),
			newUsername: null,
		}
	},
	methods: {
		changeUsername: async function () {
			this.errormsg = null;
			try {
				this.token = localStorage.getItem("token");
				await this.$axios.put(`/users/${this.token}`, { 
						username: this.newUsername
					}, null);
				localStorage.setItem("username", this.newUsername);
                this.$router.back();
			} catch (e) {
				this.errormsg = e.toString();
			}
		}
	},
	mounted() {
		this.$root.logIn();
	}
}
</script>

<template>
	<div class="container-fluid row col-md-9 ms-sm-auto col-lg-10 px-md-2">
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2 pageTitle">Change your username</h1>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<div id="usernameChange" class="formCard">
			<h5 class="card-text">This is your current username: <strong>{{ oldUsername }}</strong></h5>
			<h5 class="card-text">Insert your new username: </h5>
			<input type="string" class="form-control" v-model="newUsername" placeholder="Username">
			<button type="button" class="btn" @click="changeUsername">OK!</button>
		</div>
	</div>
</template>

<style>
#usernameChange {
	height: fit-content;
}
</style>
