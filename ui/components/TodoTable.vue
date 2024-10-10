<template>
    <div>
      <table class="todo-table">
        <thead>
          <tr>
            <th>タスク</th>
            <th>ステータス</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="todo in todos" :key="todo.ID">
            <TodoItem
              :todo="todo"
              @update-status="updateStatus"
              @delete-todo="deleteTodo"
              @edit-todo="enableEdit"
              @cancel-edit="cancelEdit"
            />
          </tr>
        </tbody>
      </table>
      <div v-if="todos.length === 0">
        <p>タスクがありません。</p>
      </div>
    </div>
  </template>
  
  <script>
  import TodoItem from './TodoItem.vue';
  
  export default {
    components: {
      TodoItem,
    },
    props: {
      todos: Array,
    },
    emits: ['update-status', 'delete-todo', 'edit-todo', 'cancel-edit'],
    methods: {
      updateStatus(todo) {
        this.$emit('update-status', todo);
      },
      deleteTodo(id) {
        this.$emit('delete-todo', id);
      },
      enableEdit(todo) {
        this.$emit('edit-todo', todo);
      },
      cancelEdit(todo) {
        this.$emit('cancel-edit', todo);
      },
    },
  };
  </script>
  
  <style scoped>
  .todo-table {
    width: 100%;
    border-collapse: collapse;
    margin-top: 20px;
  }
  
  .todo-table th,
  .todo-table td {
    padding: 10px;
    text-align: center; /* Center all text in the table */
    border: 1px solid #ddd;
    word-wrap: break-word; /* Ensure long words break to the next line */
    overflow: hidden; /* Prevent overflowing content */
    max-width: 150px; /* Set a max width for the cells */
  }
  
  .todo-table th {
    background-color: #f0f0f0;
  }
  
  .buttons button {
    background-color: #f0f0f0;
    color: #333;
    margin-left: 5px;
    border-radius: 4px;
    padding: 5px 10px;
    transition: background-color 0.3s, color 0.3s;
  }
  
  .buttons button.done {
    background-color: #333;
    color: #fff;
  }
  
  .buttons button.delete-button {
    color: white;
  }
  
  .done-task {
    text-decoration: line-through;
  }
  </style>
  