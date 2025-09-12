<template>
  <div class="todo-form">
    <input
        type="text"
        v-model="newTodoTitle"
        @keyup.enter="addTodo"
        placeholder="添加新的待办事项..."
        class="todo-input"
    >
    <button @click="addTodo" class="add-button">添加</button>
  </div>
</template>

<script>
export default {
  name: 'TodoForm',
  data() {
    return {
      newTodoTitle: ''
    }
  },
  methods: {
    addTodo() {
      if (this.newTodoTitle.trim()) {
        this.$emit('add-todo', this.newTodoTitle.trim())
        this.newTodoTitle = ''
      }
    }
  }
}
</script>

<style scoped>
.todo-form {
  display: flex;
  margin-bottom: 20px;
}

.todo-input {
  flex: 1;
  padding: 10px;
  font-size: 16px;
  border: 1px solid #ddd;
  border-radius: 4px 0 0 4px;
}

.add-button {
  padding: 10px 20px;
  background-color: #42b983;
  color: white;
  border: none;
  border-radius: 0 4px 4px 0;
  cursor: pointer;
  font-size: 16px;
}

.add-button:hover {
  background-color: #359e75;
}
</style>
