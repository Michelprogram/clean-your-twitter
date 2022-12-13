<template>
  <div>
    <p>Layout</p>
    <slot />
  </div>
</template>

<script setup></script>

<style lang="scss" scoped>
@import "@/assets/colors";
p {
  color: $red;
}
</style>
