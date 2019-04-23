<template>
  <section :class="$style.component">
    <h3>Comments</h3>
    <p>No Comments yet, be the first!</p>
    <slot></slot>
  </section>
</template>

<script>
export default {
  name: 'Comments',
};
</script>

<style lang="scss" module>
.component {
  padding: 0.5rem;
}
</style>
