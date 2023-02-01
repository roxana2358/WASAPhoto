<script>
export default {
    props: {
		'profile': Object
	},
    data: function() {
		return {
			id : this.profile.id,
			username : this.profile.username,
			postsArray : this.profile.photos,
			numberOfPhotos : this.profile.numberOfPhotos,
			followers : this.profile.followers,
			following : this.profile.following,
			banned : this.profile.banned
		}
    }
}
</script>

<template>
    <div id="profileBody" v-if="profile!=null">
		<h3>{{ username }}</h3>

		<div class="profileSection">
			<p>Number of photos:</p>
			<p>{{ numberOfPhotos }}</p>
		</div>

		<div class="profileSection">
			<p>Followers: </p>
			<p v-if="followers==null"> No followers </p>
			<p v-else>{{ followers.length }}</p>
		</div>

		<div class="profileSection">
			<p>Following: </p>
			<p v-if="following==null"> No following </p>
			<p v-else>{{ following.length }}</p>
		</div>
	</div>
</template>

<style>
</style>