/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.{html,js}"],
  daisyui: {
    themes: ["business"],
  },
  theme: {
    extend: {
      colors:{
        "BgBlack":"#181818"
      }
    },
  },
  plugins: [require("daisyui")],
  
};
