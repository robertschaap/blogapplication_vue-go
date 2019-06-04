<template>
  <main>
    <NewPost>
      <NewPostForm :handleSubmit="handleSubmit" />
    </NewPost>
  </main>
</template>

<script>
import NewPost from '@/components/NewPost';
import NewPostForm from '@/components/NewPostForm';

import ApiStatusType from '@/ApiStatusType';

export default {
  name: 'NewPostPage',
  components: {
    NewPost,
    NewPostForm,
  },
  methods: {
    async handleSubmit (form) {
      const options = {
        method: 'post',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(form),
      };

      const res = await fetch(`api/post/new`, options);
      const json = await res.json();

      if (json.status === ApiStatusType.SUCCESS) {
        // handle post
      }
    },
  },
};
</script>
