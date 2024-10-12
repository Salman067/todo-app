<template>
  <div class="taskme-dashboard">
    <!-- Header -->
    <header class="header">
      <div class="logo">
        <img
          src="/home/salman/Desktop/fullstack-examination-2024/ui/assets/task.png"
          alt="Task Icon"
          class="logo-img"
        >
        <h1>TaskMe</h1>
      </div>
      <div class="search-bar">
        <input
          v-model="keyWord"
          type="text"
          placeholder="Search..."
          @input="updatekeyWord($event.target.value)"
        >
      </div>
    </header>

    <!-- Sidebar -->
    <nav class="sidebar">
      <ul>
        <li>
          <a
            href="#"
            class="nav-item"
            :class="{ active: activeButton === 'Tasks' }"
            @click.prevent="
              setActiveButton('Tasks');
              getTodos('Tasks');
            "
          >
            <i class="fas fa-tasks"/> Tasks
          </a>
        </li>
        <li>
          <a
            href="#"
            class="nav-item"
            :class="{ active: activeButton === 'Completed' }"
            @click.prevent="
              setActiveButton('Completed');
              filterTasks('Completed');
            "
          >
            <i class="fas fa-check-circle"/> Completed
          </a>
        </li>
        <li>
          <a
            href="#"
            class="nav-item"
            :class="{ active: activeButton === 'In-Progress' }"
            @click.prevent="
              setActiveButton('In-Progress');
              filterTasks('In-Progress');
            "
          >
            <i class="fas fa-spinner"/> In Progress
          </a>
        </li>
        <li>
          <a
            href="#"
            class="nav-item"
            :class="{ active: activeButton === 'To-Do' }"
            @click.prevent="
              setActiveButton('To-Do');
              filterTasks('To-Do');
            "
          >
            <i class="fas fa-list"/> To Do
          </a>
        </li>
      </ul>
    </nav>

    <!-- Main Content -->
    <main class="main-content">
      <div class="content-header">
        <h2>{{ selectedStatus }} List</h2>
        <button class="create-task-btn" @click="openAddModal">
          + Create Task
        </button>
      </div>

      <div class="view-options">
        <div v-if="selectedStatus !== 'Tasks' && selectedStatus">
          <button
            class="view-btn"
            :class="{ active: viewMode === 'list' }"
            @click="viewMode = 'list'"
          >
            <i class="fas fa-list"/> List View
          </button>
          <!-- Sorting Dropdown -->
          <label for="sort-by" class="sort-label">Sort by:</label>
          <select v-model="sortBy" class="sort-dropdown" @change="sortTasks">
            <option value="Task">Task Title</option>
            <option value="PriorityTask">Priority</option>
            <option value="CreatedAt">Created Date</option>
          </select>
        </div>
        <div v-else>
          <!-- @click="viewMode = 'board'" -->
          <button
            class="view-btn"
            :class="{ active: viewMode === 'board' }"
             @click="handleViewToggle('board')"
          >
            <i class="fas fa-th-large"/> Board View
          </button>
          <button
            class="view-btn"
            :class="{ active: viewMode === 'list' }"
            @click="viewMode = 'list'"
          >
            <i class="fas fa-list"/> List View
          </button>
          <!-- Sorting Dropdown -->
          <label for="sort-by" class="sort-label">Sort by:</label>
          <select v-model="sortBy" class="sort-dropdown" @change="sortTasks">
            <option value="Task">Task Title</option>
            <option value="PriorityTask">Priority</option>
            <option value="CreatedAt">Created Date</option>
          </select>
        </div>
      </div>

      <div class="status-container">
        <div class="status-card to-do">
          <span class="status-indicator"/>
          <span class="status-title">To Do</span>
          <!-- <button class="add-button">+</button> -->
        </div>
        <div class="status-card in-progress">
          <span class="status-indicator"/>
          <span class="status-title">In Progress</span>
          <!-- <button class="add-button">+</button> -->
        </div>
        <div class="status-card completed">
          <span class="status-indicator"/>
          <span class="status-title">Completed</span>
          <!-- <button class="add-button">+</button> -->
        </div>
      </div>

      <div v-if="selectedStatus !== 'Tasks' && selectedStatus">
        <div class="list-view">
          <table>
            <thead>
              <tr>
                <th>Task Title</th>
                <th>Description</th>
                <th>Priority</th>
                <th>Created At</th>
                <th>Actions</th>
              </tr>
            </thead>
            <tbody>
              <!-- @click.stop="handleViewTodoDetails(task.ID)" -->
              <tr v-for="task in sortedTasks" :key="task.id">
                <td>
                  <div class="task-with-status">
                    <span
                      class="status-circle"
                      :class="getStatusClass(task.Status)"
                    />
                    <span class="task-title">{{
                      truncateText(task.Task)
                    }}</span>
                    <span
                      v-if="task.Task.length > 50"
                      class="expand-btn"
                      @click.stop="handleViewTodoDetails(task.ID)"
                    >
                      ...details
                    </span>
                  </div>
                </td>
                <td>
                  <span class="task-title">{{
                    truncateText(task.Description)
                  }}</span>
                  <span
                    v-if="task.Description.length > 50"
                    class="expand-btn"
                    @click.stop="handleViewTodoDetails(task.ID)"
                  >
                    ...details
                  </span>
                </td>
                <td>
                  <span
                    class="priority"
                    :class="(task.PriorityTask || 'unknown').toLowerCase()"
                  >
                    {{ task.PriorityTask.toUpperCase() || "Unknown" }}
                  </span>
                </td>
                <td>{{ formatDate(task.CreatedAt) }}</td>
                <td>
                  <button
                    class="action-btn edit-btn"
                    @click="openEditModal(task)"
                  >
                    Edit
                  </button>
                  <button
                    class="action-btn view1-btn"
                    @click="handleViewTodoDetails(task.ID)"
                  >
                    Details
                  </button>
                  <button
                    class="action-btn delete-btn"
                    @click="handleDeleteTodo(task.ID)"
                  >
                    Delete
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
          <p v-if="tasks.length === 0" class="no-task-available">
            No tasks available
          </p>
          <!-- Use the Pagination component -->
          <Pagination
            :current-page="currentPage"
            :total-pages="totalPages"
            @page-changed="changePage"
          />
        </div>
      </div>

      <div v-else>
        <div v-if="viewMode === 'board'" class="board-view">
          <div
            v-for="status in ['To-Do', 'In-Progress', 'Completed']"
            :key="status"
            class="task-column"
          >
            <h3>{{ status }} <button class="add-task-btn">+</button></h3>
            <div class="task-list">
              <div
                v-for="task in getTasksByStatus(status)"
                :key="task.ID"
                class="task-card"
                @click.stop="handleViewTodoDetails(task.ID)"
              >
                <div class="task-with-status">
                  <span
                    class="status-circle"
                    :class="getStatusClass(task.Status)"
                  />
                  <span class="task-title">{{ truncateText(task.Task) }}</span>
                  <span
                    v-if="task.Task.length > 50"
                    class="expand-btn"
                    @click.stop="handleViewTodoDetails(task.ID)"
                  >
                    ...details
                  </span>
                </div>
                <p>
                  <span class="task-title">{{
                    truncateText(task.Description)
                  }}</span>
                  <span
                    v-if="task.Description.length > 50"
                    class="expand-btn"
                    @click.stop="handleViewTodoDetails(task.ID)"
                  >
                    ...details
                  </span>
                </p>
                <div class="task-meta">
                  <span
                    class="priority"
                    :class="task.PriorityTask.toLowerCase()"
                  >
                    {{ task.PriorityTask.toUpperCase() }}
                  </span>
                  <span class="date">{{ formatDate(task.CreatedAt) }}</span>
                </div>
              </div>
              <p v-if="getTasksByStatus(status).length === 0">
                No tasks available
              </p>
            </div>
          </div>
        </div>

        <div v-else class="list-view">
          <table>
            <thead>
              <tr>
                <th>Task Title</th>
                <th>Description</th>
                <th>Priority</th>
                <th>Created At</th>
                <th>Actions</th>
              </tr>
            </thead>
            <tbody>
              <!-- @click.stop="handleViewTodoDetails(task.ID)" -->
              <tr v-for="task in sortedTasks" :key="task.id">
                <td>
                  <div class="task-with-status">
                    <span
                      class="status-circle"
                      :class="getStatusClass(task.Status)"
                    />
                    <span class="task-title">{{
                      truncateText(task.Task)
                    }}</span>
                    <span
                      v-if="task.Task.length > 50"
                      class="expand-btn"
                      @click.stop="handleViewTodoDetails(task.ID)"
                    >
                      ...details
                    </span>
                  </div>
                </td>
                <td>
                  <span class="task-title">{{
                    truncateText(task.Description)
                  }}</span>
                  <span
                    v-if="task.Description.length > 50"
                    class="expand-btn"
                    @click.stop="handleViewTodoDetails(task.ID)"
                  >
                    ...details
                  </span>
                </td>
                <td>
                  <span
                    class="priority"
                    :class="(task.PriorityTask || 'unknown').toLowerCase()"
                  >
                    {{ task.PriorityTask.toUpperCase() || "Unknown" }}
                  </span>
                </td>
                <td>{{ formatDate(task.CreatedAt) }}</td>
                <td>
                  <button
                    class="action-btn edit-btn"
                    @click="openEditModal(task)"
                  >
                    Edit
                  </button>
                  <button
                    class="action-btn view1-btn"
                    @click="handleViewTodoDetails(task.ID)"
                  >
                    Details
                  </button>
                  <button
                    class="action-btn delete-btn"
                    @click="handleDeleteTodo(task.ID)"
                  >
                    Delete
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
          <p v-if="tasks.length === 0" class="no-task-available">
            No tasks available
          </p>
          <!-- Use the Pagination component -->
          <Pagination
            :current-page="currentPage"
            :total-pages="totalPages"
            @page-changed="changePage"
          />
        </div>
      </div>
    </main>

    <!-- Add Task Modal -->
    <Modal
      v-if="isAddModalOpen"
      :is-open="isAddModalOpen"
      :title="'Create New Task'"
      @close="closeAddModal"
    >
      <form @submit.prevent="saveTask">
        <input v-model="currentTask.title" placeholder="Task Title" required >
        <textarea
          v-model="currentTask.description"
          placeholder="Task Description"
          required
        />
        <select v-model="currentTask.priority">
          <option value="Low">Low Priority</option>
          <option value="Medium">Medium Priority</option>
          <option value="High">High Priority</option>
        </select>
        <select v-model="currentTask.status">
          <option value="To-Do">To Do</option>
          <!-- <option value="In-Progress">In Progress</option>
          <option value="Completed">Completed</option> -->
        </select>
        <button type="submit">
          {{ "Create Task" }}
        </button>
      </form>
    </Modal>

    <!-- Edit Task Modal -->
    <Modal
      v-if="isEditModalOpen"
      :is-open="isEditModalOpen"
      :title="'Updated Task'"
      @close="closeEditModal"
    >
      <form @submit.prevent="saveEditTask">
        <input
          v-model="currentEditTask.Task"
          placeholder="Task Title"
          required
        >
        <textarea
          v-model="currentEditTask.Description"
          placeholder="Task Description"
          required
        />
        <select v-model="currentEditTask.PriorityTask">
          <option value="low">Low Priority</option>
          <option value="medium">Medium Priority</option>
          <option value="high">High Priority</option>
        </select>
        <select v-model="currentEditTask.Status">
          <option value="to-do">To Do</option>
          <option value="in-progress">In Progress</option>
          <option value="completed">Completed</option>
        </select>
        <button type="submit">Update Task</button>
      </form>
    </Modal>

    <!-- Modal to Display Todo Details -->
    <Modal
      v-if="selectedTodo"
      :is-open="!!selectedTodo"
      title="Task Details"
      close-button-text="Close"
      @close="closeTodoDetails"
    >
      <div class="task-details">
        <p><strong>Task Title:</strong> {{ selectedTodo.Task }}</p>
        <p><strong>Description:</strong> {{ selectedTodo.Description }}</p>
        <p>
          <strong>Priority:</strong>
          <span
            class="priority"
            :class="(selectedTodo.PriorityTask || 'unknown').toLowerCase()"
          >
            {{ selectedTodo.PriorityTask.toUpperCase() || "Unknown" }}
          </span>
        </p>
        <p>
          <strong>Status:</strong>
          <span :class="getStatusClass(selectedTodo.Status)">
            {{ getStatusText(selectedTodo.Status.toUpperCase()) }}
          </span>
        </p>
        <p>
          <strong>Created At:</strong> {{ formatDate(selectedTodo.CreatedAt) }}
        </p>
        <p>
          <strong>Updated At:</strong> {{ formatDate(selectedTodo.UpdatedAt) }}
        </p>
      </div>
    </Modal>
    <!-- Toast Notification -->
    <Toast :show="toast.show" :message="toast.message" :type="toast.type" />
  </div>
