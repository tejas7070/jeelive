import { createRouter, createWebHistory } from "vue-router"
import StudentList from "../pages/StudentList.vue"
import AddStudent from "../pages/AddStudent.vue"
import ViewStudent from "../pages/viewStudent.vue"

const routes = [
  { path: "/", component: StudentList, alias: "/students" },
  { path: "/add", component: AddStudent },
  {
    path: "/students/:id",
    name: "view-student",
    component: () => ViewStudent,
    props: true
  }
]

export const router = createRouter({
  history: createWebHistory(),
  routes
})
