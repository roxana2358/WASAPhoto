<script>
export default {
	data: function() {
		return {
			errormsg: null,
			username: null,
			profile: null,
			posts: [],
			followButton: false,
			banButton: false,
		}
	},
	methods: {
		getUser: async function () {
			this.errormsg = null;
			try {
				let res1 = await this.$axios.get("/users", {
					params: {
						username: this.username
					},
				});
				let res2 = await this.$axios.get(`/users/${res1.data}`, null);
				this.profile = res2.data;
				for (let i = 0; i<this.profile.numberOfPhotos; i++ ) {
					let res3 = await this.$axios.get(`posts/${this.profile.posts[i]}`, null);
					this.posts[i] = res3.data;
				}
				this.setButtons();
			} catch (e) {
				this.errormsg = e.toString();
				this.profile = null;
			}
		},
		async refresh() {
			try {
				let res = await this.$axios.get(`/users/${this.profile.id}`, null);
				this.profile = res.data;
				this.setButtons();
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		async setButtons() {
			if (this.profile.followers == null) {
				this.followButton = true;
			} else { this.followButton = !this.profile.followers.includes(localStorage.getItem("username")); }
			if (this.profile.banned == null) {
				this.banButton = true;
			} else { this.banButton = !this.profile.banned.includes(this.profile.username); }
		},
		async followUser(id) {
			try {
				await this.$axios.put(`/users/${localStorage.getItem("token")}/following/${id}`, null, null);
				this.refresh();
			} catch (e) {
				this.errormsg = e;
			}
		},
		async unfollowUser(id) {
			try {
				await this.$axios.delete(`/users/${localStorage.getItem("token")}/following/${id}`, null);
				this.refresh();
			} catch (e) {
				this.errormsg = e;
			}
		},
		async banUser(id) {
			try {
				await this.$axios.put(`/users/${localStorage.getItem("token")}/ban/${id}`, null, null);
				this.refresh();
			} catch (e) {
				this.errormsg = e;
			}
		},
		async unbanUser(id) {
			try {
				await this.$axios.delete(`/users/${localStorage.getItem("token")}/ban/${id}`, null);
				this.refresh();
			} catch (e) {
				this.errormsg = e;
			}
		}
	}
}
</script>

<template>
	<div>
		<div class="container-fluid row col-md-9 ms-sm-auto col-lg-10 px-md-2">
			<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
				<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#eye"/></svg>
				<h1 class="h2">Search profile</h1>
				<div></div>
			</div>

			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

			<div>
				<div style="display: flex; align-items: center; justify-content: space-between;">
					<div class="card text-center searchBar">
						<div class="card-body">
							<p></p>
							<p class="card-text">Insert username: </p>
							<input style="margin: auto" type="string" class="form-control" v-model="username">
							<p></p>
							<button style="margin:auto; width: 100%;" type="button" class="btn btn-primary" @click="getUser">Search</button>
						</div>
					</div>

					<div class="card searchBar" v-if="profile" style="margin: auto;">
						<UserProfile v-bind:profile="profile" v-bind:key="profile" @refresh="refresh"></UserProfile>
						<div style="display: flex; align-items: center; justify-content: space-between;">
							<a v-show="followButton && banButton" class="actionButton btn btn-success" @click="followUser(profile.id)">Follow</a>
							<a v-show="!followButton" class="actionButton btn btn-danger" @click="unfollowUser(profile.id)">Unfollow</a>
							<a v-show="banButton" class="actionButton btn btn-success" @click="banUser(profile.id)">Ban</a>
							<a v-show="!banButton" class="actionButton btn btn-danger" @click="unbanUser(profile.id)">Unban</a>
						</div>
					</div>
				</div>

				<p></p>

				<div v-if="profile && banButton">
					<UserPost v-for="post in posts" v-bind:post="post" v-bind:key="post"></UserPost>
				</div>
			</div>
		</div>
	</div>
</template>

<style>
.actionButton {
	width: 10rem;
	height: 35px;
	border: 1px solid black;
}
.searchBar {
	height: 200px;
	width: 300px;
	margin: auto;
}
</style>
