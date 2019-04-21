<template>
  <nav>
    <router-link
      v-for="link in links"
      :class="$style.link"
      :key="link.id"
      :to="link.to">
      {{ link.label }}
    </router-link>
  </nav>
</template>

<script>
export default {
  name: 'TheNavigation',
  data: () => ({
    links: [
      { to: '/', label: 'Home' },
      { to: '/posts', label: 'All Posts' },
      { to: '/posts/technology', label: 'Technology' },
      { to: '/posts/design', label: 'Design' },
      { to: '/posts/entrepreneur', label: 'Entrepreneur' },
      { to: '/posts/anything', label: 'Anything' },
    ],
  }),
};
</script>

<style lang="scss" module>
.link {
  padding-top: 0.5rem;
  padding-bottom: 0.5rem;
  padding-right: 0.6rem;
  text-decoration: none;
  color: #666;

  &:hover {
    text-decoration: underline;
  }
}

@media (max-width: 600px) {
  .link {
    display: block;
    padding: 0.5rem 0;
  }
}
</style>
