<script>
export default {
	props: ["present", "username", "photo", "likes", "comments"]
}
</script>

<template>
    <div v-if="present" class="card" style="width: 18rem;">
        <h5 class="card-title">{{ username }}</h5>
        <img :src=photo class="card-img-top" alt="photo">
        <div class="card-body">
            <a href="#" class="btn btn-primary">Like</a>
            <a href="#" class="btn btn-primary">Comment</a>
        </div>
    </div>
</template>