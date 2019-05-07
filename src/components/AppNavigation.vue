<template>
  <nav :class="$style.component">
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
import Categories from '../Categories';

export default {
  name: 'AppNavigation',
  data: function () {
    return {
      links: [
        { to: '/', label: 'Home' },
        ...this.blaat(),
      ],
    };
  },
  methods: {
    blaat: function () {
      return Object.keys(Categories).map(key => ({
        to: `/posts/${key}`,
        label: Categories[key],
      }));
    },
  },
};
</script>

<style lang="scss" module>
.component {
  padding: 0 0.5rem;
}

.link {
  padding-top: 0.5rem;
  padding-bottom: 0.5rem;
  margin-right: 0.6rem;
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
