<template>
  <div v-if="isOpen" class="modal">
    <div class="modal-content">
      <h2>{{ title }}</h2>
      <slot></slot> <!-- Slot for the content -->
      <div class="modal-buttons">
        <button @click="onPrimaryAction">{{ primaryActionText }}</button>
        <!-- Conditionally display the Close button -->
        <button v-if="title !== 'Task Details'" @click="onClose">{{ closeButtonText }}</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    isOpen: {
      type: Boolean,
      required: true,
    },
    title: {
      type: String,
      required: true,
    },
    primaryActionText: {
      type: String,
      default: "Confirm",
    },
    closeButtonText: {
      type: String,
      default: "Cancel",
    },
  },
  methods: {
    onPrimaryAction() {
      this.$emit("primaryAction");
    },
    onClose() {
      this.$emit("close");
    },
  },
};
</script>

<style scoped>
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



.modal-buttons button {
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
</style>
