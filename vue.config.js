module.exports = {
  css: {
    loaderOptions: {
      css: {
        modules: true,
      },
    },
  },
  devServer: {
    proxy: 'http://localhost:4000',
  },
};
