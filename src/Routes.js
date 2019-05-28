const Routes = {
  HOME: '/',

  // Order is specific for regex matching
  POSTS_SINGLE: '/posts/:id([0-9])',
  POSTS_ALL: '/posts/:category(\\D+)',
  POSTS_NEW: '/posts/new',

  USERS_LOGIN: '/users/login',
  USERS_LOGOUT: 'users/logout',
  USERS_NEW: '/users/new',
};

export default Routes;
