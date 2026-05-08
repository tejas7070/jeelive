import { createRouter, createWebHistory } from "vue-router"
import Home from "../pages/Home.vue"
import Login from "../pages/Login.vue"
import StudentList from "../pages/StudentList.vue"
import AddStudent from "../pages/AddStudent.vue"
import ViewStudent from "../pages/viewStudent.vue"
import { hasToken } from "../services/api"

const routes = [
  { path: "/login", name: "login", component: Login },
  { path: "/", name: "home", component: Home, meta: { requiresAuth: true } },
  { path: "/students", name: "students", component: StudentList, meta: { requiresAuth: true } },
  { path: "/add", name: "add-student", component: AddStudent, meta: { requiresAuth: true } },
  { path: "/students/:id/edit", name: "edit-student", component: AddStudent, meta: { requiresAuth: true } },
  {
    path: "/students/:id",
    name: "view-student",
    component: ViewStudent,
    props: true,
    meta: { requiresAuth: true }
  }
]

export const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to) => {
  if (to.name === "login" && hasToken()) {
    return { path: "/" }
  }

  if (to.meta.requiresAuth && !hasToken()) {
    return {
      name: "login",
      query: { redirect: to.fullPath }
    }
  }

  return true
})
