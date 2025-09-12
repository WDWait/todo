import axios from 'axios'

export default {
    // 获取所有待办事项
    getAllTodos() {
        return axios.get('/todos')
    },

    // 根据ID获取待办事项
    getTodoById(id) {
        return axios.get(`/todos/${id}`)
    },

    // 创建新的待办事项
    createTodo(title) {
        return axios.post('/todos', { title })
    },

    // 更新待办事项
    updateTodo(id, todo) {
        return axios.put(`/todos/${id}`, todo)
    },

    // 切换待办事项的完成状态
    toggleTodoCompletion(id, completed) {
        return axios.patch(`/todos/${id}/toggle`, { completed })
    },

    // 删除待办事项
    deleteTodo(id) {
        return axios.delete(`/todos/${id}`)
    }
}
