const config = {
  content: [
    "./src/**/*.{html,js,svelte,ts}",
    "./node_modules/tw-elements/dist/js/**/*.js"
  ],

  theme: {
    extend: {}
  },

  plugins: [require("tw-elements/dist/plugin")],
  darkMode: 'class',
};

module.exports = config;