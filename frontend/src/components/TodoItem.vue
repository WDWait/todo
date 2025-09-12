<template>
  <div class="todo-item">
    <input
        type="checkbox"
        v-model="todo.completed"
        @change="toggleCompletion"
        class="todo-checkbox"
    >
    <span
        :class="{ 'completed': todo.completed }"
        @dblclick="startEditing"
        v-if="!isEditing"
        class="todo-title"
    >
      {{ todo.title }}
    </span>
    <input
        type="text"
        v-model="editedTitle"
        @blur="finishEditing"
        @keyup.enter="finishEditing"
        @keyup.esc="cancelEditing"
        v-if="isEditing"
        ref="editInput"
        class="edit-input"
    >
    <div class="todo-actions">
      <button @click="startEditing" class="edit-button">编辑</button>
      <button @click="deleteTodo" class="delete-button">删除</button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'TodoItem',
  props: {
    todo: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      isEditing: false,
      editedTitle: ''
    }
  },
  methods: {
    toggleCompletion() {
      this.$emit('toggle-completion', this.todo.id, this.todo.completed)
    },
    deleteTodo() {
      this.$emit('delete-todo', this.todo.id)
    },
    startEditing() {
      this.isEditing = true
      this.editedTitle = this.todo.title
      // 确保输入框在编辑模式下获得焦点
      this.$nextTick(() => {
        this.$refs.editInput.focus()
      })
    },
    finishEditing() {
      if (!this.isEditing) return

      const title = this.editedTitle.trim()
      if (title) {
        this.$emit('update-todo', {
          ...this.todo,
          title
        })
      } else {
        this.deleteTodo()
      }
      this.isEditing = false
    },
    cancelEditing() {
      this.isEditing = false
    }
  }
}
</script>

<style scoped>
.todo-item {
  display: flex;
  align-items: center;
  padding: 10px;
  border-bottom: 1px solid #eee;
}

.todo-checkbox {
  margin-right: 10px;
  width: 18px;
  height: 18px;
}

.todo-title {
  flex: 1;
  font-size: 16px;
  transition: color 0.2s;
}

.completed {
  color: #888;
  text-decoration: line-through;
}

.edit-input {
  flex: 1;
  padding: 8px;
  font-size: 16px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.todo-actions {
  margin-left: 10px;
}

.edit-button, .delete-button {
  padding: 5px 10px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
  margin-left: 5px;
}

.edit-button {
  background-color: #42b983;
  color: white;
}

.edit-button:hover {
  background-color: #359e75;
}

.delete-button {
  background-color: #e74c3c;
  color: white;
}

.delete-button:hover {
  background-color: #c0392b;
}
</style>
