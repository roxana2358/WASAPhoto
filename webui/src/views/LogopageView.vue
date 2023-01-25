<script>
export default {
    methods: {
		clearData: () => { localStorage.clear(); },
		next: async function() { this.$router.replace("/login"); }
    },
	mounted() { this.clearData(); }
}
</script>

<template>
	<div style="margin: 150px; text-align: center;">
		<div>
			<h1 style="color:dodgerblue; font-weight: bold; font-family:'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif">WELCOME TO</h1>
		</div>
		<div>
			<img src="/logo1.png" alt="logo">
		</div>
		<div>
			<button style="width:500px" type="button" class="btn btn-primary" @click="next">
				Log in
			</button>
		</div>
	</div>
</template>

<style>
</style>
