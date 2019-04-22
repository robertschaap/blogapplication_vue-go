<template>
  <header :class="$style.component">
    <div :class="$style.top">
      <h1 :class="$style.title">.Dotted Pixels</h1>
      <nav>
        <LinkButton label="Sign In" to="/" />
        <LinkButton label="Sign Up" to="/" />
      </nav>
    </div>
    <AppNavigation />
  </header>
</template>

<script>
import LinkButton from '@/components/LinkButton';
import AppNavigation from '@/components/AppNavigation';

export default {
  name: 'AppHeader',
  components: {
    LinkButton,
    AppNavigation,
  },
};
</script>

<style lang="scss" module>
.component {
  margin-bottom: 1rem;
  padding: 1rem 0;
}

.top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1rem 0;
}

@media (max-width: 600px) {
  .top {
    display: block;
  }

  .title {
    margin-bottom: 1rem;
  }
}
</style>
