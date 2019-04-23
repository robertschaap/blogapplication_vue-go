<template>
  <main>
    <SinglePostBannerImage />
    <SinglePostContent />
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

export default {
  name: 'SinglePostPage',
  data: () => ({
    comments: [{
      body: 'Body',
      firstName: 'First',
      lastName: 'Last',
    }],
  }),
  components: {
    SinglePostBannerImage,
    SinglePostComment,
    SinglePostComments,
    SinglePostContent,
  },
};
</script>
