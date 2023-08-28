/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.{html,js}", "node_modules/preline/dist/*.js"],
  daisyui: {
    themes: ["business"],
  },
  theme: {
    extend: {
      colors: {
        "BgBlack": "#181818"
      }
    },
  },
  plugins: [require("daisyui"), require("preline/plugin"), require("@tailwindcss/typography")],

};
