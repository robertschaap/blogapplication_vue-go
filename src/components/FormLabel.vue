<template>
  <label :class="$style.component">{{ label }}</label>
</template>

<script>
export default {
  name: 'FormLabel',
  props: {
    label: String,
  },
};
</script>

<style lang="scss" module>
.component {
  display: block;
  width: 100%;
}
</style>
