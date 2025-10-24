const API_HOST = process.env.API_HOST || 'localhost'
const nextConfig = {
  reactStrictMode: true,
  async rewrites() {
    return [
      {
        source: '/api/:path*',
        destination: `http://${API_HOST}:8080/api/:path*`,
      },
      {
        source: '/swagger/:path*',
        destination: `http://${API_HOST}:8080/swagger/:path*`,
      },
    ]
  },
}

export default nextConfig;
