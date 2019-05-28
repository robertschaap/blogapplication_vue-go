<template>
  <main>
    <PostCategoryHeader label="All Posts" />
    <PostList>
      <PostListTile
        v-for="post in posts"
        :key="post.id"
        :author="post.author"
        :title="post.title" />
    </PostList>
  </main>
</template>

<script>
import PostCategoryHeader from '@/components/PostCategoryHeader';
import PostList from '@/components/PostList';
import PostListTile from '@/components/PostListTile';

import ApiStatusType from '@/ApiStatusType';

export default {
  name: 'AllPostsPage',
  components: {
    PostCategoryHeader,
    PostList,
    PostListTile,
  },
  data: () => ({
    posts: [],
  }),
  created () {
    this.fetchPosts();
  },
  methods: {
    async fetchPosts () {
      const res = await fetch('/api/posts/all');
      const json = await res.json();

      if (json.status === ApiStatusType.SUCCESS) {
        this.posts = json.data;
      }
    },
  },
};
</script>
