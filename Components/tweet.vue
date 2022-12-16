<template>
  <div class="tweet">
    .header+.content+.foote
    {{ name }}
  </div>
</template>

<script setup lang="ts">
const props = defineProps({
  name: {
    type: String,
    required: true,
  },
  pseudo: {
    type: String,
    required: true,
  },
  text: {
    type: String,
    required: true,
  },
  date: {
    type: String,
    required: true,
  },
  avatar: {
    type: String,
    required: true,
  },
  remove: {
    type: Boolean,
    required: true,
  },
});
</script>

<style lang="scss" scoped></style>
