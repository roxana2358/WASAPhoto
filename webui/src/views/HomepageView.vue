<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			time: null,
			stream: [],
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			this.time = new Date().toString().split(" ")[4];
			try {
				let res = await this.$axios.get("/users/" + localStorage.getItem('token') + "/stream", null);
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
		<SideMenu></SideMenu>

		<div class="container-fluid row col-md-9 ms-sm-auto col-lg-10 px-md-2">
			<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
				<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
				<h1 class="h2">Home page</h1>
				<div class="btn-toolbar mb-2 mb-md-0">
					<div class="btn-group me-2">
						<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
							Refresh (last update {{ time }})
						</button>
					</div>
				</div>
			</div>

			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
			<LoadingSpinner v-if="loading"></LoadingSpinner>

			<div v-if="!loading">
				<div class="card" v-if="stream == null">
					<div class="card-body">
						<p>No posts to be shown.</p>
					</div>
				</div>

				<div class="card" v-else v-for="p in stream" :key="p">
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
		</div>
	</div>
</template>

<style>
</style>
