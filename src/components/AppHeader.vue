<template>
  <header :class="$style.component">
    <div :class="$style.top">
      <router-link :class="$style.title" :to="homeRoute">
        <h1>.Dotted Pixels</h1>
      </router-link>
      <nav v-if="isLoggedIn">
        <LinkButton label="+" :to="newPostRoute" />
        <LinkButton label="Log Out" :to="logoutRoute" />
      </nav>
      <nav v-else="">
        <LinkButton label="Sign In" :to="loginRoute" />
        <LinkButton label="Sign Up" :to="newUserRoute" />
      </nav>
    </div>
    <AppNavigation />
  </header>
</template>

<script>
import Routes from '../Routes';
import LinkButton from '@/components/LinkButton';
import AppNavigation from '@/components/AppNavigation';

export default {
  name: 'AppHeader',
  data: () => ({
    homeRoute: Routes.HOME,
    loginRoute: Routes.USERS_LOGIN,
    logoutRoute: Routes.USERS_LOGOUT,
    newPostRoute: Routes.POSTS_NEW,
    newUserRoute: Routes.USERS_NEW,
  }),
  components: {
    LinkButton,
    AppNavigation,
  },
  props: {
    isLoggedIn: Boolean,
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

.title {
  text-decoration: none;
}

@media (max-width: 600px) {
  .top {
    display: block;
  }

  .title {
    display: block;
    margin-bottom: 1rem;
  }
}
</style>
