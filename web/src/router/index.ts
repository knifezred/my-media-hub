import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'discovery',
      component: () => import('../views/DiscoveryPage.vue'),
    },
    {
      path: '/search',
      name: 'search',
      component: () => import('../views/SearchPage.vue'),
    },
    {
      path: '/library',
      name: 'library',
      component: () => import('../views/LibraryPage.vue'),
    },
    {
      path: '/media/:id',
      name: 'media-detail',
      component: () => import('../views/MediaDetailPage.vue'),
    },
    {
      path: '/favorites',
      name: 'favorites',
      component: () => import('../views/FavoritesPage.vue'),
    },
    {
      path: '/history',
      name: 'history',
      component: () => import('../views/HistoryPage.vue'),
    },
    {
      path: '/stats',
      name: 'stats',
      component: () => import('../views/StatsPage.vue'),
    },
    {
      path: '/settings',
      name: 'settings',
      component: () => import('../views/SettingsPage.vue'),
    },
  ],
})

export default router
