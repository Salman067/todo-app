<template>
  <div class="todo-main">
    <!-- Toast Component -->
    <div v-if="toast.show" class="toast" :class="toast.type">
      {{ toast.message }}
    </div>
    <h1>TODO List</h1>
    <div v-if="statusMessage" class="status-message">{{ statusMessage }}</div>
    <div class="input-group">
      <button @click="openAddModal">Add Task</button>
    </div>

    <!-- Search Inputs -->
    <div class="search-group">
      <input
        v-model="searchTask"
        placeholder="Search tasks"
        @input="searchTodos"
      />
      <select v-model="searchStatus" @change="searchTodos">
        <option value="">All Statuses</option>
        <option value="created">Created</option>
        <option value="in-progress">In Progress</option>
        <option value="done">Done</option>
      </select>
    </div>

    <table v-if="paginatedTodos.length > 0" class="todo-table">
      <thead>
        <tr>
          <th>ID</th>
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
            <button class="delete-button" @click="deleteTodo(todo.ID)">
              ✖
            </button>
            <button class="view-button" @click="viewTodoDetails(todo.ID)">
              Details
            </button>
            <button class="edit-button" @click="openEditModal(todo)">
              Edit
            </button>
          </td>
        </tr>
      </tbody>
    </table>

    <div v-else>
      <p>No tasks available.</p>
    </div>

    <!-- Pagination Controls -->
    <div class="pagination">
      <button
        @click="changePage(currentPage - 1)"
        :disabled="currentPage === 1"
      >
        Previous Page
      </button>
      <span>Page {{ currentPage }} / {{ totalPages }}</span>
      <button
        @click="changePage(currentPage + 1)"
        :disabled="currentPage === totalPages"
      >
        Next Page
      </button>
    </div>

   <!-- Add Task Modal -->
   <div v-if="isAddModalOpen" class="modal">
      <div class="modal-content">
        <h2>Add New Task</h2>
        <textarea
          v-model="newTask"
          placeholder="Enter new task"
          class="edit-input"
        ></textarea>
        <div class="modal-buttons">
          <button @click="addTodo">Add</button>
          <button @click="closeAddModal">Cancel</button>
        </div>
      </div>
    </div>

    <!-- Edit Task Modal -->
    <div v-if="isEditModalOpen" class="modal">
      <div class="modal-content">
        <h2>Edit Task</h2>
        <textarea
          v-model="editTaskContent"
          placeholder="Update task"
          class="edit-input"
        ></textarea>
        <div class="modal-buttons">
          <button @click="saveEditTask">Save</button>
          <button @click="closeEditModal">Cancel</button>
        </div>
      </div>
    </div>

    <!-- Modal to Display Todo Details -->
    <div v-if="selectedTodo" class="modal">
      <div class="modal-content">
        <h2>Task Details</h2>
        <p><strong>Task:</strong> {{ selectedTodo.Task }}</p>
        <p>
          <strong>Status:</strong>
          <span :class="getStatusClass(selectedTodo.Status)">{{
            getStatusText(selectedTodo.Status)
          }}</span>
        </p>
        <p>
          <strong>Created At:</strong> {{ formatDate(selectedTodo.CreatedAt) }}
        </p>
        <p>
          <strong>Updated At:</strong> {{ formatDate(selectedTodo.UpdatedAt) }}
        </p>
        <button @click="closeTodoDetails">Close</button>
      </div>
    </div>

  </div>
</template>

