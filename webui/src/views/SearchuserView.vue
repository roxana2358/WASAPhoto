<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			username: null,
			user: null,
			followButton: false,
			banButton: false,
		}
	},
	methods: {
		getUser: async function () {
			this.loading = true;
			this.errormsg = null;
			try {
				let res1 = await this.$axios.get("/users", {
					params: {
						username: this.username
					},
					headers: {
						Authorization: "Bearer "+localStorage.getItem('token')
					}
				});
				let res2 = await this.$axios.get(`/users/${res1.data}`, {
					headers: {
						Authorization: "Bearer "+localStorage.getItem('token')
					}
				});
				this.user = res2.data;
				if (this.user.followers == null) {
					this.followButton = true;
				} else { this.followButton = !this.user.followers.includes(localStorage.getItem("username")); }
				if (this.user.banned == null) {
					this.banButton = true;
				} else { this.banButton = !this.user.banned.includes(this.user.username); }
			} catch (e) {
				this.errormsg = e.toString();
				this.user = null;
			}
			this.loading = false;
		},
		async followUser(id) {
			this.loading = true;
			this.errormsg = null;
			try {
				await this.$axios.put(`/users/${localStorage.getItem("token")}/following/${id}`, null ,{
					headers: {
						'Authorization': localStorage.getItem("token")
					}
				});
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			this.refresh();
		},
		async unfollowUser(id) {
			this.loading = true;
			this.errormsg = null;
			try {
				await this.$axios.delete(`/users/${localStorage.getItem("token")}/following/${id}`, {
					headers: {
						'Authorization': localStorage.getItem("token")
					}
				});
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			this.refresh();
		},
		async banUser(id) {
			this.loading = true;
			this.errormsg = null;
			try {
				await this.$axios.put(`/users/${localStorage.getItem("token")}/ban/${id}`, null ,{
					headers: {
						'Authorization': localStorage.getItem("token")
					}
				});
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			this.refresh();
		},
		async unbanUser(id) {
			this.loading = true;
			this.errormsg = null;
			try {
				await this.$axios.delete(`/users/${localStorage.getItem("token")}/ban/${id}`, {
					headers: {
						'Authorization': localStorage.getItem("token")
					}
				});
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			this.refresh();
		},
		async refresh() {
			this.loading = true;
			try {
				let res1 = await this.$axios.get("/users", {
					params: {
						username: this.username
					},
					headers: {
						Authorization: "Bearer "+localStorage.getItem('token')
					}
				});
				let res2 = await this.$axios.get(`/users/${res1.data}`, {
					headers: {
						Authorization: "Bearer "+localStorage.getItem('token')
					}
				});
				this.user = res2.data;
				if (this.user.followers == null) {
					this.followButton = true;
				} else { this.followButton = !this.user.followers.includes(localStorage.getItem("username")); }
				if (this.user.banned == null) {
					this.banButton = true;
				} else { this.banButton = !this.user.banned.includes(this.user.username); }
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		}
	}
}
</script>

<template>
	<SideMenu></SideMenu>

	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Search user</h1>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<LoadingSpinner v-if="loading"></LoadingSpinner>

		<div class="mb-3">
			<label for="description" class="form-label">Insert username</label>
			<input type="string" class="form-control" id="Username" v-model="username">
		</div>

		<div>
			<button style="width:200px" v-if="!loading" type="button" class="btn btn-primary" @click="getUser">
				Search
			</button>
			<LoadingSpinner v-if="loading"></LoadingSpinner>
		</div>

		<div class="card" v-if="!loading && user!=null">
			<div class="card-header">
				Username: {{ user.username }}
			</div>
			<div class="card-body">
				<p class="card-text">
					Number of photos: {{ this.user.numberOfPhotos }}<br />
					Followers: {{ this.user.followers }}<br />
					Following: {{ this.user.following }}
				</p>
				<a v-if="followButton && banButton" href="javascript:" class="btn btn-success" @click="followUser(user.id)">Follow</a>
				<a v-if="!followButton" href="javascript:" class="btn btn-danger" @click="unfollowUser(user.id)">Unfollow</a>

				<a v-if="banButton" href="javascript:" class="btn btn-success" @click="banUser(user.id)">Ban</a>
				<a v-if="!banButton" href="javascript:" class="btn btn-danger" @click="unbanUser(user.id)">Unban</a>
			</div>
		</div>
	</div>

</template>

<style>
</style>
