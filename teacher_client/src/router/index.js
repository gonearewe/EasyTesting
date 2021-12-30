import Vue from 'vue'
import Router from 'vue-router'
/* Layout */
import Layout from '@/layout'

Vue.use(Router)

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      component: () => import('@/views/dashboard/index'),
      meta: {title: '首页', icon: 'dashboard'}
    }]
  },

  {
    path: '/teacher',
    component: Layout,
    children: [
      {
        path: 'teacher',
        component: () => import('@/views/teacher'),
        meta: {title: '教师管理', icon: 'form'}
      }
    ]
  },
  {
    path: '/student',
    component: Layout,
    children: [
      {
        path: 'student',
        component: () => import('@/views/student'),
        meta: {title: '学生管理', icon: 'form'}
      }
    ]
  },
  {
    path: '/exam',
    component: Layout,
    redirect: '/exam/index',
    children: [
      {
        path: 'index',
        component: () => import('@/views/exam'),
        meta: {title: '考试管理', icon: 'form'}
      },
      {
        path: 'detail/:id(\\d+)',
        component: () => import('@/views/exam/detail'),
        meta: {title: '考试结果', icon: 'form'},
        hidden: true
      }
    ]
  },
  {
    path: '/question',
    component: Layout,
    meta: {
      title: '试题管理',
      icon: 'form'
    },
    children: [
      {
        path: 'mcq',
        component: () => import('@/views/question/mcq'),
        meta: {
          title: '单选题',
          icon: 'form'
        },
      },
      {
        path: 'maq',
        component: () => import('@/views/question/maq'),
        meta: {
          title: '多选题',
          icon: 'form'
        },
      },
      // {
      //   path: 'bfq',
      //   component: () => import('@/views/question/bfq'),
      //   meta: {
      //     title: '填空题',
      //     icon: 'form'
      //   },
      // },
      // {
      //   path: 'tfq',
      //   component: () => import('@/views/question/tfq'),
      //   meta: {
      //     title: '判断题',
      //     icon: 'form'
      //   },
      // },
      // {
      //   path: 'crq',
      //   component: () => import('@/views/question/crq'),
      //   meta: {
      //     title: '代码阅读题',
      //     icon: 'form'
      //   },
      // },
      // {
      //   path: 'cq',
      //   component: () => import('@/views/question/cq'),
      //   meta: {
      //     title: '编程题',
      //     icon: 'form'
      //   },
      // }
    ]
  },

  // 404 page must be placed at the end !!!
  {path: '*', redirect: '/404', hidden: true}
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({y: 0}),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
