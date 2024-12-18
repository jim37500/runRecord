import { createRouter, createWebHistory } from 'vue-router';
import APIService from '../services/APIService';
import UserService from '../services/UserService';
import { loginStore } from '../stores/LoginStore';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  // 換頁後移動至頂端
  scrollBehavior() {
    return { top: 0 };
  },
  routes: [
    {
      path: '/',
      meta: { requiresAuth: true },
      component: () => import('../views/root-page.vue'),
      children: [
        { path: '', name: '', component: () => import('../views/home-page.vue') },
        { path: 'activity', name: '', component: () => import('../views/my-activity.vue') },
        { path: 'activity/:activityid', name: '', component: () => import('../views/activity-detail.vue') },
        { path: 'admin/article', name: '', component: () => import('../views/activity-detail.vue') },
      ],
    },
  ],
});

router.beforeEach((to, from, next) => {
  const store = loginStore();

  if (to.matched.some((record) => record.meta.requiresAuth)) {
    // 若 有Cookie有權杖 則 取得 當前使用者資訊 後導到首頁
    if (document.cookie.indexOf('JimAndyToken=') > -1) {
      APIService.defaults.headers.common.Authorization = JSON.parse(document.cookie.split('JimAndyToken=')[1].split(';')[0]);
      store.IsLoggedIn = true;
      UserService.GetCurrentUser() // 取得 當前使用者資訊
        .then(() => next())
        .catch(() => next({ path: '/', replace: true }));
      return;
    }
  }

  // 若 無適配路徑 則 導到首頁
  if (!to.matched.length) {
    next({ path: '/', replace: true });
    return;
  }

  next();
});

export default router;
