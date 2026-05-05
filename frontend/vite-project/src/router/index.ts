import { createRouter, createWebHistory } from "vue-router"
import Home from "../pages/Home.vue"
import StudentList from "../pages/StudentList.vue"
import AddStudent from "../pages/AddStudent.vue"
import ViewStudent from "../pages/viewStudent.vue"

const routes = [
  { path: "/", name: "home", component: Home },
  { path: "/students", name: "students", component: StudentList },
  { path: "/add", name: "add-student", component: AddStudent },
  { path: "/students/:id/edit", name: "edit-student", component: AddStudent },
  {
    path: "/students/:id",
    name: "view-student",
    component: ViewStudent,
    props: true
  }
]

export const router = createRouter({
  history: createWebHistory(),
  routes
})
