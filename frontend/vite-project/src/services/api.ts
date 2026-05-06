import axios from "axios";

const Api = "http://localhost:8080/api";

export const getStudents = async () => axios.get(`${Api}/students`);
export const getStudentById = async (id: number) => {
  const res = await axios.get(`${Api}/students/${id}`)
  return res.data
}
export const updateStudent = async (id: number, data: any ) => {
  const res =  await axios.put(`${Api}/students/${id}`, data)
  return res.data
}
export const addStudent = (data: any) => axios.post(`${Api}/students`, data);
export const runCap = () => axios.post(`${Api}/run-cap`)
export const deleteStudent = (id: number) => axios.delete(`${Api}/students/${id}`);
export const getStats = () => axios.get(`${Api}/stats/seats`)