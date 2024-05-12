import {createWebHistory, createRouter} from "vue-router";

import MainPage from "../pages/MainPage.vue"
import DetailPage from "../pages/DetailPage.vue"

export const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', name: 'main', component: MainPage },
    { path: '/events/:id', name: 'detail', component: DetailPage },
  ],
  scrollBehavior(to, from, savedPosition) {
    // always scroll to top
    return { top: 0 }
  },
})