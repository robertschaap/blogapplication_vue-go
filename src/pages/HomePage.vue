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
    posts: [],
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
  created () {
    this.fetchPosts();
  },
  methods: {
    async fetchPosts () {
      const res = await fetch('/api/posts');
      const json = await res.json();

      this.posts = json;
    },
  },
};
</script>
