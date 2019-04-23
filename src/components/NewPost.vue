<template>
  <section :class="$style.component">
    <h2>Create Something New</h2>
    <slot></slot>
  </section>
</template>

<script>
export default {
  name: 'NewPost',
};
</script>

<style lang="scss" module>
.component {
  padding: 0.5em;
}
</style>
