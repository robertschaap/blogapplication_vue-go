<template>
  <footer :class="$style.component" role="contentinfo">
    .Dotted Pixels is a mad creation of Robert Schaap. This project is in no way affiliated with the actual domain www.dottedpixels.com which was registered but not in use at the time of this project.
  </footer>
</template>

<script>
export default {
  name: 'AppFooter',
};
</script>

<style lang="scss" module>
.component {
  border-top: 2px solid black;
  margin-top: 2rem;
  padding: 1rem 0;
  font-size: 0.8rem;
  text-align: right;
}
</style>
