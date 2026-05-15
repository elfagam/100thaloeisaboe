import axios from 'axios'

const API_BASE_URL = 'http://localhost:8081/api'

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json'
  }
})

export const apiService = {
  async getBooks() {
    const response = await api.get('/books')
    return response.data
  },
  async getBookById(id) {
    const response = await api.get(`/books/${id}`)
    return response.data
  },
  async register(email, password) {
    const response = await api.post('/auth/register', { email, password })
    return response.data
  },
  async activate(code) {
    const response = await api.post('/auth/activate', { code })
    return response.data
  },
  async login(email, password) {
    const response = await api.post('/auth/login', { email, password })
    return response.data
  }
}
