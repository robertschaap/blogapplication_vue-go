<template>
  <main>
    <NewPost>
      <NewPostForm />
    </NewPost>
  </main>
</template>

<script>
import NewPost from '@/components/NewPost';
import NewPostForm from '@/components/NewPostForm';

export default {
  name: 'NewPostPage',
  components: {
    NewPost,
    NewPostForm,
  },
};
</script>
