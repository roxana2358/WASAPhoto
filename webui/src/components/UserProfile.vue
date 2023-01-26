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
			banned : this.profile.banned,
			posts : [],
			photos : []
		}
    },
    methods: {
        notifyError: async function(e) {
            this.$emit(e);
        }
    }
}
</script>

<template>
    <div v-if="profile!=null">
		<div style="padding: 10px;">
			<h5 class="card-title text-center">{{ username }}</h5>

			<div class="section">
				<p>Number of photos:</p>
				<p>{{ numberOfPhotos }}</p>
			</div>

			<div class="section">
				<p>Followers: </p>
				<p v-if="followers==null"> No followers </p>
				<p v-else>{{ followers.length }}</p>
			</div>

			<div class="section">
				<p>Following: </p>
				<p v-if="following==null"> No following </p>
				<p v-else>{{ following.length }}</p>
			</div>
		</div>
	</div>
</template>

<style>
.section {
	display: flex; 
	align-items: center; 
	justify-content: space-between;
}
</style>