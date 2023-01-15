<template>
  <div class="container-loader">
    <div class="custom-loader"></div>
    <p>{{ message }} ...</p>
  </div>
</template>

<script setup lang="ts">
type LoaderProps = {
  message: string;
};

const props = defineProps<LoaderProps>();
</script>

<style lang="scss" scoped>
@import "@/assets/scss/colors.scss";

.container-loader {
  height: 70%;

  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  gap: 1em;

  .custom-loader {
    width: 50px;
    height: 50px;
    border-radius: 50%;
    background: $main;
    -webkit-mask: radial-gradient(
      circle closest-side at 50% 40%,
      #0000 94%,
      #000
    );
    transform-origin: 50% 40%;
    animation: s5 1s infinite linear;
  }

  p {
    font-size: 3vw;
  }

  @keyframes s5 {
    100% {
      transform: rotate(1turn);
    }
  }
}
</style>
