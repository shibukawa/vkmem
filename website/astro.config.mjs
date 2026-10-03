// @ts-check
import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

// Published to https://shibukawa.github.io/vkmem/ by .github/workflows/docs.yml.
export default defineConfig({
  site: 'https://shibukawa.github.io',
  base: '/vkmem',
  trailingSlash: 'always',
  // the Go guide was one page before it was split by test boundary
  redirects: {
    '/go/': '/vkmem/guides/go/basics/',
    '/ja/go/': '/vkmem/ja/guides/go/basics/',
  },
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
          items: [
            {
              label: 'Go',
              items: [
                { label: 'Basics', translations: { ja: '基本' }, slug: 'guides/go/basics' },
                { label: 'Unit tests', translations: { ja: 'ユニットテスト' }, slug: 'guides/go/testing' },
                { label: 'API tests', translations: { ja: 'APIテスト' }, slug: 'guides/go/api-testing' },
                { label: 'E2E tests', translations: { ja: 'E2Eテスト' }, slug: 'guides/go/e2e-testing' },
              ],
            },
            { slug: 'python' },
            { slug: 'node' },
            { slug: 'java' },
          ],
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
