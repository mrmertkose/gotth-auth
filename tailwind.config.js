/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/*.{html,templ}",
    "./views/**/*.{html,templ}",
    "./views/**/**/*.{html,templ}"
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}

