module.exports = {
  reactStrictMode: true,
  async rewrites() {
    return [
      { source: '/apps/:appid', destination: 'http://127.0.0.1:8000/loadapp/:appid' }
    ];
  }
};
