<template>
  <main>
    <NewUser>
      <NewUserForm :handleSubmit="handleSubmit" />
    </NewUser>
  </main>
</template>

<script>
import NewUser from '@/components/NewUser';
import NewUserForm from '@/components/NewUserForm';

import ApiStatusType from '@/ApiStatusType';

export default {
  name: 'NewUserPage',
  components: {
    NewUser,
    NewUserForm,
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

      const res = await fetch(`/api/users/new`, options);
      const json = await res.json();

      if (json.status === ApiStatusType.SUCCESS) {
        // handle post
      }
    },
  },
};
</script>
