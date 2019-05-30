<template>
  <main v-if="body">
    <SinglePostBannerImage />
    <SinglePostContent
      :author="author"
      :body="body" />
    <SinglePostComments>
      <SinglePostComment
        v-for="(comment, index) in comments"
        :key="index"
        :body="comment.body"
        :firstName="comment.firstName"
        :lastName="comment.lastName" />
    </SinglePostComments>
  </main>
</template>

<script>
import SinglePostBannerImage from '@/components/SinglePostBannerImage';
import SinglePostComment from '@/components/SinglePostComment';
import SinglePostComments from '@/components/SinglePostComments';
import SinglePostContent from '@/components/SinglePostContent';

import ApiStatusType from '@/ApiStatusType';

export default {
  name: 'SinglePostPage',
  data: () => ({
    author: null,
    body: null,
    comments: null,
  }),
  components: {
    SinglePostBannerImage,
    SinglePostComment,
    SinglePostComments,
    SinglePostContent,
  },
  created () {
    this.fetchPost();
  },
  methods: {
    async fetchPost () {
      const { id } = this.$route.params;

      const res = await fetch(`/api/post/${id}`);
      const json = await res.json();

      if (json.status === ApiStatusType.SUCCESS) {
        this.body = json.data.body;
        this.author = json.data.author;
        this.comments = json.data.comments;
      }
    },
  },
};
</script>
