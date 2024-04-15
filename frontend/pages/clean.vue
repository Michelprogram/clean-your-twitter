<template>
  <div class="container">
    <div class="dashboard">
      <Select class="select-date" />
      <Preview class="preview" />
      <Clean class="clean-info" />
    </div>
    <div class="shapes">
      <img src="/images/svg/shapes/fraise.svg" alt="" />
      <img src="/images/svg/shapes/patate.svg" alt="" />
      <img src="/images/svg/shapes/pill-2.svg" alt="" />
      <img src="/images/svg/shapes/rectangle.svg" alt="" />
    </div>
  </div>
</template>

<script setup lang="ts">
import Clean from "@/components/clean/clean.vue";
import Select from "@/components/clean/select.vue";
import Preview from "@/components/clean/preview.vue";
</script>

<style lang="scss" scoped>
@forward "@/assets/scss/variables/fonts";
@import "@/assets/scss/colors";

.container {
  overflow: hidden;
  position: relative;
  padding-top: 3em;

  .shapes {
    img {
      position: absolute;

      &:first-child {
        right: -3em;
        bottom: 1em;
      }

      &:nth-child(2) {
        top: -2em;
        right: 0em;
      }

      &:nth-child(3) {
        top: 8em;
      }

      &:last-child {
        left: 19em;
        bottom: 2em;
      }
    }
  }
}

.dashboard {
  position: relative;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  grid-template-rows: repeat(2, 1fr) 0.4fr;
  grid-column-gap: 2em;
  grid-row-gap: 1em;

  width: 90%;
  margin-left: 5%;

  .preview,
  .select-date,
  .clean-info {
    background-color: $highlighted;
    border-radius: 15px;
    color: $main;
    z-index: 1;
  }
}

.select-date {
  grid-area: 1 / 1 / 3 / 2;
  height: 70vh;
  display: flex;
  flex-direction: column;
  row-gap: 1em;
}
.preview {
  grid-area: 1 / 2 / 3 / 5;
}
.clean-info {
  height: 10vh;
  grid-area: 3 / 1 / 4 / 5;

  display: flex;
  align-items: center;
  justify-content: space-around;
}
</style>
