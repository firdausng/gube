const colors = require('tailwindcss/colors')

const primary = {
  50: '#f4f8fa',
  100: '#f4f8fa',
  200: '#f4f8fa',
  300: '#f4f8fa',
  400: '#f4f8fa',
  500: '#6a94be',
  600: '#5d82b3',
  700: '#5271a3',
  800: '#475e86',
  900: '#3d4f6b',
  950: '#283143',
}
/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './src/**/*.{html,js,svelte,ts}'
  ],
  theme: {
    extend: {
      typography: ({theme}) => ({
        DEFAULT: {
          css: {
            maxWidth: '2000px',
          },
        },
      }),
      colors: {
        transparent: 'transparent',
        current: 'currentColor',
        // black: colors.black,
        // white: colors.white,
        // gray: colors.slate,
        // green: colors.emerald,
        // purple: colors.violet,
        // yellow: colors.amber,
        // pink: colors.fuchsia,
        primary: {
          50: '#f4f8fa',
          100: '#e5edf4',
          200: '#d1e0ec',
          300: '#b1cddf',
          400: '#8cb2ce',
          500: '#6a94be',
          600: '#5d82b3',
          700: '#5271a3',
          800: '#475e86',
          900: '#3d4f6b',
          950: '#283143',
        },
        app: {
          DEFAULT: "#6a94be",
          darkest: '#283143',
          dark: '#3d4f6b',
          light: '#e5edf4',
          lightest: '#f4f8fa',
        },
      }
    },
  },
  plugins: [],
  darkMode: 'class',
}

