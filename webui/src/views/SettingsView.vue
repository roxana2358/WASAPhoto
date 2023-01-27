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
	}
}
</script>

<template>
	<div>

		<div class="container-fluid row col-md-9 ms-sm-auto col-lg-10 px-md-2">
			<div
				class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
				<h1 class="h2">Change your username</h1>
			</div>

			<div style="align-items: center;">
				<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

				<div class="card text-center" style="width:fit-content; margin: auto;">
					<div class="card-body">
						<p class="card-text">This is yout current username: {{ oldUsername }}</p>
						<p class="card-text">Insert your new username: </p>
						<input style="width:200px; margin: auto;" type="string" class="form-control" id="Username" v-model="newUsername">
						<p></p>
						<button style="width:200px" type="button" class="btn btn-primary" @click="changeUsername">OK!</button>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<style>
</style>
