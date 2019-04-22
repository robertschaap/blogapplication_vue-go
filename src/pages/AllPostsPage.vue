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

export default {
  name: 'AllPostsPage',
  components: {
    PostCategoryHeader,
    PostList,
    PostListTile,
  },
  data: () => ({
    posts: [
      { id: 0, title: 'Title', author: 'Author' },
      { id: 1, title: 'Title', author: 'Author' },
      { id: 2, title: 'Title', author: 'Author' },
      { id: 3, title: 'Title', author: 'Author' },
    ],
  }),
};
</script>
