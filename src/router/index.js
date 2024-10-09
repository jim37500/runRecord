import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: () => import('../views/root-page.vue'),
      children: [
        { path: '', name: '', component: () => import('../views/home-page.vue') },
        { path: 'runrecord', name: '', component: () => import('../views/run-record.vue') },
      ],
    },
  ],
});

export default router;
