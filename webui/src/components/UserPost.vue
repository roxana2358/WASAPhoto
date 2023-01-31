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
				alert(e);
			}
        },
        commentPhoto: async function() {
            try {
                if (this.newComment==null) {
                    alert("Your comment is not a real comment!");
                    return;
                }
                let res = await this.$axios.post(`/posts/${this.photoId}/comments`, {
                    comment : this.newComment
                });
                this.newComment = null;
                this.refresh();
            } catch(e) {
                alert(e);
            }
        },
        deleteComment: async function(comId) {
            try {
                await this.$axios.delete(`/posts/${this.photoId}/comments/${comId}`);
                this.refresh();
            } catch(e) {
                alert(e);
            }
        },
        likePhoto: async function() {
            try {
                let res = await this.$axios.put(`/posts/${this.photoId}/likes/${localStorage.getItem("token")}`);
                this.refresh();
            } catch(e) {
                alert(e);
            }
        },
        unlikePhoto: async function() {
            try {
                await this.$axios.delete(`/posts/${this.photoId}/likes/${localStorage.getItem("token")}`);
                this.refresh();
            } catch(e) {
                alert(e);
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
                alert(e);
            }
        }
    },
    mounted() {
        this.getPhoto();
    }
}
</script>

<template>
    <div class="formCard post" id="postContainer" v-if="post!=null">
        <div id="postImage">
            <img v-bind:src="photo" class="img-fluid rounded-start" alt="photo">
        </div>
        <div id="postInteractive">
            <div class="parallelSection">
                <h5 class="card-title">{{ username }}</h5>
                <div class="parallelSection">
                    <img width="17" height="17" src="/comment.png">
                    <h5 class="postInfo">{{ comments==null ? 0 : comments.length }}</h5>
                    <a v-if="!liked" @click="likePhoto">
                        <img width="20" height="20" src="/like-empty.png">
                    </a>
                    <a v-else @click="unlikePhoto">
                        <img width="20" height="20" src="/like-full.png">
                    </a>
                    <h5 class="postInfo">{{ likes==null ? 0 : likes.length }}</h5>
                </div>
            </div>

            <hr>

            <div class="comments" v-if="comments!=null">
                <div class="parallelSection comment" v-for="com in comments" :key="com">
                    <span><strong>{{ com.username }}:</strong> {{ com.comment }}</span>
                    <button type="button" class="btn" v-if="com.userId==token" @click="deleteComment(com.commentId)">Delete</button>
                </div>
            </div>

            <div class="parallelSection" v-if="comments==null">
                <p>Be the first to comment!</p>
            </div>

            <hr>

            <div class="parallelSection newComment">
                <input type="string" class="form-control" v-model="newComment" placeholder="Write your comment here...">
                <button type="button" class="btn" @click="commentPhoto">Send</button>
            </div>

            <p class="parallelSection"><small class="text-muted">Posted on {{ time }} of {{ date }}</small></p>

        </div>
    </div>
</template>

<style>
</style>