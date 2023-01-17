<template>
  <div class="container-loader">
    <div class="custom-loader"></div>
    <p class="information">
      Tweets between {{ displayDate(twitter.getFrom) }} and
      {{ displayDate(twitter.getTo) }}
    </p>
    <p>{{ message }} ...</p>
  </div>
</template>

<script setup lang="ts">
import { useTweetStore } from "@/store/tweets";
import Formatter from "@/utils/dates";

type LoaderProps = {
  message: string;
};
defineProps<LoaderProps>();

const twitter = useTweetStore();

const displayDate = (date: string) => {
  return Formatter.short.format(new Date(date));
};
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

  .information {
    font-size: 1vw;
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
