import axios from "axios"

const api = axios.create({
  baseURL: "http://localhost:3000"
})

api.interceptors.request.use((config) => {
  const token = localStorage.getItem("token")

  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }

  return config
})

export const setToken = (token: string) => {
  localStorage.setItem("token", token)
}

export const clearToken = () => {
  localStorage.removeItem("token")
}

export const hasToken = () => Boolean(localStorage.getItem("token"))

export const getStudents = async () => api.get("/r/students")
export const getStudentById = async (id: string) => {
  const res = await api.get(`/r/students/${id}`)
  console.log(res.data)
  return res.data[0]
}
export const updateStudent = async (id: string, data: any) => {
  const res = await api.put(`/r/students/${id}`, data)
  return res.data
}
export const addStudent = (data: any) => api.post("/r/students", data)
export const runCap = () => api.post("/r/run-cap")
export const deleteStudent = (id: string) => api.delete(`/r/students/${id}`)
export const getStats = () => api.get("/r/stats/seats")

// Authentication API
export const login = async (data: any) => {
  const res = await api.post("/login", data)
  return res.data
}

/// Branches

export const getBranches = async () => {
  const res = await api.get("/branches")
  return res.data
}

export const getBranch = async (id: string) => {
  const res = await api.get(`/branches/${id}`)
  return res.data
}
export const updateBranch = async (id: string, data: any) => {
  const res = await api.put(`/branches/${id}`, data)
  return res.data
}
export const addBranch = (data: any) => api.post("/branches", data)
export const deleteBranch = (id: string) => api.delete(`/branches/${id}`)
