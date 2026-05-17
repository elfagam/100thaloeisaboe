import { createRouter, createWebHistory } from 'vue-router'
import Home from '../pages/Home.vue'
import Launching from '../pages/Launching.vue'
import History from '../pages/History.vue'
import Dashboard from '../pages/Dashboard.vue'
import Gallery from '../pages/Gallery.vue'

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
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard
  },
  {
    path: '/gallery',
    name: 'Gallery',
    component: Gallery
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
