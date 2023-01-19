<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			stream: [],
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let res = await this.$axios.get("/users/" + localStorage.getItem('token') + "/stream", { 
					headers: { 
						Authorization: localStorage.getItem('token')
					} 
				});
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
	<SideMenu></SideMenu>

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
		<LoadingSpinner v-if="loading"></LoadingSpinner>


		<div class="card" v-if="stream.length == 0">
			<div class="card-body">
				<p>No posts to be shown.</p>
			</div>
		</div>

		<div class="card" v-if="!loading" v-for="p in stream">
			<div class="card-header">
				Username: {{ p.Username }}
			</div>
			<div class="card-body">
				<p class="card-text">
					Date: {{ p.date }}<br />
					Time: {{ p.time }}<br />
					Likes: {{ p.likes }}<br />
					Comments: {{ p.comments }}
				</p>
			</div>
		</div>
	</div>

</template>

<style>
</style>
