// @ts-check
import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

// Published to https://shibukawa.github.io/vkmem/ by .github/workflows/docs.yml.
export default defineConfig({
  site: 'https://shibukawa.github.io',
  base: '/vkmem',
  trailingSlash: 'always',
  integrations: [
    starlight({
      title: 'vkmem',
      description: 'Real Valkey for tests, inside your process. No Docker.',
      customCss: ['./src/styles/home.css'],
      defaultLocale: 'root',
      locales: {
        root: { label: 'English', lang: 'en' },
        ja: { label: '日本語', lang: 'ja' },
      },
      social: [{ icon: 'github', label: 'GitHub', href: 'https://github.com/shibukawa/vkmem' }],
      editLink: { baseUrl: 'https://github.com/shibukawa/vkmem/edit/main/website/' },
      sidebar: [
        { slug: 'getting-started' },
        { slug: 'performance' },
        {
          label: 'Guides',
          translations: { ja: 'ガイド' },
          items: [{ slug: 'go' }, { slug: 'python' }, { slug: 'node' }, { slug: 'java' }],
        },
        {
          label: 'Reference',
          translations: { ja: 'リファレンス' },
          items: [{ slug: 'compatibility' }, { slug: 'architecture' }],
        },
      ],
    }),
  ],
});
