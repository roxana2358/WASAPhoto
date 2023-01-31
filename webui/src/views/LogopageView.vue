<script>
export default {
    methods: {
		next: async function() { this.$router.replace("/login"); }
    },
	mounted() { 
		localStorage.clear(); 
	}
}
</script>

<template>
	<div id="logoPage">
		<div id="greeting">
			<h1>WELCOME TO</h1>
		</div>
		<div id="logo">
			<img src="/logo.png" alt="logo">
		</div>
		<div id="login">
			<button type="button" class="btn" @click="next">
				Log in
			</button>
		</div>
	</div>
</template>

<style>
#logoPage {
	margin: 100px; 
	text-align: center;
}
#greeting h1 {
	color: #34495E;
	font-size: 70px;
	font-weight: bolder;
	font-family: 'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif;
	margin: 20px;
}
#logo {
	margin: 50px;
}
#login button {
	width: 500px;
	margin: 20px;
	background-color:#34495E;
	color:white;
}
</style>