<script>
export default {
  data() {
    return {
      newTask: "",
      todos: [],
      statusMessage: "",
      selectedTodo: null,
      searchTask: "",
      searchStatus: "",
      currentPage: 1,
      perPage: 10,
      totalTodos: 0,
      isAddModalOpen: false,
      toast: {
        show: false,
        message: "",
        type: "success",
        timeout: null,
      },
      editTaskContent: '',
      editTaskId: null,
      isEditModalOpen: false,
    };
  },
  computed: {
    paginatedTodos() {
      return this.todos;
    },
    totalPages() {
      return Math.ceil(this.totalTodos / this.perPage);
    },
  },
  mounted() {
    this.fetchTodos();
  },
  methods: {
    openAddModal() {
      this.isAddModalOpen = true;
    },
    closeAddModal() {
      this.isAddModalOpen = false;
      this.newTask = "";
    },
    openEditModal(todo) {
      this.editTaskContent = todo.Task;
      this.editTaskId = todo.ID;
      this.isEditModalOpen = true;
    },
    closeEditModal() {
      this.isEditModalOpen = false;
      this.editTaskContent = '';
      this.editTaskId = null;
    },
    showToast(message, type = "success", duration = 2000) {
      clearTimeout(this.toast.timeout);
      this.toast.show = true;
      this.toast.message = message;
      this.toast.type = type;
      this.toast.timeout = setTimeout(() => {
        this.toast.show = false;
      }, duration);
    },
    async fetchTodos() {
      try {
        const params = new URLSearchParams({
          page: this.currentPage.toString(),
          per_page: this.perPage.toString(),
        });

        if (this.searchTask.trim()) {
          params.append("task", this.searchTask.trim());
        }

        if (this.searchStatus) {
          params.append("status", this.searchStatus);
        }

        const url = `/api/v1/todos?${params.toString()}`;
        const response = await fetch(url);
        if (!response.ok) {
          throw new Error(
            `Failed to get todo list. statusCode: ${response.status}`
          );
        }
        const data = await response.json();
        this.todos = data.data;
        this.totalTodos = data.total;
        this.currentPage = data.page;
        this.perPage = data.per_page;
      } catch (error) {
        console.error("Error fetching todos:", error);
        this.showToast("Failed to retrieve tasks", "error");
      }
    },
    async addTodo() {
      if (this.newTask.trim() === "") return;

      try {
        const response = await fetch("/api/v1/todos", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            task: this.newTask,
            Status: "created",
          }),
        });

        if (!response.ok)
          throw new Error(
            `Failed to create todo. statusCode: ${response.status}`
          );

        this.newTask = "";
        this.showToast("Task added successfully", "success");
        this.fetchTodos();
        this.closeAddModal();
      } catch (error) {
        console.error("Error creating todo:", error);
        this.showToast("Failed to create task", "error");
      }
    },
    async saveEditTask() {
      if (!this.editTaskContent.trim()) {
        this.statusMessage = 'Please enter a task.';
        return;
      }

      try {
        const response = await fetch(`/api/v1/todos/${this.editTaskId}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            task: this.editTaskContent,
          }),
        });

        if (!response.ok)
          throw new Error(
            `Failed to edit todo. statusCode: ${response.status}`
          );

        this.showToast('Task updated successfully', 'success');
        this.fetchTodos();
        this.closeEditModal();
      } catch (error) {
        console.error('Error editing task:', error.message);
        this.showToast('Failed to update task', 'error');
      }
    },
    enableEdit(todo) {
      todo.isEditing = true;
    },
    async editTodo(todo) {
      if (!todo.Task.trim()) {
        this.statusMessage = "Please enter a task.";
        return;
      }

      todo.isEditing = false;

      try {
        const response = await fetch(`/api/v1/todos/${todo.ID}`, {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            task: todo.Task,
          }),
        });

        if (!response.ok)
          throw new Error(
            `Failed to edit todo. statusCode: ${response.status}`
          );
        this.showToast("Task edited successfully", "success");
        this.fetchTodos();
      } catch (error) {
        console.error("Error editing todo:", error);
        this.showToast("Failed to edit task", "error");
      }
    },
    cancelEdit(todo) {
      todo.isEditing = false;
    },
    async updateStatus(todo) {
      try {
        const statuses = ["created", "in-progress", "done"];
        const currentIndex = statuses.indexOf(todo.Status);
        const nextIndex = (currentIndex + 1) % statuses.length;
        const nextStatus = statuses[nextIndex];

        const response = await fetch(`/api/v1/todos/${todo.ID}`, {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            Status: nextStatus,
          }),
        });

        if (!response.ok)
          throw new Error(
            `Failed to update todo Status. statusCode: ${response.status}`
          );
        this.showToast("Task status updated", "success");
        this.fetchTodos();
      } catch (error) {
        console.error("Error updating todo Status:", error);
        this.showToast( "Failed to update status", "error" );
      }
    },
    async deleteTodo(id) {
      try {
        const response = await fetch(`/api/v1/todos/${id}`, {
          method: "DELETE",
          headers: {
            "Content-Type": "application/json",
          },
        });

        if (!response.ok)
          throw new Error(
            `Failed to delete todo. statusCode: ${response.status}`
          );
        this.showToast( "Task deleted successfully", "success" );
        if (this.todos.length === 1 && this.currentPage > 1) {
          this.changePage(this.currentPage - 1);
        } else {
          this.fetchTodos();
        }
      } catch (error) {
        console.error("Error deleting todo:", error);
        this.showToast(  "Failed to delete task", "error" );
      }
    },
    async viewTodoDetails(id) {
      try {
        const response = await fetch(`/api/v1/todos/${id}`, {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
          },
        });

        if (!response.ok)
          throw new Error(
            `Failed to fetch todo details. statusCode: ${response.status}`
          );

        const data = await response.json();
        this.selectedTodo = data.data;
      } catch (error) {
        console.error("Error fetching todo details:", error);
        this.showToast(  "Failed to retrieve task details", "error" );
      }
    },
    closeTodoDetails() {
      this.selectedTodo = null;
    },
    searchTodos() {
      this.currentPage = 1;
      this.fetchTodos();
    },
    async changePage(page) {
      if (page >= 1 && page <= this.totalPages) {
        this.currentPage = page;
        await this.fetchTodos();
      }
    },
    formatDate(dateString) {
      const options = {
        year: "numeric",
        month: "2-digit",
        day: "2-digit",
        hour: "2-digit",
        minute: "2-digit",
        second: "2-digit",
      };
      return new Date(dateString).toLocaleString("en-US", options);
    },
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
  },
};
</script>

<style scoped>
.todo-main {
  max-width: 900px;
  margin: 20px auto;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  font-family: Arial, sans-serif;
}
.input-group,
.search-group {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
}

select {
  width: 28%;
}
h1 {
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
.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 20px;
}
.pagination button {
  margin: 0 10px;
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

.modal-buttons button{
  margin-right: 10px;
}
.modal {
  position: fixed;
  z-index: 1;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: rgba(0, 0, 0, 0.4);
}
.modal-content {
  background-color: #fefefe;
  margin: 15% auto;
  padding: 20px;
  border: 1px solid #888;
  width: 80%;
  max-width: 600px;
  border-radius: 8px;
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

.toast {
  position: fixed;
  top: 20px;
  right: 20px;
  padding: 10px 20px;
  border-radius: 4px;
  color: white;
  font-weight: bold;
  z-index: 1000;
  transition: opacity 0.3s ease;
}

.toast.success {
  background-color: #28a745;
}

.toast.error {
  background-color: #dc3545;
}
</style>
