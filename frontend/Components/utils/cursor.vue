<template>
  <div class="cursor circle" ref="cursor"></div>
</template>

<script setup lang="ts">
const position = { x: 0, y: 0 };

type Shape = "circle" | "line";

onMounted(() => {
  const cursor = document.querySelector(".cursor") as HTMLElement;

  const setCursorCSS = () => {
    cursor.style.top = `${position.y}px`;
    cursor.style.left = `${position.x}px`;
  };

  window.addEventListener("mousemove", (e: MouseEvent) => {
    position.x = e.pageX;
    position.y = e.pageY;

    requestAnimationFrame(setCursorCSS);
  });
});
</script>

<style lang="scss" scoped>
@import "@/assets/scss/colors.scss";

$size: 30px;

.cursor {
  position: absolute;
  content: "";
  width: $size;
  pointer-events: none;
  transform: translate(-50%, -50%);
  z-index: 10000;

  &.circle {
    height: $size;
    border-radius: 50px;
    border: 2px solid $main;
  }

  &.line {
    height: 2px;
    background-color: $main;
  }
}
</style>
