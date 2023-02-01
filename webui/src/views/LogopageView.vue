<script>
export default {
    methods: {
		next() { this.$router.replace("/login"); }
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
</style>