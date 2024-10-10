import axios from "axios";

const API_URL = "http://localhost:8080/api/v1/todos";

export const getTodos = () => axios.get(API_URL);

export const addTodo = (task) =>
  axios.post(API_URL, {
    task,
    Status: "created",
  });

export const updateTodo = (id, updates) =>
  axios.put(`${API_URL}/${id}`, updates);

export const deleteTodo = (id) => axios.delete(`${API_URL}/${id}`);
