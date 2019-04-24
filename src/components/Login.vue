<template>
  <section :class="$style.component">
    <h2>Sign In</h2>
    <slot></slot>
  </section>
</template>

<script>
export default {
  name: 'Login',
};
</script>

<style lang="scss" module>
.component {
  padding: 0.5em;
}
</style>
