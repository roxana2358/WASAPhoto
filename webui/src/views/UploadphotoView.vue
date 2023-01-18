<script>
import SideMenu from '../components/SideMenu.vue'
export default {
    data: function () {
        return {
            errormsg: null,
            loading: false
        };
    },
    methods: {},
    mounted() {
    },
    components: { SideMenu }
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
	</div>

</template>

<style>
</style>