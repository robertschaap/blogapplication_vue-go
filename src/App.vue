<template>
  <div id="app">
    <AppHeader />
    <router-view/>
    <AppFooter />
  </div>
</template>

<script>
import AppFooter from '@/components/AppFooter';
import AppHeader from '@/components/AppHeader';

export default {
  name: 'App',
  components: {
    AppFooter,
    AppHeader,
  },
};
</script>

<style lang="scss">
@import './assets/normalize.css';

html {
  box-sizing: border-box;
}

*, *:before, *:after {
  box-sizing: inherit;
}

body {
  margin: 0 auto;
  width: 80%;
  max-width: 850px;
  font-family: 'Open Sans', 'Helvetica Neue', 'Arial', sans-serif;
  font-size: 0.9rem;
}

h1, h2, h3, h4, h5 {
  font-family: 'Roboto Slab', serif;
  margin: 0;
}

h1 {
  font-size: 3rem;
  color: black;
}

p {
  margin-top: 0;
}

@media (max-width: 600px) {
  body {
    width: 90%;
  }

  h1 {
    font-size: 2rem;
  }
}
</style>
