<template>
  <div class="cursor" ref="cursor" :class="className" :style="cursorCSS"></div>
</template>

<script setup lang="ts">
import { CSSProperties } from "vue";

type Position = {
  x: number;
  y: number;
};

const cursor = ref<HTMLDivElement | null>(null);
const className = ref("");

const position = ref<Position>({ x: 0, y: 0 });

const route = useRoute();

const updatePosition = (e: MouseEvent) => {
  position.value.x = e.pageX;
  position.value.y = e.pageY;
};

const cursorCSS = computed((): CSSProperties => {
  return {
    top: `${position.value.y}px`,
    left: `${position.value.x}px`,
  };
});

const eventButtons = () => {
  const buttons = document.querySelectorAll<HTMLElement>(".button-container");

  buttons.forEach((e) => {
    e.addEventListener(
      "mouseover",
      () => (className.value = "cursor over-button")
    );
    e.addEventListener("mouseleave", () => (className.value = "cursor"));
  });
};

const eventTexts = () => {
  const texts = document.querySelectorAll<HTMLElement>(".cursor-text");

  texts.forEach((e) => {
    e.addEventListener(
      "mouseover",
      () => (className.value = "cursor over-text")
    );
    e.addEventListener("mouseleave", () => (className.value = "cursor"));
  });
};

const eventTitle = () => {
  const titles = document.querySelectorAll<HTMLElement>(".cursor-title");

  titles.forEach((e) => {
    e.addEventListener(
      "mouseover",
      () => (className.value = "cursor over-title")
    );
    e.addEventListener("mouseleave", () => (className.value = "cursor"));
  });
};

onMounted(() => {
  eventTexts();
  eventButtons();
  eventTitle();
  window.addEventListener("mousemove", updatePosition);
});

watch(route, () => {
  eventTexts();
  eventButtons();
  eventTitle();
});
</script>

<style lang="scss" scoped>
@import "@/assets/scss/colors.scss";

$size: 30px;

.cursor {
  position: absolute;
  content: "";
  width: $size;
  height: $size;
  pointer-events: none;
  transform: translate(-50%, -50%);
  z-index: 10000;
  transition: background-color 0.5s ease-in, opacity 1s ease-in-out,
    border-radius 0.5s ease-in, height 0.5s ease-in, width 0.5s ease-in;
  border: 2px solid $main;
  border-radius: 50px;

  &.over {
    &-text {
      opacity: 0.1;
      background-color: $main;
    }
    &-title {
    }
    &-button {
      height: 10px;
      width: 10px;
      background-color: $background;
    }
  }
}
</style>
