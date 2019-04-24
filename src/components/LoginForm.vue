<template>
  <form :class="$style.component">
    <FormInput
      v-model="form.userName"
      label="Username" />
    <FormInput
      v-model="form.password"
      label="Password"
      type="password" />

    <FormSubmitButton
      label="Submit" />
    <FormCheckbox
      v-model="form.remember"
      label="Remember me 24 hours"
      type="checkbox" />
  </form>
</template>

<script>
import FormInput from '@/components/FormInput';
import FormCheckbox from '@/components/FormCheckbox';
import FormSubmitButton from '@/components/FormSubmitButton';

export default {
  name: 'LoginForm',
  data: () => ({
    form: {
      userName: '',
      password: '',
      remember: '',
    },
  }),
  components: {
    FormInput,
    FormCheckbox,
    FormSubmitButton,
  },
};
</script>

<style lang="scss" module>
.component {
  width: 50%;
}
</style>
