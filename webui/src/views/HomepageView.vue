<script>
//import token from 'login'
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			stream: null,
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let res = await this.$axios.get("/users/" + 1 + "/stream", { headers: { 'Authorization': 1 } });
				this.stream = res.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Home page</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
						Refresh (getMyStream)
					</button>
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
