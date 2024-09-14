/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/*.templ","./view/components/*.templ"],
  theme: {
    
    
    extend: {
      colors: {
        primary: {
          50: '#f9fafb',
          100: '#f3f4f6',
          200: '#e5e7eb',
          300: '#d1d5db',
          400: '#9ca3af',
          500: '#6b7280',
          600: '#4b5563',
          700: '#374151',
          800: '#1f2937',
          900: '#111827',
        },
      },
    },
  },
  plugins: [],
  // 添加以下配置
  future: {
    // 移除 Tailwind 添加的变体前缀
    removeDeprecatedGapUtilities: true,
    purgeLayersByDefault: true,
  },
  // 配置为使用 'class' 而不是 'className'
  experimental: {
    classRegex: [
      'class\\s*=\\s*"([^"]*)"',
      'class\\s+([^"]*)(?=\\s)(?!\\s*[=])',
      'class\\s*=\\s*\'([^\']*)\'',
    ],
  },
}

