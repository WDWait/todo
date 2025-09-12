<template>
  <div class="todo-app">
    <h1>待办事项</h1>
    <TodoForm @add-todo="addTodo" />

    <div class="todo-filters">
      <button
          :class="{ 'active': filter === 'all' }"
          @click="filter = 'all'"
          class="filter-button"
      >
        全部
      </button>
      <button
          :class="{ 'active': filter === 'active' }"
          @click="filter = 'active'"
          class="filter-button"
      >
        未完成
      </button>
      <button
          :class="{ 'active': filter === 'completed' }"
          @click="filter = 'completed'"
          class="filter-button"
      >
        已完成
      </button>
    </div>

    <div class="todo-list">
      <TodoItem
          v-for="todo in filteredTodos"
          :key="todo.id"
          :todo="todo"
          @toggle-completion="toggleTodoCompletion"
          @update-todo="updateTodo"
          @delete-todo="deleteTodo"
      />
    </div>

    <div class="todo-stats" v-if="todos.length > 0">
      <span>{{ remainingCount }} 项未完成</span>
      <button
          @click="clearCompleted"
          class="clear-button"
          v-if="completedCount > 0"
      >
        清除已完成
      </button>
    </div>
  </div>
</template>

<script>
import TodoForm from '../components/TodoForm.vue'
import TodoItem from '../components/TodoItem.vue'
import todoService from '../services/todoService'

export default {
  name: 'TodoList',
  components: {
    TodoForm,
    TodoItem
  },
  data() {
    return {
      todos: [],
      filter: 'all'
    }
  },
  computed: {
    filteredTodos() {
      switch (this.filter) {
        case 'active':
          return this.todos.filter(todo => !todo.completed)
        case 'completed':
          return this.todos.filter(todo => todo.completed)
        default:
          return this.todos
      }
    },
    remainingCount() {
      return this.todos.filter(todo => !todo.completed).length
    },
    completedCount() {
      return this.todos.filter(todo => todo.completed).length
    }
  },
  created() {
    this.fetchTodos()
  },
  methods: {
    async fetchTodos() {
      try {
        const response = await todoService.getAllTodos()
        this.todos = response.data
      } catch (error) {
        console.error('Error fetching todos:', error)
      }
    },

    async addTodo(title) {
      try {
        const response = await todoService.createTodo(title)
        const newTodo = {
          id: response.data.id,
          title,
          completed: false,
          created_at: new Date().toISOString()
        }
        this.todos.unshift(newTodo)
      } catch (error) {
        console.error('Error adding todo:', error)
      }
    },

    async toggleTodoCompletion(id, completed) {
      try {
        await todoService.toggleTodoCompletion(id, completed)
        const todo = this.todos.find(t => t.id === id)
        if (todo) {
          todo.completed = completed
        }
      } catch (error) {
        console.error('Error toggling todo:', error)
      }
    },

    async updateTodo(updatedTodo) {
      try {
        await todoService.updateTodo(updatedTodo.id, updatedTodo)
        const index = this.todos.findIndex(t => t.id === updatedTodo.id)
        if (index !== -1) {
          this.todos[index] = updatedTodo
        }
      } catch (error) {
        console.error('Error updating todo:', error)
      }
    },

    async deleteTodo(id) {
      try {
        await todoService.deleteTodo(id)
        this.todos = this.todos.filter(todo => todo.id !== id)
      } catch (error) {
        console.error('Error deleting todo:', error)
      }
    },

    async clearCompleted() {
      try {
        const completedTodos = this.todos.filter(todo => todo.completed)
        for (const todo of completedTodos) {
          await todoService.deleteTodo(todo.id)
        }
        this.todos = this.todos.filter(todo => !todo.completed)
      } catch (error) {
        console.error('Error clearing completed todos:', error)
      }
    }
  }
}
</script>

<style scoped>
.todo-app {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
  font-family: Arial, sans-serif;
}

h1 {
  text-align: center;
  color: #333;
}

.todo-filters {
  display: flex;
  justify-content: center;
  margin-bottom: 20px;
}

.filter-button {
  padding: 8px 16px;
  margin: 0 5px;
  background-color: #f5f5f5;
  border: 1px solid #ddd;
  border-radius: 4px;
  cursor: pointer;
}

.filter-button.active {
  background-color: #42b983;
  color: white;
  border-color: #42b983;
}

.todo-list {
  border: 1px solid #ddd;
  border-radius: 4px;
  margin-bottom: 20px;
}

.todo-stats {
  display: flex;
  justify-content: space-between;
  color: #666;
  font-size: 14px;
}

.clear-button {
  padding: 5px 10px;
  background-color: #e74c3c;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.clear-button:hover {
  background-color: #c0392b;
}
</style>