</template>

<script>
import "../style/style.css";
import Toast from "../components/NotificationToast.vue";
import Pagination from "../components/TaskPagination.vue";
import Modal from "../components/AppModal.vue";
import {
  fetchTodos,
  addTodo,
  editTodo,
  fetchTasksWithoutPagination,
  deleteTodo,
  viewTodoDetails,
  fetchTodosForStatus,
} from "../services/api";

export default {
  components: {
    Toast,
    Pagination,
    Modal,
  },

  data() {
    return {
      keyWord: "",
      searchStatus: "",
      viewMode: "list",
      activeButton: "Tasks",
      sortBy: "PriorityTask",
      tasks: [],
      isModalOpen: false,
      isEditing: false,
      isDetailsModalOpen: false,
      isAddModalOpen: false,
      isEditModalOpen: false,
      currentTask: {
        title: "",
        description: "",
        priority: "Low",
        status: "To-Do",
      },
      currentEditTask: {
        Task: "",
        Description: "",
        PriorityTask: "",
        Status: "",
      },
      toast: {
        show: false,
        message: "",
        type: "success",
        timeout: null,
      },
      currentPage: 1,
      perPage: 10,
      totalTodos: 0,
      selectedTodo: null,
      selectedStatus: "",
    };
  },

  computed: {
    paginatedTodos() {
      return this.tasks;
    },
    totalPages() {
      return Math.ceil(this.totalTodos / this.perPage);
    },
    sortedTasks() {
      return this.tasks.slice().sort((a, b) => {
        if (this.sortBy === "PriorityTask") {
          const priorityOrder = { high: 1, medium: 2, low: 3 };
          return (
            priorityOrder[a.PriorityTask.toLowerCase()] -
            priorityOrder[b.PriorityTask.toLowerCase()]
          );
        } else if (this.sortBy === "CreatedAt") {
          return new Date(b.CreatedAt) - new Date(a.CreatedAt);
        } else {
          return a.Task.localeCompare(b.Task);
        }
      });
    },
  },
  mounted() {
    this.getTodos();
  },
  methods: {
    sortTasks() {},
    setActiveButton(button) {
      this.activeButton = button;
    },
    formatDate(date) {
      const options = { year: "numeric", month: "long", day: "numeric" };
      return new Date(date).toLocaleDateString(undefined, options);
    },
    showToast(message, type = "success", duration = 2000) {
      if (this.toast.timeout) {
        clearTimeout(this.toast.timeout);
      }
      this.toast.show = true;
      this.toast.message = message;
      this.toast.type = type;
      this.toast.timeout = setTimeout(() => {
        this.toast.show = false;
      }, duration);
    },
    async changePage(page) {
      if (page >= 1 && page <= this.totalPages) {
        this.currentPage = page;
        if (this.selectedStatus !== "Tasks" && this.selectedStatus) {
          console.log("page: ", this.currentPage);
          await this.filterTasks(this.selectedStatus, "change-page");
        } else {
          await this.getTodos(this.selectedStatus);
        }
      }
    },
    openAddModal() {
      this.isAddModalOpen = true;
      this.currentTask = {
        title: "",
        description: "",
        priority: "Low",
        status: "To-Do",
      };
    },
    closeAddModal() {
      this.isAddModalOpen = false;
      this.currentTask = {
        title: "",
        description: "",
        priority: "Low",
        status: "To-Do",
      };
    },

    openEditModal(task) {
      this.isEditModalOpen = true;
      this.currentEditTask = {
        ID: task.ID,
        Task: task.Task,
        Description: task.Description,
        PriorityTask: task.PriorityTask,
        Status: task.Status,
      };
    },
    closeEditModal() {
      this.isEditModalOpen = false;
      this.currentEditTask = {
        Task: "",
        Description: "",
        PriorityTask: "",
        Status: "",
      };
    },

    async saveEditTask() {
      const updatedTask = {
        ID: this.currentEditTask.ID,
        Task: this.currentEditTask.Task,
        Description: this.currentEditTask.Description,
        PriorityTask: this.currentEditTask.PriorityTask,
        Status: this.currentEditTask.Status,
      };
      try {
        await editTodo(updatedTask);
        this.isEditModalOpen = false;
        this.getTodos();
        this.closeEditModal();
        this.showToast("Task updated successfully.");
      } catch (error) {
        console.error(error.message);
        this.showToast("Failed to update task", "error");
      }
    },

    async saveTask() {
      try {
        await addTodo(this.currentTask);
        this.currentTask = {
          title: "",
          description: "",
          priority: "Low",
          status: "To-Do",
        };
        this.showToast("Task added successfully", "success");
        this.getTodos();
        this.closeAddModal();
      } catch (error) {
        console.log(error.message);
        this.showToast("Failed to create task", "error");
      }
    },

    async getTodos(status) {
      this.selectedStatus = status;
      try {
        const data = await fetchTodos(
          this.currentPage,
          this.perPage,
          this.keyWord
        );

        this.tasks = data.data;
        this.totalTodos = data.total;
        this.currentPage = data.page;
        this.perPage = data.per_page;
      } catch (error) {
        console.log(error.message);
        this.showToast("Failed to retrieve tasks", "error");
      }
    },

    getTasksByStatus(status) {
      return this.tasks.filter(
        (task) => task.Status.toLowerCase() === status.toLowerCase()
      );
    },

    getStatusClass(status) {
      return {
        "status-todo": status === "to-do",
        "status-inprogress": status === "in-progress",
        "status-completed": status === "completed",
      };
    },

  //   showToast(message, type = "success", duration = 2000) {
  //   if (this.toast.timeout) {
  //     clearTimeout(this.toast.timeout); 
  //   }

  //   this.toast.show = true;
  //   this.toast.message = message;
  //   this.toast.type = type;

  //   this.toast.timeout = setTimeout(() => {
  //     this.toast.show = false;
  //   }, duration);
  // },

    truncateText(text, length = 50) {
      if (!text) return "";
      return text.length > length ? text.slice(0, length) : text;
    },

    getStatusText(status) {
      const statusTexts = {
        "to-do": "To-Do",
        "in-progress": "In Progress",
        completed: "Completed",
      };
      return statusTexts[status] || status;
    },

    async handleDeleteTodo(id) {
      try {
        await deleteTodo(id);
        this.showToast("Task deleted successfully", "success");
        this.getTodos();
      } catch (error) {
        console.error("Error deleting task:", error);
        this.showToast("Failed to delete task", "error");
      }
    },
    async handleViewTodoDetails(id) {
      try {
        this.selectedTodo = await viewTodoDetails(id);
        this.isDetailsModalOpen = true;
        console.log("selected", this.selectedTodo);
      } catch (error) {
        console.log("error", error);
        this.showToast("Failed to retrieve task details", "error");
      }
    },
    closeTodoDetails() {
      this.selectedTodo = null;
      this.isDetailsModalOpen = false;
    },

    async filterTasks(status, pageCalulate) {
      this.selectedStatus = status;
      if (pageCalulate !== "change-page") {
        this.currentPage = 1;
      }
      try {
        const data = await fetchTodosForStatus(
          this.currentPage,
          this.perPage,
          this.selectedStatus.toLowerCase()
        );
        this.tasks = data.data;
        this.totalTodos = data.total;
        this.currentPage = data.page;
        this.perPage = data.per_page;
        console.log("fatch current page", this.currentPage);
      } catch (error) {
        console.log(error.message);
        this.showToast("Failed to retrieve tasks", "error");
      }
    },

    updatekeyWord(newValue) {
      this.keyWord = newValue;
      this.searchTodos();
    },
    searchTodos() {
      this.currentPage = 1;
      this.getTodos();
    },

    async fetchTasksWithoutPaginations(status) {
      this.selectedStatus = status;
      try {
        const data = await fetchTasksWithoutPagination();
        this.tasks = data.data;
      } catch (error) {
        console.log(error.message);
        this.showToast("Failed to retrieve tasks", "error");
      }
    },
    handleViewToggle(view) {
      this.viewMode = view;
      if (view === 'board') {
        this.fetchTasksWithoutPaginations(); 
      }
    },
  },
};
</script>
