const Routes = {
  HOME: '/',

  POSTS_ALL: '/posts/:category(\\D+)',
  POSTS_NEW: '/posts/new',
  POSTS_SINGLE: '/posts/:id([0-9a-fA-F]{24})',

  USERS_LOGIN: '/users/login',
  USERS_LOGOUT: 'users/logout',
  USERS_NEW: '/users/new',
};

export default Routes;
