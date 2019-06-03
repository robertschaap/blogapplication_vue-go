<template>
  <main>
    <Login>
      <LoginForm :handleSubmit="handleSubmit" />
    </Login>
  </main>
</template>

<script>
import Login from '@/components/Login';
import LoginForm from '@/components/LoginForm';

import ApiStatusType from '@/ApiStatusType';

export default {
  name: 'LoginPage',
  components: {
    Login,
    LoginForm,
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

      const res = await fetch(`/api/users/login`, options);
      const json = await res.json();

      if (json.status === ApiStatusType.SUCCESS) {
        // handle login
      }
    },
  },
};
</script>
