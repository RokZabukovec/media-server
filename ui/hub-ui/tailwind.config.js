/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './src/components/*.{html,vue,js}',
    './src/views/*.{html,vue,js}',
  ],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],}