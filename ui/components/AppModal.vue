<template>
  <div
    v-if="isOpen"
    class="modal"
    role="dialog"
    aria-modal="true"
    @keydown.esc="onClose"
  >
    <div class="modal-content" tabindex="-1">
      <h2>{{ title }}</h2>
      <slot/>
      <div class="modal-buttons">
        <button class="cancel-btn" @click="onClose">
          {{ closeButtonText }}
        </button>
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
    closeButtonText: {
      type: String,
      default: "Cancel",
    },
  },
  emits: ['close'], // Declare the "close" event
  methods: {
    onClose() {
      this.$emit("close");
    },
  },
};
</script>

<style scoped>
.modal {
  position: fixed;
  z-index: 1000;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.4);
  display: flex;
  align-items: center;
  justify-content: center;
}

.modal-content {
  background-color: #fff;
  padding: 20px;
  border-radius: 8px;
  width: 100%;
  max-width: 600px;
}

.modal-buttons {
  display: flex;
  justify-content: flex-end;
}

.cancel-btn {
  width: 100%;
  margin-top: 7px;
}

button {
  padding: 8px 16px;
  border: none;
  border-radius: 4px;
  background-color: #4a90e2;
  color: white;
  cursor: pointer;
}

button:hover {
  background-color: #357abd;
}
</style>
