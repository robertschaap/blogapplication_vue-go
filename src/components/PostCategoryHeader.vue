<template>
  <section :class="$style.component">
    <h2 v-if="label" :class="$style.header">{{ label }}</h2>
  </section>
</template>

<script>
export default {
  name: 'PostCategoryHeader',
  props: {
    label: String,
  },
};
</script>

<style lang="scss" module>
.component {
  position: relative;
  background-color: #eee;
  margin-bottom: 2rem;
  background: linear-gradient(135deg, #ffffff 0%,#cccccc 80%,#c4c4c4 67%,#c4c4c4 80%,#7db9e8 100%);
  min-height: 150px;
}

.header {
  font-size: 3rem;
  padding: 1rem;
  background-color: white;
  font-weight: normal;
  position: absolute;
  top: 2rem;
  right: 2rem;
}
</style>
