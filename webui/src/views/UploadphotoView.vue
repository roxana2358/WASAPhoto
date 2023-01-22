<script>
export default {
    data: function () {
        return {
            errormsg: null,
            loading: false,
            image: undefined,
            previewImage: undefined,
        };
    },
    methods: {
        selectImage: async function() {
            this.loading = true;
			this.errormsg = null;
			try {
                this.image = this.$refs.file.files.item(0);
                this.previewImage = URL.createObjectURL(this.image);
            } catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
        },
        uploadImage: async function() {
            this.loading = true;
			this.errormsg = null;
			try {
				let formData = new FormData();
				formData.append("file", this.image)
				await this.$axios.post("/posts", formData, {
					headers: {
						Authorization: "Bearer "+localStorage.getItem("token"),
                        "Content-Type": "multipart/form-data",
					}
				});
				this.$router.push("/profile")
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
        }
    },
    mounted() {
    }
}
</script>

<template>
    <SideMenu></SideMenu>

	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Upload your photo</h1>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<LoadingSpinner v-if="loading"></LoadingSpinner>

        <div class="mb-3">
			<label for="description" class="form-label">Insert your photo</label>
			<input type="file" accept="image/png" ref="file" @change="selectImage">
		</div>

		<div>
			<button style="width:200px" v-if="!loading && previewImage!=undefined" type="button" class="btn btn-success" @click="uploadImage">
				Upload!
			</button>
			<LoadingSpinner v-if="loading"></LoadingSpinner>
		</div>

        <div v-if="previewImage">
            <div>
                <img style="width:500px; height:500px;" class="preview my-3" :src="previewImage" alt="" />
             </div>
        </div>
	</div>

</template>

<style>
</style>