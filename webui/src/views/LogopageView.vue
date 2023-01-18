<script>
export default {
	data: function() {
		return {
			loading: false,
		}
	},
    methods: {
		next: async function() {
			this.$router.replace("/login")
		}
    }
}
</script>

<template>
	<div style="position:absolute; left:730px; top: 150px">
		<div>
			<h1>Welcome to Wasa Photo</h1>
		</div>

		<div>
			<img src="../../public/logo.png" alt="logo">
		</div>

		<div>
			<button style="width:480px" v-if="!loading" type="button" class="btn btn-primary" @click="next">
				Log in
			</button>
			<LoadingSpinner v-if="loading"></LoadingSpinner>
		</div>
	</div>
</template>

<style>
</style>
