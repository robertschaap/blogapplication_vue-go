<template>
  <div :class="$style.component">
    <FormLabel :label="label" />
    <slot></slot>
  </div>
</template>

<script>
import FormLabel from '@/components/FormLabel';

export default {
  name: 'FormGroup',
  props: {
    label: String,
  },
  components: {
    FormLabel,
  },
};
</script>

<style lang="scss" module>
.component {
  margin-top: 1rem;
}
</style>
