<template>
  <div :class="$style.component">
    <div :class="$style.image" />
    <div :class="$style.body">
      <p :class="$style.author">{{ firstName }} {{ lastName }}</p>
      <p>{{ body}}</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'SinglePostComment',
  props: {
    body: String,
    firstName: String,
    lastName: String,
  },
};
</script>

<style lang="scss" module>
.component {
  display: flex;
}

.image {
  flex: none;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background-color: #eee;
  margin-right: 1rem;
}

.body {
  flex: auto;
}

.author {
  margin-bottom: 0;
  border-bottom: 1px solid #eee;
}
</style>
