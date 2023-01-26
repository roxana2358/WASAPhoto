<script>
export default {
    props: {
        'post': Object
    },
    data: function() {
        return {
            userId : this.post.userId,
            username : this.post.username,
            photoId : this.post.photoId,
            date : this.post.date,
            time : this.post.time,
            likes : this.post.likes,
            comments : this.post.comments,
            photo : null,
            newComment : null,
            liked : false,
            token : localStorage.getItem("token")
        }
    },
    methods: {
        getPhoto: async function() {
            try {
                // get photo
                let res = await this.$axios.get(`photos/${this.photoId}`, {
                        responseType:  'arraybuffer',
                        headers: {
                            Accept: "image/*"
                        }
                });
                let blob = new Blob([res.data]);
                this.photo= URL.createObjectURL(blob);
                // set like button
                if (this.likes == null) {
                    this.liked = false;
                } else {
                    this.liked = this.likes.includes(Number(localStorage.getItem("token")));
                }
			} catch (e) {
				//this.notifyError(e);
                console.log(e);
			}
        },
        notifyError: async function(e) {
            this.$emit(e)
        },
        commentPhoto: async function() {
            try {
                if (this.newComment==null) {
                    alert("Non e un commento");
                    return;
                }
                let res = await this.$axios.post(`/posts/${this.photoId}/comments`, {
                    comment : this.newComment
                });
                this.refresh();
            } catch(e) {
                //this.notifyError(e);
                console.log(e);
            }
        },
        deleteComment: async function(comId) {
            try {
                await this.$axios.delete(`/posts/${this.photoId}/comments/${comId}`);
                this.refresh();
            } catch(e) {
                //this.notifyError(e);
                console.log(e);
            }
        },
        likePhoto: async function() {
            try {
                let res = await this.$axios.put(`/posts/${this.photoId}/likes/${localStorage.getItem("token")}`);
                this.refresh();
            } catch(e) {
                //this.notifyError(e);
                console.log(e);
            }
        },
        unlikePhoto: async function() {
            try {
                await this.$axios.delete(`/posts/${this.photoId}/likes/${localStorage.getItem("token")}`);
                this.refresh();
            } catch(e) {
                //this.notifyError(e);
                console.log(e);
            }
        },
        refresh: async function() {
            try {
                let res = await this.$axios.get(`posts/${this.photoId}`, null);
                this.likes = res.data.likes;
                this.comments = res.data.comments;
                // set like button
                if (this.likes == null) {
                    this.liked = false;
                } else {
                    this.liked = this.likes.includes(Number(localStorage.getItem("token")));
                }
            } catch(e) {
                //this.notifyError(e);
                console.log(e);
            }
        }
    },
    mounted() {
        this.getPhoto();
    }
}
</script>

<template>

    <div v-if="post!=null" class="card mb-3">
        <div class="row g-0">
            <div class="col-md-4">
                <img v-bind:src="photo" class="img-fluid rounded-start" alt="photo">
            </div>
            <div class="col-md-8">
                <div class="card-body">
                    <div class="postSection">
                        <h5 class="card-title">{{ username }}</h5>
                        <button v-if="!liked" @click="likePhoto">Like</button>
                        <button v-else style="background-color:red;" @click="unlikePhoto">Like</button>
                        <p>#likes: {{ likes==null ? 0 : likes.length }}</p>
                    </div>
                    <p></p>
                    <div class="postSection" v-for="com in comments" :key="com">
                        <p><strong>{{ com.username }}:</strong> {{ com.comment }}</p>
                        <button v-if="com.userId==token" @click="deleteComment(com.commentId)">Delete</button>
                    </div>
                    <div>
                        <input style="width:391px;" v-model="newComment" placeholder="Write your comment here...">
                        <button @click="commentPhoto">OK!</button>
                    </div>
                    <p></p>
                    <p class="postSection"><small class="text-muted">Post from {{ time }} of {{ date }}</small></p>
                </div>
            </div>
        </div>
    </div>

</template>

<style>
.postSection {
    display: flex; 
    align-items: center; 
    justify-content: space-between;
}
</style>