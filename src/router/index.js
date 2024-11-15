import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  // 換頁後移動至頂端
  scrollBehavior() {
    return { top: 0 };
  },
  routes: [
    {
      path: '/',
      component: () => import('../views/root-page.vue'),
      children: [
        { path: '', name: '', component: () => import('../views/home-page.vue') },
        { path: 'activity', name: '', component: () => import('../views/my-activity.vue') },
        { path: 'activity/:activityid', name: '', component: () => import('../views/activity-detail.vue') },
      ],
    },
  ],
});

export default router;
