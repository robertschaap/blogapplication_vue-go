<template>
  <section :class="$style.component" />
</template>

<script>
export default {
  name: 'SinglePostBannerImage',
};
</script>

<style lang="scss" module>
.component {
  position: relative;
  min-height: 300px;
  background-color: #eee;
  margin-bottom: 2rem;
  background-repeat: no-repeat;
  background-size: cover;
  background-position: center top;
}
</style>
