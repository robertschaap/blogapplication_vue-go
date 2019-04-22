<template>
  <main>
    <PostList>
      <PostListTile
        v-for="post in homePosts"
        :key="post.id"
        :author="post.author"
        :isLarge="post.isLarge"
        :isMedium="post.isMedium"
        :title="post.title" />
    </PostList>
  </main>
</template>

<script>
import PostList from '@/components/PostList';
import PostListTile from '@/components/PostListTile';

export default {
  name: 'HomePage',
  components: {
    PostList,
    PostListTile,
  },
  data: () => ({
    posts: [
      { id: 0, title: 'Title', author: 'Author' },
      { id: 1, title: 'Title', author: 'Author' },
      { id: 2, title: 'Title', author: 'Author' },
      { id: 3, title: 'Title', author: 'Author' },
      { id: 4, title: 'Title', author: 'Author' },
      { id: 5, title: 'Title', author: 'Author' },
      { id: 6, title: 'Title', author: 'Author' },
      { id: 7, title: 'Title', author: 'Author' },
    ],
  }),
  computed: {
    homePosts: function () {
      return this.posts.map((post, index) => ({
        ...post,
        isLarge: index < 2,
        isMedium: index >= 2 && index < 4,
      }));
    },
  },
};
</script>
