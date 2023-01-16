<template>
  <div class="container-loader">
    <div class="custom-loader"></div>
    <p class="information">
      Tweet between {{ displayDateFrom(twitter.getFrom) }} to
      {{ displayDateFrom(twitter.getTo) }}
    </p>
    <p>{{ message }} ...</p>
  </div>
</template>

<script setup lang="ts">
import { useTweetStore } from "@/store/tweets";

type LoaderProps = {
  message: string;
};

const twitter = useTweetStore();

const props = defineProps<LoaderProps>();

const displayDateFrom = (date: string) => {
  const options: Intl.DateTimeFormatOptions = {
    hour: "numeric",
    minute: "numeric",
    month: "long",
    year: "numeric",
    day: "numeric",
  };
  const formatter = Intl.DateTimeFormat("en-US", options);
  return formatter.format(new Date(date));
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
