<template>
  <section
    :class="$style.component"
    :style="{ backgroundImage: `url(/banners/${randomImageId()}.jpg`}" />
</template>

<script>
export default {
  name: 'SinglePostBannerImage',
  methods: {
    randomImageId () {
      return Math.floor(Math.random()* 6) + 1;
    },
  },
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
