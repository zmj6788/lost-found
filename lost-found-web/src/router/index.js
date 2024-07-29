import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'login',
      component: () => import('../views/login.vue')
    },
    {
      path: '/admin',
      name: 'admin',
      component: () => import('../views/admin/admin.vue'),
      children: [
        {
          path: "",
          name: "admin_index",
          redirect: "/admin/home"
        },
        {
          path: "home",
          name: "home",
          component: () => import("../views/admin/home/home.vue")
        },
        {
          path: 'lost_view',
          name: 'lost_view',
          component: () => import('../views/admin/lost/lost_list.vue')
        },
        {
          path: 'found_view',
          name: 'found_view',
          component: () => import('../views/admin/found/found_list.vue')
        },
        {
          path: 'user_list',
          name: 'user_list',
          component: () => import('../views/admin/user/user_list.vue')
        },
        {
          path: 'image_list',
          name: 'image_list',
          component: () => import('../views/admin/image/image_list.vue')
        }
      ]
    },
    {
      path: '/web',
      name: 'web',
      component: () => import('../views/web/web.vue'),
      children: [
        {
          path: "",
          name: "web_index",
          redirect: "/web/index"
        },
        {
          path: 'article',
          name: 'article',
          component: () => import('../views/web/article/article.vue')
        },
        {
          path: 'found',
          name: 'found',
          component: () => import('../views/web/found/found.vue')
        },
        {
          path: 'index',
          name: 'index',
          component: () => import('../views/web/index/index.vue')
        }
        ,
        {
          path: 'user',
          name: 'user',
          component: () => import('../views/web/user/user.vue')
        }
        ,
        {
          path: 'show/:id',
          name: 'show',
          component: () => import('../views/web/show/show.vue')
        }
      ]
    }
  ]
})

export default router
