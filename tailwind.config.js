/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./content/*.md", "./content/**/*.md", "./content/**/**/*.md", "./layouts/**/*.html"],
  theme: {
    extend: {
      colors: {
        'dark-orange': '#af781b',
        'oil': {
          '50': '#f5f5f1',
          '100': '#e6e5db',
          '200': '#cfccb9',
          '300': '#b3ad91',
          '400': '#9d9372',
          '500': '#8e8364',
          '600': '#7a6d54',
          '700': '#635745',
          '800': '#554a3e',
          '900': '#4a4139',
          '950': '#2c2520',
        },
        'millbrook': {
          '50': '#f8f7ee',
          '100': '#eeecd3',
          '200': '#ded9aa',
          '300': '#cbbf79',
          '400': '#bba954',
          '500': '#ac9546',
          '600': '#94783a',
          '700': '#775b31',
          '800': '#644c2f',
          '900': '#5b442e',
          '950': '#322216',
        },
        'carrot-orange': {
          '50': '#fdf7ed',
          '100': '#f9e7cc',
          '200': '#f2ce95',
          '300': '#ecaf5d',
          '400': '#e89940',
          '500': '#df7421',
          '600': '#c5551a',
          '700': '#a43a19',
          '800': '#862e1a',
          '900': '#6e2819',
          '950': '#3f1209',
        },
        'hawaiian-tan': {
          '50': '#fcf9ea',
          '100': '#f8f0c9',
          '200': '#f3e195',
          '300': '#ecc858',
          '400': '#e5b02a',
          '500': '#d59a1d',
          '600': '#b87616',
          '700': '#905415',
          '800': '#7a4419',
          '900': '#68391b',
          '950': '#3c1d0c',
        },
        'janna': {
          '50': '#fcf8ee',
          '100': '#f8efda',
          '200': '#ecd49b',
          '300': '#e2b967',
          '400': '#dba244',
          '500': '#d3842d',
          '600': '#ba6625',
          '700': '#9b4a22',
          '800': '#7f3c21',
          '900': '#69321e',
          '950': '#3b180d',
        },
        'rock': {
          '50': '#f5f3f1',
          '100': '#e6dfdb',
          '200': '#cec1ba',
          '300': '#b29c92',
          '400': '#9c7e73',
          '500': '#8d6e65',
          '600': '#785b56',
          '700': '#624746',
          '800': '#503c3c',
          '900': '#4a393a',
          '950': '#291f20',
        }
      }
    }
  },
  plugins: [],
}

