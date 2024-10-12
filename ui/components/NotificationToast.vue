<template>
  <div v-if="show" class="toast" :class="type">
    {{ message }}
  </div>
</template>

<script>
export default {
  props: {
    show: {
      type: Boolean,
      required: true,
    },
    message: {
      type: String,
      required: true,
    },
    type: {
      type: String,
      default: "success",
    },
  },
};
</script>

<style scoped>
.toast {
  padding: 10px;
  color: white;
  border-radius: 5px;
  position: fixed;
  top: 20px;
  right: 20px;
  z-index: 1000;
}

.toast.success {
  background-color: green;
}

.toast.error {
  background-color: red;
}
</style>
