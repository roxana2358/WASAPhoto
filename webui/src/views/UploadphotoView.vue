<script>
export default {
    data: function () {
        return {
            errormsg: null,
            image: undefined,
            previewImage: undefined,
        };
    },
    methods: {
        selectImage() {
			this.errormsg = null;
			try {
                this.image = this.$refs.file.files.item(0);
                this.previewImage = URL.createObjectURL(this.image);
            } catch (e) {
				this.errormsg = e.toString();
			}
        },
        uploadImage: async function() {
			this.errormsg = null;
			if (this.image == null || this.image == undefined) {
				this.errormsg = "You did not choose a photo!";
				return;
			}
			try {
				let formData = new FormData();
				formData.append("file", this.image)
				await this.$axios.post("/posts", formData, {
					headers: {
                        "Content-Type": "multipart/form-data",
					}
				});
				this.$router.push("/profile")
			} catch (e) {
				this.errormsg = e.toString();
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
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#plus-circle"/></svg>
			<h1 class="h2 pageTitle">Upload your photo</h1>
			<div></div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<div id="uploadPhoto">
			<div class="formCard">
				<h5 class="card-title">Choose your photo:</h5>
				<input type="file" accept="image/*" ref="file" @change="selectImage">
				<button type="button" class="btn" @click="uploadImage">Upload!</button>
			</div>

			<div id="previewImage" v-if="previewImage">
				<img height="400" class="preview" :src="previewImage" alt="" />
			</div>
		</div>
	</div>
</template>

<style>
</style>