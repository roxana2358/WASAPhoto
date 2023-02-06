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
			this.profile = null;
			this.posts = [];
			if (this.username===localStorage.getItem("username")) {
				this.errormsg = "Why would you ask for your own profile?";
				return;
			}
			try {
				let res1 = await this.$axios.get("/users", {
					params: {
						username: this.username
					},
				});
				let res2 = await this.$axios.get(`/users/${res1.data.id}`, null);
				this.profile = res2.data;
				for (let i = 0; i<this.profile.numberOfPhotos; i++ ) {
					let res3 = await this.$axios.get(`posts/${this.profile.posts[i]}`, null);
					this.posts[i] = res3.data;
				}
				this.posts.reverse();
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
				this.posts = [];
				for (let i = 0; i<this.profile.numberOfPhotos; i++ ) {
					let res3 = await this.$axios.get(`posts/${this.profile.posts[i]}`, null);
					this.posts[i] = res3.data;
				}
				this.posts.reverse();
				this.setButtons();
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		setButtons() {
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
	},
	mounted() {
		this.$root.logIn();
	}
}
</script>

<template>
	<div class="container-fluid row col-md-9 ms-sm-auto col-lg-10 px-md-2">
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#eye"/></svg>
			<h1 class="h2 title">Search profile</h1>
			<div></div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<div>
			<div id="searchBox">
				<div class="formCard">
					<h5 class="card-title">Insert username:</h5>
					<input type="string" class="form-control" v-model="username" placeholder="Username">
					<button type="button" class="btn" @click="getUser">Search</button>
				</div>

				<div id="profileInfo" class="formCard" v-if="profile">
					<UserProfile v-bind:profile="profile" v-bind:key="profile" @refresh="refresh"></UserProfile>
					<div id="interactionButtons">
						<button v-show="followButton && banButton" class="myButton success" @click="followUser(profile.id)">Follow</button>
						<button v-show="!followButton" class="myButton danger" @click="unfollowUser(profile.id)">Unfollow</button>
						<button v-show="banButton" class="myButton danger" @click="banUser(profile.id)">Ban</button>
						<button v-show="!banButton" class="myButton success" @click="unbanUser(profile.id)">Unban</button>
					</div>
				</div>
			</div>

			<div v-if="profile && banButton">
				<UserPost v-for="post in posts" v-bind:post="post" v-bind:key="post"></UserPost>
			</div>
		</div>
	</div>
</template>

<style>
</style>