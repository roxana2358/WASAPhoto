<script>
export default {
	data: function() {
		return {
			errormsg: null,
			profile: null,
			posts: []
		}
	},
	methods: {
		changeUsernamePage: async function() {
			this.$router.push("/settings");
		},
		getProfile: async function() {
			this.errormsg = null;
			try {
				// get profile
				let res1 = await this.$axios.get(`/users/${localStorage.getItem("token")}`, null);
				this.profile = res1.data;
				// get posts
				for (let i = 0; i<this.profile.numberOfPhotos; i++ ) {
					let res2 = await this.$axios.get(`posts/${this.profile.posts[i]}`, null);
					this.posts[i] = res2.data;
				}
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		deletePhoto: async function(photoId) {
			try {
				console.log(photoId);
				await this.$axios.delete(`posts/${photoId}`);
				this.posts = [];
				this.getProfile();
			} catch(e) {
				this.errormsg = e.toString();
			}
		}
	},
	mounted() {
		this.getProfile();
	}
}
</script>

<template>
	<div>

		<div class="container-fluid row col-md-9 ms-sm-auto col-lg-10 px-md-2">
			<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
				<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#smile"/></svg>
				<h1 class="h2">Personal profile</h1>
				<div class="btn-toolbar mb-2 mb-md-0">
					<div class="btn-group me-2">
						<button type="button" class="btn btn-sm btn-outline-info" @click="changeUsernamePage">
							Change username
						</button>
					</div>
				</div>
			</div>

			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

			<div v-if="profile">
				<UserProfile class="card mb-3" style="margin: auto;" v-bind:profile="profile" v-bind:key="profile"></UserProfile>
				
				<div v-for="post in posts" v-bind:key="post" style="display: flex; align-items: center; justify-content: space-between;">
					<UserPost v-bind:post="post"></UserPost>
					<button @click="deletePhoto(post.photoId)">Delete</button>
				</div>
			</div>
		</div>
	</div>
</template>

<style>
</style>
