<template>
  <section :class="$style.component">
    <div :class="$style.header">
      <div :class="$style.headerDescription">
        <h2>That feeling when you're coding in a new framework</h2>
        <p>Robert schaap</p>
        <p>Robert is aa true hero. He taught himself to code with the help of courses, books and a bit of magic.</p>
      </div>
      <div :class="$style.headerAvatar" />
    </div>

    <p>There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don't look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn't anything embarrassing hidden in the middle of text. All the Lorem Ipsum generators on the Internet tend to repeat predefined chunks as necessary, making this the first true generator on the Internet. It uses a dictionary of over 200 Latin words, combined with a handful of model sentence structures, to generate Lorem Ipsum which looks reasonable. The generated Lorem Ipsum is therefore always free from repetition, injected humour, or non-characteristic words etc.</p>
  </section>
</template>

<script>
export default {
  name: 'SinglePostContent',
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
  }
}

</style>
