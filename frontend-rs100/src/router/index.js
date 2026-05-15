import { createRouter, createWebHistory } from 'vue-router'
import Home from '../pages/Home.vue'
import Launching from '../pages/Launching.vue'
import History from '../pages/History.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/launching',
    name: 'Launching',
    component: Launching
  },
  {
    path: '/history',
    name: 'History',
    component: History
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/'
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
