import Vue from 'vue';
import Router from 'vue-router';
import Routes from './routes';

import AllPostsPage from './pages/AllPostsPage.vue';
import HomePage from './pages/HomePage.vue';
import LoginPage from './pages/LoginPage.vue';
import NewPostPage from './pages/NewPostPage.vue';
import NewUserPage from './pages/NewUserPage.vue';
import SinglePostPage from './pages/SinglePostPage.vue';

Vue.use(Router);

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    { path: Routes.HOME, component: HomePage },

    { path: Routes.POSTS_ALL, component: AllPostsPage },
    { path: Routes.POSTS_NEW, component: NewPostPage },
    { path: Routes.POSTS_SINGLE, component: SinglePostPage },

    { path: Routes.USERS_LOGIN, component: LoginPage },
    { path: Routes.USERS_NEW, component: NewUserPage },
  ],
});
