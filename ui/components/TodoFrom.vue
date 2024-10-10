<template>
    <div class="input-group">
      <input
        :value="newTask"
        placeholder="新しいタスクを入力"
        @input="updateTask" 
        @keyup.enter="addTodo"
      />
      <button @click="addTodo">追加</button>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      value: String, // Assuming you are passing the new task value as a prop
    },
    data() {
      return {
        newTask: this.value, // Initialize local data with prop value
      };
    },
    methods: {
      updateTask(event) {
        this.newTask = event.target.value; // Update local data
        this.$emit('update:value', this.newTask); // Emit change event
      },
      addTodo() {
        // Emit the current newTask value to the parent component
        this.$emit('add-todo', this.newTask);
        this.newTask = ''; // Clear input after adding
      },
    },
    watch: {
      value(newValue) {
        this.newTask = newValue; // Update local data when prop changes
      },
    },
  };
  </script>
  
  <style scoped>
  .input-group {
    display: flex;
    margin-bottom: 20px;
  }
  
  input {
    flex: 1;
    padding: 10px;
    border: 1px solid #ddd;
    border-radius: 4px;
    margin-right: 10px;
  }
  </style>
  