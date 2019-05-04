<template>
  <article
    :class="{
      [$style.component]: true,
      [$style.noSpan]: isLarge || isMedium,
    }">
    <router-link to="/">
      <div
        v-if="isLarge || isMedium"
        :class="{
          [$style.imageLarge]: isLarge,
          [$style.imageMedium]: isMedium,
        }"
        :style="{ backgroundImage: `url(/banners/${randomImageId()}.jpg`}" />
      <div :class="$style.text">
        <h2>{{ title }}</h2>
        <p>{{ author }}</p>
      </div>
    </router-link>
  </article>
</template>

<script>
export default {
  name: 'PostListTile',
  props: {
    author: String,
    isLarge: Boolean,
    isMedium: Boolean,
    title: String,
  },
  methods: {
    randomImageId () {
      return Math.floor(Math.random()* 6) + 1;
    },
  },
};
</script>

<style lang="scss" module>
.component {
  grid-column: auto / span 2;
  border-bottom: 1px solid #ccc;
  margin-bottom: 1rem;

  &.noSpan {
    grid-column: auto / span 1;
  }

  a {
    text-decoration: none;
    color: black;
  }

  &:hover {
    box-shadow: 0px 5px 35px -10px rgba(204,204,204,1);
  }
}

.image {
  background-repeat: no-repeat;
  background-size: cover;
  background-position: center top;
  background-color: #eee;

  &Large {
    @extend .image;
    min-height: 200px;
  }

  &Medium {
    @extend .image;
    min-height: 100px;
  }
}

.text {
  padding: 0.5rem;
}
</style>
