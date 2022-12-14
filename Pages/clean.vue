<template>
  <div class="container">
    <div class="dashboard">
      <div class="select-date"></div>
      <div class="preview"></div>
      <div class="clean-info"><p>test</p></div>
    </div>
    <div class="shapes">
      <img src="/images/svg/shapes/fraise.svg" alt="" />
      <img src="/images/svg/shapes/patate.svg" alt="" />
      <img src="/images/svg/shapes/pill-2.svg" alt="" />
      <img src="/images/svg/shapes/rectangle.svg" alt="" />
    </div>
  </div>
</template>

<script setup></script>

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
  display: grid;
  position: relative;
  grid-template-columns: repeat(4, 1fr);
  grid-template-rows: repeat(2, 1fr) 0.4fr;
  grid-column-gap: 2em;
  grid-row-gap: 1em;

  width: 90%;
  margin-left: 5%;

  div {
    background-color: $highlighted;
    border-radius: 15px;
    color: $main;
    z-index: 1;
  }
}

.select-date {
  grid-area: 1 / 1 / 3 / 2;
  height: 70vh;
}
.preview {
  grid-area: 1 / 2 / 3 / 5;
}
.clean-info {
  height: 10vh;
  grid-area: 3 / 1 / 4 / 5;
}
</style>
