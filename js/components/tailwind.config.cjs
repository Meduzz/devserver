/** @type {import('tailwindcss').Config}*/
const config = {
  content: [
    "./src/**/*.{html,js,svelte,ts}"
  ],

  theme: {
    extend: {
      colors: {
        primary: '#7e5bef'
      }
    },
  },

  plugins: [
  ],
};

module.exports = config;
