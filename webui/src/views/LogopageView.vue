<script>
export default {
	data: function() {
		return {
			loading: false,
		}
	},
    methods: {
		clearData: async function() {
			localStorage.clear();
		},
		next: async function() {
			this.$router.replace("/login")
		}
    },
	mounted() {
		this.clearData()
	}
}
</script>

<template>
	<div style="position:absolute; left:720px; top: 150px; width:500px">
		<div>
			<h1>Welcome to Wasa Photo</h1>
		</div>

		<div>
			<img src="/logo.png" alt="logo">
		</div>

		<div>
			<button style="width:500px" v-if="!loading" type="button" class="btn btn-primary" @click="next">
				Log in
			</button>
		</div>
		<LoadingSpinner v-if="loading"></LoadingSpinner>
	</div>
</template>

<style>
</style>
