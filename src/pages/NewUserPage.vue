<template>
  <main>
    <NewUser>
      <NewUserForm />
    </NewUser>
  </main>
</template>

<script>
import NewUser from '@/components/NewUser';
import NewUserForm from '@/components/NewUserForm';

export default {
  name: 'NewUserPage',
  components: {
    NewUser,
    NewUserForm,
  },
};
</script>
