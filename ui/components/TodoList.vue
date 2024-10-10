<template>
  <table v-if="paginatedTodos.length > 0" class="todo-table">
    <thead>
      <tr>
        <th>SL.No</th>
        <th>Task</th>
        <th>Status</th>
        <th>Actions</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(todo, i) in paginatedTodos" :key="todo.ID">
        <td>{{ (currentPage - 1) * 10 + i + 1 }}</td>
        <td>
          <div
            :class="{
              'done-task': todo.Status === 'done',
              'task-content': true,
            }"
          >
            {{ truncateText(todo.Task) }}
            <span
              v-if="todo.Task.length > 50"
              class="expand-btn"
              @click.stop="viewTodoDetails(todo.ID)"
            >
              ...details
            </span>
          </div>
        </td>
        <td class="status">
          <span :class="getStatusClass(todo.Status)">{{
            getStatusText(todo.Status)
          }}</span>
        </td>
        <td class="buttons">
          <button
            :class="{ done: todo.Status === 'done' }"
            @click="updateStatus(todo)"
          >
            ✓
          </button>
          <button class="delete-button" @click="deleteTodo(todo.ID)">✖</button>
          <button class="view-button" @click="viewTodoDetails(todo.ID)">
            Details
          </button>
          <button class="edit-button" @click="openEditModal(todo)">Edit</button>
        </td>
      </tr>
    </tbody>
  </table>
  <div v-else>
    <p>No tasks available.</p>
  </div>
</template>

<script>
export default {
  props: {
    paginatedTodos: Array,
    currentPage: Number,
  },
  methods: {
    truncateText(text, length = 50) {
      return text.length > length ? text.slice(0, length) : text;
    },
    getStatusClass(status) {
      return {
        "status-created": status === "created",
        "status-in-progress": status === "in-progress",
        "status-done": status === "done",
      };
    },
    getStatusText(status) {
      const statusTexts = {
        created: "Created",
        "in-progress": "In Progress",
        done: "Done",
      };
      return statusTexts[status] || status;
    },
    updateStatus(todo) {
      this.$emit("update-status", todo);
    },
    deleteTodo(id) {
      this.$emit("delete-todo", id);
    },
    viewTodoDetails(id) {
      this.$emit("view-todo-details", id);
    },
    openEditModal(todo) {
      this.$emit("edit-todo", todo);
    },
  },
};
</script>

<style scoped>
.input-group,
.search-group {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
}

select {
  width: 28%;
}
p{
  text-align: center;
}
button {
  padding: 10px 15px;
  border: none;
  border-radius: 4px;
  background-color: #007bff;
  color: white;
  cursor: pointer;
  font-size: 10px;
  transition: background-color 0.3s;
}

.buttons button {
  margin-right: 3px;
}
button:hover {
  background-color: #0056b3;
}
button.done {
  background-color: #28a745;
}
button.delete-button {
  background-color: #dc3545;
}
button.view-button {
  background-color: #17a2b8;
}
.status-message {
  color: #dc3545;
  margin-bottom: 10px;
}
.done-task {
  text-decoration: line-through;
  color: #6c757d;
}

.todo-table {
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 20px;
  table-layout: fixed;
}

.todo-table th,
.todo-table td {
  border: 1px solid #ddd;
  padding: 12px;
  text-align: center;
  word-wrap: break-word;
  overflow-wrap: break-word;
}
.todo-table th {
  background-color: #f8f9fa;
  font-weight: bold;
}

.todo-table tr:nth-child(even) {
  background-color: #f2f2f2;
}

.todo-table th,
.todo-table td {
  width: 25%;
}
.task-content {
  max-width: 300px;
  white-space: normal;
  overflow: hidden;
  text-overflow: ellipsis;
}
.expand-btn {
  color: #007bff;
  cursor: pointer;
  margin-left: 5px;
}
.expand-btn:hover {
  text-decoration: underline;
}

.status-created {
  background-color: #ffc107;
  color: #000;
  padding: 3px 8px;
  border-radius: 12px;
  font-size: 12px;
}
.status-in-progress {
  background-color: #17a2b8;
  color: #fff;
  padding: 3px 8px;
  border-radius: 12px;
  font-size: 12px;
}
.status-done {
  background-color: #28a745;
  color: #fff;
  padding: 3px 8px;
  border-radius: 12px;
  font-size: 12px;
}
.edit-input {
  width: 100%;
  min-height: 60px;
  resize: vertical;
}
</style>
