<template>
  <button :class="$style.component" type="submit">{{ label }}</button>
</template>

<script>
export default {
  name: 'FormSubmitButton',
  props: {
    label: String,
  },
};
</script>

<style lang="scss" module>
.component {
  display: block;
  border: 1px solid #666;
  background-color: #666;
  color: white;
  padding: 0.5rem;
  margin-top: 1rem;
  transition: 0.3s;

  &:hover {
    background-color: transparent;
    border: 1px solid #666;
    color: #666;
  }
}
</style>
