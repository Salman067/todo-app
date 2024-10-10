<template>
  <div class="todo-main">
    <h1>TODO List</h1>
    <div class="input-group">
      <button @click="openAddModal">Add Task</button>
    </div>

    <!-- Use the Toast component -->
    <Toast :show="toast.show" :message="toast.message" :type="toast.type" />
    
    <!-- Search Component -->
    <SearchControls
      :searchTask="searchTask"
      :searchStatus="searchStatus"
      @updateSearchTask="updateSearchTask"
      @updateSearchStatus="updateSearchStatus"
    />

    <!-- Task List Component -->
    <TodoList
      :paginatedTodos="paginatedTodos"
      :currentPage="currentPage"
      @update-status="updateStatus"
      @delete-todo="handleDeleteTodo"
      @view-todo-details="handleViewTodoDetails"
      @edit-todo="openEditModal"
    />

    <!-- Use the Pagination component -->
    <Pagination
      :currentPage="currentPage"
      :totalPages="totalPages"
      @page-changed="changePage"
    />

    <!-- Add Task Modal -->
    <Modal
      v-if="isAddModalOpen"
      :isOpen="isAddModalOpen"
      title="Add New Task"
      primaryActionText="Add"
      closeButtonText="Cancel"
      @primaryAction="handleAddTodo"
      @close="closeAddModal"
    >
      <textarea
        v-model="newTask"
        placeholder="Enter new task"
        class="edit-input"
      ></textarea>
    </Modal>

    <!-- Edit Task Modal -->
    <Modal
      v-if="isEditModalOpen"
      :isOpen="isEditModalOpen"
      title="Edit Task"
      primaryActionText="Save"
      closeButtonText="Cancel"
      @primaryAction="saveEditTask"
      @close="closeEditModal"
    >
      <textarea
        v-model="editTaskContent"
        placeholder="Update task"
        class="edit-input"
      ></textarea>
    </Modal>

    <!-- Modal to Display Todo Details -->
    <Modal
      v-if="selectedTodo"
      :isOpen="!!selectedTodo"
      title="Task Details"
      primaryActionText="Close"
      :closeButtonText="''"
      @primaryAction="closeTodoDetails"
    >
      <p><strong>Task:</strong> {{ selectedTodo.Task }}</p>
      <p>
        <strong>Status:</strong>
        <span :class="getStatusClass(selectedTodo.Status)">
          {{ getStatusText(selectedTodo.Status) }}
        </span>
      </p>
      <p>
        <strong>Created At:</strong> {{ formatDate(selectedTodo.CreatedAt) }}
      </p>
      <p>
        <strong>Updated At:</strong> {{ formatDate(selectedTodo.UpdatedAt) }}
      </p>
    </Modal>
  </div>
</template>

<script>
import Toast from "../components/Toast.vue";
import Pagination from "../components/Pagination.vue";
import SearchControls from "../components/SearchControls.vue";
import TaskList from "../components/TodoList.vue";
import Modal from "../components/Modal.vue";
import { fetchTodos, addTodo, editTodo, updateTodoStatus, deleteTodo, viewTodoDetails} from "../services/api";

export default {
  components: {
    Toast,
    Pagination,
    SearchControls,
    TaskList,
    Modal,
  },

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
      toast: {
        show: false,
        message: "",
        type: "success",
        timeout: null,
      },
      isAddModalOpen: false,
      editTaskContent: "",
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
    this.getTodos();
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
      this.editTaskContent = "";
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

    enableEdit(todo) {
      todo.isEditing = true;
    },

    cancelEdit(todo) {
      todo.isEditing = false;
    },

    closeTodoDetails() {
      this.selectedTodo = null;
    },

    updateSearchTask(newValue) {
      this.searchTask = newValue;
      this.searchTodos();
    },

    updateSearchStatus(newValue) {
      this.searchStatus = newValue;
      this.searchTodos();
    },

    searchTodos() {
      this.currentPage = 1;
      this.getTodos();
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

    async changePage(page) {
      if (page >= 1 && page <= this.totalPages) {
        this.currentPage = page;
        await this.getTodos();
      }
    },

    async getTodos() {
      try {
        const data = await fetchTodos(this.currentPage, this.perPage, this.searchTask, this.searchStatus);
        this.todos = data.data;
        this.totalTodos = data.total;
        this.currentPage = data.page;
        this.perPage = data.per_page;
      } catch (error) {
        this.showToast("Failed to retrieve tasks", "error");
      }
    },

    async handleAddTodo() {
      try {
        await addTodo(this.newTask);
        this.newTask = "";
        this.showToast("Task added successfully", "success");
        this.getTodos(); 
        this.closeAddModal();
      } catch (error) {
        this.showToast("Failed to create task", "error");
      }
    },

    async saveEditTask() {
      try {
        await editTodo(this.editTaskId, this.editTaskContent);
        this.showToast("Task updated successfully", "success");
        this.getTodos(); 
        this.closeEditModal();
      } catch (error) {
        this.showToast("Failed to update task", "error");
      }
    },

    async updateStatus(todo) {
      try {
        await updateTodoStatus(todo);
        this.showToast("Task status updated", "success");
        this.getTodos(); 
      } catch (error) {
        this.showToast("Failed to update status", "error");
      }
    },

    async handleDeleteTodo(id) {
      try {
        await deleteTodo(id);
        this.showToast("Task deleted successfully", "success");
        if (this.todos.length === 1 && this.currentPage > 1) {
          this.changePage(this.currentPage - 1); 
        } else {
          this.getTodos(); 
        }
      } catch (error) {
        this.showToast("Failed to delete task", "error");
      }
    },

    async handleViewTodoDetails(id) {
      try {
        this.selectedTodo = await viewTodoDetails(id);
      } catch (error) {
        this.showToast("Failed to retrieve task details", "error");
      }
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

button:hover {
  background-color: #0056b3;
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
