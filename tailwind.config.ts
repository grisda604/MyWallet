/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./app/components/**/*.{js,vue,ts}",
    "./app/layouts/**/*.vue",
    "./app/pages/**/*.vue",
    "./app/plugins/**/*.{js,ts}",
    "./app/app.vue",
    "./app/error.vue",
    "./components/**/*.{js,vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./app.vue",
    "./error.vue"
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ['IBM Plex Sans Thai', 'sans-serif'],
      },
      animation: {
        'scroll-up': 'scroll-up 30s linear infinite', // ปรับความเร็วตรงนี้ได้
        'scroll-down': 'scroll-down 30s linear infinite',
      },
      keyframes: {
        'scroll-up': {
          '0%': { transform: 'translateY(0)' },
          '100%': { transform: 'translateY(-50%)' }, // เลื่อนขึ้น 50% (1 ชุด)
        },
        'scroll-down': {
          '0%': { transform: 'translateY(-50%)' }, // เริ่มจากชุดที่ 2
          '100%': { transform: 'translateY(0)' }, // เลื่อนกลับมาที่ชุดที่ 1
        },
      },
      colors: {
        brand: {
          50: '#fef3f2',
          100: '#fee5e3',
          200: '#fed0cc',
          300: '#fcaca7',
          400: '#f87171',
          500: '#ef4444',
          600: '#dc2626',
          700: '#b91c1c',
          800: '#991b1b',
          900: '#7f1d1d',
        }
      }
    },
  },
  plugins: [],
}