<template>
  <div class="container">
    <div class="dashboard">
      <div class="select-date">
        <p class="title">Select</p>
        <div class="inputs">
          <Input title="From" data="" />
          <Input title="To" data="" />
        </div>
        <p class="title">Suggestions</p>
        <div class="info">
          <p class="suggestion">
            Most of your tweets are behind <span id="start"></span> and
            <span id="end"></span>
          </p>
        </div>
        <Button
          text="Find tweets"
          :action="() => {}"
          :fill="true"
          class="find-tweets"
        />
      </div>
      <div class="preview">
        <div class="title">
          <p>Preview</p>
          <hr />
        </div>

        <div class="content">
          <div class="waiting">
            <p>Waiting to choose dates ...</p>
            <img v-if="!finded" src="/images/svg/waiting.svg" alt="" />
          </div>
        </div>
      </div>
      <div class="clean-info"></div>
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
import Input from "@/Components/input.vue";
import Button from "@/Components/button.vue";

const finded = false;
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

  .title {
    font-size: 3em;
    font-weight: 600;
    margin: 0.5em;
  }
}

.select-date {
  grid-area: 1 / 1 / 3 / 2;
  height: 70vh;
  display: flex;
  flex-direction: column;
  align-items: baseline;
  row-gap: 1em;

  .title:nth-child(3) {
    font-size: 1.5em;
    font-weight: 500;
  }

  .inputs {
    width: 80%;

    div:first-child {
      margin-bottom: 1em;
    }
  }

  .info {
    font-size: 1.2em;
    width: 90%;
  }

  .find-tweets {
  }
}
.preview {
  grid-area: 1 / 2 / 3 / 5;

  hr {
    height: 2px;
    color: $main;
    border: 1px solid $main;
  }

  .content {
    .waiting {
      display: flex;
      justify-content: center;
      align-items: center;
      width: 90%;
      margin: 2em;

      p {
        font-size: 3em;
        font-family: "Faustina", serif;
        font-weight: 700;
        font-style: normal;
        letter-spacing: 0.02em;
      }

      img {
        width: 50%;
      }
    }
  }
}
.clean-info {
  height: 10vh;
  grid-area: 3 / 1 / 4 / 5;
}
</style>
