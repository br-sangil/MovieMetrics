/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {},
    backgroundImage: {
      'movie-posters': "url('../public/MoviePosters.jpeg')",
    },
  },
  plugins: [],
}
