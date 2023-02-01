<script>
export default {
	data: function() {
		return {
			errormsg: null,
			profile: null,
			posts: [],
			trashIcons: []
		}
	},
	methods: {
		changeUsernamePage() {
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
					this.trashIcons[i] = "/trash.png";
				}
				this.posts.reverse();
			} catch (e) {
				this.errormsg = e.toString();
			}
		},
		deletePhoto: async function(photoId, index) {
			try {
				await this.$axios.delete(`posts/${photoId}`);
				this.posts = [];
				this.trashIcons[index] = "/trash.png";
				this.getProfile();
			} catch(e) {
				this.errormsg = e.toString();
			}
		},
		changeIcon(index) {
			if (this.trashIcons[index] == "/trash.png") {
				this.trashIcons[index] = "/delete.png";
			} else {
				this.trashIcons[index] = "/trash.png";
			}
		}
	},
	mounted() {
		this.$root.logIn();
		this.getProfile();
	}
}
</script>

<template>
	<div class="container-fluid row col-md-9 ms-sm-auto col-lg-10 px-md-2">
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#smile"/></svg>
			<h1 class="h2 pageTitle">Personal profile</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="myButton" @click="changeUsernamePage">
						Change username
					</button>
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<div v-if="profile">
			<UserProfile class="formCard profile" v-bind:profile="profile" v-bind:key="profile"></UserProfile>
			
			<div class="parallelSection" v-for="post in posts" v-bind:key="post">
				<UserPost v-bind:post="post"></UserPost>
				<button id="deletePost" @click="deletePhoto(post.photoId, posts.indexOf(post))" 
										@mouseenter="changeIcon(posts.indexOf(post))" 
										@mouseleave="changeIcon(posts.indexOf(post))">
					<img width="30" height="30" :src=trashIcons[posts.indexOf(post)]>
				</button>
			</div>
		</div>
	</div>
</template>

<style>
</style>