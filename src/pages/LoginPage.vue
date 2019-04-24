<template>
  <main>
    <Login>
      <LoginForm />
    </Login>
  </main>
</template>

<script>
import Login from '@/components/Login';
import LoginForm from '@/components/LoginForm';

export default {
  name: 'LoginPage',
  components: {
    Login,
    LoginForm,
  },
};
</script>
