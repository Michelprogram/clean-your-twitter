<template>
  <div class="paragraphe-component">
    <h3 class="title cursor-title">{{ title }}</h3>
    <p class="text cursor-text" v-html="description"></p>
  </div>
</template>

<script setup lang="ts">
defineProps({
  title: { type: String },
  description: { type: String },
});
</script>

<style lang="scss" scoped>
@forward "@/assets/scss/variables/fonts";

@import "@/assets/scss/colors.scss";

.paragraphe-component {
  color: $main;

  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  gap: 1em;

  margin: 2em;
  .title {
    font-family: "Faustina", serif;
    font-size: 3vw;
    font-weight: 500;
    letter-spacing: 0.02em;
  }

  .text {
    font-size: 1.5vw;
    white-space: pre-line;

    .link {
      background-color: red;
    }
  }
}
</style>
