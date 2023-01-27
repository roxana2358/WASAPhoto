<script>
export default {
	data: function() {
		return {
			errormsg: null,
			time: null,
			stream: []
		}
	},
	methods: {
		async refresh() {
			this.errormsg = null;
			this.time = new Date().toString().split(" ")[4];
			try {
				let res = await this.$axios.get("/users/" + localStorage.getItem('token') + "/stream", null);
				this.stream = res.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		logout() {
			this.$root.logOut();
		}
	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
	<div>
		<div class="container-fluid row col-md-9 ms-sm-auto col-lg-10 px-md-2">
			<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
				<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
				<h1 class="h2">Home</h1>
				<div class="btn-toolbar mb-2 mb-md-0">
					<div class="btn-group me-2">
						<button type="button" class="btn btn-sm btn-outline-info" @click="refresh">
							Refresh (last update {{ time }})
						</button>
					</div>
					<div class="btn-group me-2">
						<button type="button" class="btn btn-sm btn-outline-info" @click="logout">
							Log out
						</button>
					</div>
				</div>
			</div>

			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

			<div>
				<div v-if="stream == null" class="card">
					<div class="card-body">
						<p>No posts to be shown.</p>
					</div>
				</div>

				<UserPost v-else v-for="post in stream" v-bind:post="post" :key="post"></UserPost>
			</div>
		</div>
	</div>
</template>

<style>
</style>
