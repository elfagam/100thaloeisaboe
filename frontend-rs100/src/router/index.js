import { createRouter, createWebHistory } from 'vue-router'
import Home from '../pages/Home.vue'
import Launching from '../pages/Launching.vue'
import History from '../pages/History.vue'
import Dashboard from '../pages/Dashboard.vue'
import Gallery from '../pages/Gallery.vue'
import Reflection from '../pages/Reflection.vue'
import Preshow from '../pages/Preshow.vue'
import Remote from '../pages/Remote.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/remote',
    name: 'Remote',
    component: Remote
  },
  {
    path: '/preshow',
    name: 'Preshow',
    component: Preshow
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
    path: '/reflection',
    name: 'Reflection',
    component: Reflection
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
