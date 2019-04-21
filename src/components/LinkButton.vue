<template>
  <router-link :class="$style.component" to="to">{{ label }}</router-link>
</template>

<script>
export default {
  name: 'LinkButton',
  props: {
    'label': String,
    'to': String,
  },
};
</script>

<style lang="scss" module>
.component {
    border: 1px solid #666;
    border-radius: 5px;
    transition: 0.3s;
    padding: 0.3rem 0.5rem;
    text-decoration: none;
    color: #666;

    &:not(:first-child) {
      margin-left: 5px;
    }

    &:hover {
      text-decoration: none;
      color: white;
      background-color: #666;
    }
}
</style>
