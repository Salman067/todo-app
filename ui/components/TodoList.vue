<template>
  <div>
    <div class="status-container">
      <div class="status-card to-do">
        <span class="status-indicator"/>
        <span class="status-title">To Do</span>
        <button class="add-button">+</button>
      </div>
      <div class="status-card in-progress">
        <span class="status-indicator"/>
        <span class="status-title">In Progress</span>
        <button class="add-button">+</button>
      </div>
      <div class="status-card completed">
        <span class="status-indicator"/>
        <span class="status-title">Completed</span>
        <button class="add-button">+</button>
      </div>
    </div>

    <table v-if="paginatedTodos.length > 0" class="todo-table">
      <thead>
        <tr>
          <th>Task Title</th>
          <th>Priority</th>
          <th>Created At</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="todo in paginatedTodos" :key="todo.ID">
          <td>
            <div class="task-with-status">
              <span :class="getStatusClass(todo.Status)" class="status-circle"/>
              <span class="task-title">{{ truncateText(todo.Task) }}</span>
              <span
                v-if="todo.Task.length > 50"
                class="expand-btn"
                @click.stop="viewTodoDetails(todo.ID)"
              >
                ...details
              </span>
            </div>
          </td>
          <td class="priority-task">
            <span :class="getPriorityClass(todo.PriorityTask)">
              {{ getPriorityText(todo.PriorityTask) }}
            </span>
          </td>
          <td class="created-at">{{ formatCreatedAt(todo.CreatedAt) }}</td>
          <td class="buttons">
            <button :class="{ done: todo.Status === 'done' }" @click="updateStatus(todo)">
              ✓
            </button>
            <button class="delete-button" @click="deleteTodo(todo.ID)">✖</button>
            <button class="view-button" @click="viewTodoDetails(todo.ID)">Details</button>
            <button class="edit-button" @click="openEditModal(todo)">Edit</button>
          </td>
        </tr>
      </tbody>
    </table>
    <div v-else>
      <p>No tasks available.</p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    paginatedTodos: {
      type: Array,
      default: () => [],  // Default value for paginatedTodos
    },
    currentPage: {
      type: Number,
      default: 1,  // Default value for currentPage
    },
  },
  emits: ['update-status', 'delete-todo', 'view-todo-details', 'edit-todo'],  // Declare all emitted events
  methods: {
    truncateText(text, length = 50) {
      return text.length > length ? text.slice(0, length) : text;
    },
    getPriorityClass(PriorityTask) {
      return {
        "priority-low": PriorityTask === "low",
        "priority-medium": PriorityTask === "medium",
        "priority-high": PriorityTask === "high",
      };
    },
    getPriorityText(PriorityTask) {
      const priorityTexts = {
        low: "Low",
        medium: "Medium",
        high: "High",
      };
      return priorityTexts[PriorityTask] || PriorityTask;
    },
    getStatusClass(status) {
      return {
        "status-created": status === "created",
        "status-in-progress": status === "in-progress",
        "status-done": status === "done",
      };
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
    formatCreatedAt(date) {
      if (!date) return '';
      const options = { day: 'numeric', month: 'short', year: 'numeric' };
      return new Date(date).toLocaleDateString('en-GB', options).replace(/ /g, '-');  // Format to "9-Feb-2024"
    },
  },
};
</script>

<style scoped>
.todo-table {
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 20px;
  table-layout: fixed; /* Ensures columns obey the set width */
}

.todo-table th,
.todo-table td {
  border: 1px solid #ddd;
  padding: 12px;
  text-align: center;
  word-wrap: break-word;
  overflow-wrap: break-word;
}

.todo-table th:first-child,
.todo-table td:first-child {
  width: 25%;
}

.todo-table th:not(:first-child),
.todo-table td:not(:first-child) {
  width: 10%;
}

.todo-table th {
  background-color: #f8f9fa;
  font-weight: bold;
}

.todo-table tr:nth-child(even) {
  background-color: #f2f2f2;
}

/* Task content style */
.task-with-status {
  display: flex;
  align-items: center;
}

.task-title {
  margin-left: 10px;
}

.expand-btn {
  color: #007bff;
  cursor: pointer;
  margin-left: 5px;
}

/* Priority styles */
.priority-medium {
  background-color: #ffc107;
  color: #000;
  padding: 3px 8px;
  border-radius: 100px;
  font-size: 14px;
}

.priority-high {
  background-color: #ff7f00;
  color: #fff;
  padding: 3px 8px;
  border-radius: 100px;
  font-size: 14px;
}

.priority-low {
  background-color: #696969;
  color: #fff;
  padding: 3px 8px;
  border-radius: 100px;
  font-size: 14px;
}

/* Status circle styles */
.status-circle {
  width: 12px;
  height: 12px;
  border-radius: 50%;
}

.status-created {
  background-color: #007bff;
}

.status-in-progress {
  background-color: #ffc107;
}

.status-done {
  background-color: #28a745;
}

/* Button styles */
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

button:hover {
  background-color: #0056b3;
}

button.done {
  background-color: #28a745;
}

.buttons button {
  margin-right: 5px;
}

button.delete-button {
  background-color: #dc3545;
}

button.edit-button {
  padding: 10px;
}

button.view-button {
  background-color: #17a2b8;
  padding: 10px;
}

.status-container {
  display: flex;
  justify-content: space-between;
  padding: 20px;
  background-color: #f8f9fa;
}

.status-card {
  display: flex;
  align-items: center;
  padding: 15px;
  border-radius: 8px;
  border: 1px solid #ddd;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  background-color: #fff;
  flex: 1;
  margin: 0 10px;
}

.status-card .status-indicator {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  margin-right: 10px;
}

.to-do .status-indicator {
  background-color: #007bff;
}

.in-progress .status-indicator {
  background-color: #ffc107;
}

.completed .status-indicator {
  background-color: #28a745;
}

.status-title {
  flex: 1;
  font-size: 16px;
  font-weight: 600;
}

.add-button {
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  padding: 5px 10px;
  cursor: pointer;
  font-size: 14px;
  transition: background-color 0.3s;
}

.add-button:hover {
  background-color: #0056b3;
}
</style>
