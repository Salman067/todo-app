<template>
  <tr>
    <td>
      <input
        v-if="todo.isEditing"
        v-model="todo.Task"
        class="edit-input"
        @blur="editTodo"
        @keyup.enter="editTodo"
        @keyup.esc="cancelEdit"
      />
      <span
        v-else
        :class="{ 'done-task': todo.Status === 'done' }"
        @click="enableEdit"
      >{{ todo.Task }}</span>
    </td>
    <td class="status">{{ todo.Status }}</td>
    <td class="buttons">
      <button
        :class="{ done: todo.Status === 'done' }"
        @click="updateStatus"
      >
        ✔️
      </button>
      <button class="delete-button" @click="deleteTodo">🗑️</button>
    </td>
  </tr>
</template>

<script>
export default {
  props: {
    todo: Object,
  },
  emits: ['update-status', 'delete-todo', 'edit-todo', 'cancel-edit'],
  methods: {
    enableEdit() {
      this.$emit('edit-todo', this.todo);
    },
    editTodo() {
      if (!this.todo.Task.trim()) {
        return;
      }
      this.$emit('edit-todo', this.todo);
    },
    cancelEdit() {
      this.$emit('cancel-edit', this.todo);
    },
    updateStatus() {
      this.$emit('update-status', this.todo);
    },
    deleteTodo() {
      this.$emit('delete-todo', this.todo.ID);
    },
  },
};
</script>

<style scoped>
.edit-input {
  flex: 1;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
}
</style>
