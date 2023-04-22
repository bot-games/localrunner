import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router'
import docsView from '@/views/DocsView.vue'
import battlesView from "@/views/BattlesView.vue";
import battleView from "@/views/BattleView.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    component: docsView
  },
  {
    path: '/battles',
    component: battlesView
  },
  {
    path: '/battle/:battleId',
    component: battleView
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
