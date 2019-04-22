<template>
  <article :class="$style.component">
    <router-link to="/">
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
    title: String,
  },
};
</script>

<style lang="scss" module>
.component {
  grid-column: auto / span 2;
  border-bottom: 1px solid #ccc;
  margin-bottom: 1rem;

  a {
    text-decoration: none;
    color: black;
  }

  &:hover {
    box-shadow: 0px 5px 35px -10px rgba(204,204,204,1);
  }
}

.text {
  padding: 0.5rem;
}
</style>
