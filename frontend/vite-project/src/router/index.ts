import { createRouter, createWebHistory } from "vue-router"
import StudentList from "../pages/StudentList.vue"
import AddStudent from "../pages/AddStudent.vue"

const routes = [
  { path: "/", component: StudentList, alias: "/students" },
  { path: "/add", component: AddStudent }
]

export const router = createRouter({
  history: createWebHistory(),
  routes
})
