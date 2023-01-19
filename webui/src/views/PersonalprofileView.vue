<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			user: null,
			photos: [],
			immagine: undefined,
		}
	},
	methods: {
		changeUsernamePage: async function() {
			this.$router.push("/settings");
		},
		getPhotos: async function() {
			this.loading = true;
			this.errormsg = null;
			try {
				for (let i = 0; i<this.user.photos.length; i++ ) {
					this.$axios.get(`posts/${this.user.photos[i]}`, {
							headers: {
								Authorization: localStorage.getItem("token"),
								Accept: 'image/png'
							}
					}).then( (res) => {
						// TO-DO
						this.photos[i] = res.data;
					});
				};
				this.immagine = this.photos[0];
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let res = await this.$axios.get(`/users/${localStorage.getItem("token")}`, {
					headers: {
						Authorization: localStorage.getItem("token") 
					}
				});
				this.user = res.data;
				this.getPhotos();
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		}
	},
	mounted() {
		this.refresh();
	}
}
</script>

<template>
	<SideMenu></SideMenu>

	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Personal profile</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="changeUsernamePage">
						Change username
					</button>
				</div>
			</div>

		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<LoadingSpinner v-if="loading"></LoadingSpinner>

		<div class="mb-3">
			<p class="card-text">
				Hello there!<br />
			</p>
		</div>

		<div class="card" v-if="!loading && user!=null">
			<div class="card-header">
				Username: {{ this.user.username }}
			</div>
			<div class="card-body">
				<p class="card-text">
					Number of photos: {{ this.user.numberOfPhotos }}<br />
					Followers: {{ this.user.followers }}<br />
					Following: {{ this.user.following }}<br />
				</p>
			</div>
		</div>

		<div>
            <div>
				<label for="description" class="form-label">{{ photos }}</label>
				<img :src="immagine" alt="ph">
            </div>
		</div>

	</div>

</template>

<style>
</style>
