<template>
  <section :class="$style.component">
    <h2>Join Us!</h2>
    <slot></slot>
  </section>
</template>

<script>
export default {
  name: 'NewUser',
};
</script>

<style lang="scss" module>
.component {
  padding: 0.5em;
}
</style>
