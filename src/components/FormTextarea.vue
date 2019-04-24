<template>
  <FormGroup :label="label">
    <textarea
      :value="value"
      @input="$emit('input', $event.target.value)"
      :class="$style.component"/>
  </FormGroup>
</template>

<script>
import FormGroup from '@/components/FormGroup';

export default {
  name: 'FormTextarea',
  props: {
    label: String,
    value: String,
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
