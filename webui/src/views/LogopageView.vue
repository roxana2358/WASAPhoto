<script>
export default {
    methods: {
		next: async function() {
			this.$router.replace("/login")
		}
    }
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Welcome to Wasa Photo</h1>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<div>
			<button v-if="!loading" type="button" class="btn btn-primary" @click="next">
				Log in
			</button>
			<LoadingSpinner v-if="loading"></LoadingSpinner>
		</div>
	</div>
</template>

<style>
</style>
