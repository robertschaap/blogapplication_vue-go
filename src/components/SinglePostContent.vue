<template>
  <section :class="$style.component">
    <div :class="$style.header">
      <div :class="$style.headerDescription">
        <h2>{{ body.title }}</h2>
        <p>{{ author.firstName }} {{ author.lastName }}</p>
        <p>{{ author.bio }}</p>
      </div>
      <div :class="$style.headerAvatar">
        <img :class="$style.headerAvatarImage" src="/avatars/1.jpg" />
      </div>
    </div>

    <p>{{ body.body }}</p>
  </section>
</template>

<script>
export default {
  name: 'SinglePostContent',
  props: {
    author: {
      type: Object,
      required: true,
    },
    body: {
      type: Object,
      required: true,
    },
  },
};
</script>

<style lang="scss" module>
.component {
  padding: 0.5rem;
}

.header {
  display: flex;
  margin-bottom: 1rem;

  &Description {
    flex: auto;
  }

  &Avatar {
    flex: none;
    width: 60px;
    height: 60px;
    border-radius: 50%;
    background-color: #eee;
    margin-left: 1rem;
    overflow: hidden;

    &Image {
      width: 100%;
      height: 100%;
    }
  }
}
</style>
