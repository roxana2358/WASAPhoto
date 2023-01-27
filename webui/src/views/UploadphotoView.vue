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
        selectImage: async function() {
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
    }
}
</script>

<template>
	<div>

		<div class="container-fluid row col-md-9 ms-sm-auto col-lg-10 px-md-2">
			<div
				class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
				<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#plus-circle"/></svg>
				<h1 class="h2">Upload your photo</h1>
				<div></div>
			</div>

			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

			<div class="mb-3">
				<label for="description" class="form-label">Insert your photo</label>
				<input type="file" accept="image/png" ref="file" @change="selectImage">
			</div>

			<div>
				<button style="width:200px" v-if="previewImage!=undefined" type="button" class="btn btn-success" @click="uploadImage">
					Upload!
				</button>
				<LoadingSpinner></LoadingSpinner>
			</div>

			<div v-if="previewImage">
				<div>
					<img style="height:500px;" class="preview my-3" :src="previewImage" alt="" />
				</div>
			</div>
		</div>
	</div>
</template>

<style>
</style>