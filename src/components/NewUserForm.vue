<template>
  <form :class="$style.component">
    <FormInput
      v-model="form.firstName"
      label="First Name" />
    <FormInput
      v-model="form.lastName"
      label="Last Name" />
    <FormInput
      v-model="form.email"
      label="E-mail" />
    <FormInput
      v-model="form.userName"
      label="User Name" />
    <FormTextarea
      v-model="form.bio"
      label="Bio" />
    <FormInput
      v-model="form.password"
      label="Password (Minimum 8 Characters)" />
    <FormInput
      v-model="form.passwordConfirm"
      label="Confirm Password" />

    <FormSubmitButton
      label="Submit" />
  </form>
</template>

<script>
import FormInput from '@/components/FormInput';
import FormSubmitButton from '@/components/FormSubmitButton';
import FormTextarea from '@/components/FormTextarea';

export default {
  name: 'NewUserForm',
  data: () => ({
    form: {
      firstName: '',
      lastName: '',
      email: '',
      userName: '',
      bio: '',
      password: '',
      passwordConfirm: '',
    },
  }),
  components: {
    FormInput,
    FormSubmitButton,
    FormTextarea,
  },
};
</script>

<style lang="scss" module>
.component {
  width: 50%;
}
</style>
