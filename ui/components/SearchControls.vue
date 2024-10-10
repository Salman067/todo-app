<template>
  <div class="search-group">
    <input
      :value="searchTask"
      placeholder="Search tasks"
      @input="$emit('updateSearchTask', $event.target.value)"
    />
    <select :value="searchStatus" @change="$emit('updateSearchStatus', $event.target.value)">
      <option value="">All Statuses</option>
      <option value="created">Created</option>
      <option value="in-progress">In Progress</option>
      <option value="done">Done</option>
    </select>
  </div>
</template>

<script>
export default {
  props: {
    searchTask: String,
    searchStatus: String
  }
}
</script>
