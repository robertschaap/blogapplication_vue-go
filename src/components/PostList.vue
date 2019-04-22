<template>
  <section :class="$style.component">
    <slot></slot>
  </section>
</template>

<script>
export default {
  name: 'PostList',
  components: {},
};
</script>

<style lang="scss" module>
.component {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  grid-column-gap: 1rem;
}
</style>
