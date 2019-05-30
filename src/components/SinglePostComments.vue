<template>
  <section :class="$style.component">
    <h3>Comments</h3>
    <p v-if="!hasComments">No Comments yet, be the first!</p>
    <slot></slot>
  </section>
</template>

<script>
export default {
  name: 'Comments',
  computed: {
    hasComments () {
      return this.$slots.default;
    },
  },
};
</script>

<style lang="scss" module>
.component {
  padding: 0.5rem;
}
</style>
