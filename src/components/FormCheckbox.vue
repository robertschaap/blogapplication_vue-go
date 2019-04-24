<template>
  <FormGroup :label="label">
    <input
      :type="type"
      :checked="checked"
      @change="$emit('input', $event.target.checked)"
      :class="$style.component" />
  </FormGroup>
</template>

<script>
import FormGroup from '@/components/FormGroup';

export default {
  name: 'FormCheckbox',
  props: {
    label: String,
    checked: Boolean,
    type: {
      type: String,
      default: 'checkbox',
      validator: prop => prop === 'checkbox',
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
