<template>
  <form :class="$style.component">
    <FormInput
      v-model="form.title"
      label="Title" />
    <FormTextarea
      v-model="form.body"
      label="Body" />
    <FormSelect
      v-model="form.category"
      :options="categoryOptions"
      label="Category" />

    <FormSubmitButton
      label="Submit" />
  </form>
</template>

<script>
import FormInput from '@/components/FormInput';
import FormSelect from '@/components/FormSelect';
import FormSubmitButton from '@/components/FormSubmitButton';
import FormTextarea from '@/components/FormTextarea';

const categoryOptions = [
  { value: 'technology', label: 'Technology' },
  { value: 'design', label: 'Design' },
  { value: 'entrepreneur', label: 'Entrepreneur' },
  { value: 'anything', label: 'Anything' },
];

export default {
  name: 'NewPostForm',
  data: () => ({
    form: {
      title: '',
      body: '',
      category: 'technology',
    },
    categoryOptions,
  }),
  components: {
    FormInput,
    FormSelect,
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
