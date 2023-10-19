import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Content from '../views/Content.vue'
import Publish from '../views/Publish.vue'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import RegisterE from '../views/RegisterE.vue'
const originalPush = VueRouter.prototype.push;
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err);
}
Vue.use(VueRouter)

  const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/post/:id',
    name: 'Content',
    component: Content
  },
  {
    path: '/publish',
    name: 'Publish',
    component: Publish,
    meta: { requireAuth: true }
  },
  {
    path: '/login',
    name:"Login",
    component: Login
  },
  {
    path: '/register',
    name:"Register",
    component: Register
  },
  {
    path: '/registere',
    name:"RegisterE",
    component: RegisterE
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
