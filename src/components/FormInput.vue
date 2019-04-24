<template>
  <FormGroup :label="label">
    <input
      :type="type"
      :value="value"
      @input="$emit('input', $event.target.value)"
      :class="$style.component" />
  </FormGroup>
</template>

<script>
import FormGroup from '@/components/FormGroup';

export default {
  name: 'FormInput',
  props: {
    label: String,
    value: String,
    type: {
      type: String,
      default: 'text',
      validator: prop => ['text', 'password'].some(v => v === prop),
    },
  },
  components: {
    FormGroup,
  },
};
</script>

<style lang="scss" module>
.component {
  display: block;
  width: 100%;
  border: 1px solid #eee;
}
</style>
