<template>
  <FormGroup :label="label">
    <textarea
      v-bind:value="value"
      v-on:input="$emit('input', $event.target.value)"
    :class="$style.component"/>
  </FormGroup>
</template>

<script>
import FormGroup from '@/components/FormGroup';

export default {
  name: 'FormTextarea',
  props: {
    label: String,
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
