/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  env: {
    API_HOST: "http://localhost:8080/"
  }
}

module.exports = nextConfig
