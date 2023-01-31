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
#loginPage {
	margin: 150px; 
	text-align: center;
}
#loginTitle {
	color: #34495E;
	font-size: 70px;
	font-weight: bolder;
	font-family: 'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif;
	margin: 40px;
}
.formCard {
	border-color: #34495E;
	border-width: 2px;
	border-style: solid;
	border-radius: 10px;
	width: 400px;
	height: 200px;
	margin: auto;
}
.formCard h5 {
	margin: 5%;
	color: #34495E;
	font-size: 20px;
	text-align: center;
}
.formCard input {
	margin: 5%;
	width: 90%;
	position: static;
}
.formCard button {
	margin: 5%;
	width: 90%;
	position: static;
	color:white;
	background-color: #34495E;
	text-align: center;
}
</style>
