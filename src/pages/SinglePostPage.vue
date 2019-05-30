<template>
  <main>
    <SinglePostBannerImage />
    <SinglePostContent
      :author="post.author"
      :body="post.body" />
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
    post: {
      author: {
        bio: 'Robert is a true hero. He taught himself to code with the help of courses, books and a bit of magic.',
        firstName: 'Robert',
        lastName: 'Schaap',
      },
      body: {
        body: 'There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don\'t look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn\'t anything embarrassing hidden in the middle of text. All the Lorem Ipsum generators on the Internet tend to repeat predefined chunks as necessary, making this the first true generator on the Internet. It uses a dictionary of over 200 Latin words, combined with a handful of model sentence structures, to generate Lorem Ipsum which looks reasonable. The generated Lorem Ipsum is therefore always free from repetition, injected humour, or non-characteristic words etc.',
        title: 'That feeling when you\'re coding in a new framework',
      },
    },
  }),
  components: {
    SinglePostBannerImage,
    SinglePostComment,
    SinglePostComments,
    SinglePostContent,
  },
};
</script>
