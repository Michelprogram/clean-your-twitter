<template>
  <div>
    <p class="footer-description">
      Website created by
      <a href="https://fr.linkedin.com/in/dorian-gauron-534716187"
        >Gauron dorian</a
      >
      | Version :
      <span id="version"> {{ version }}</span>
    </p>
  </div>
</template>

<script setup lang="ts">
const runtimeConfig = useRuntimeConfig();

const version = runtimeConfig.public.version;
</script>

<style lang="scss" scoped>
@import "@/assets/scss/colors.scss";

div {
  padding: 1.5em;
}

.footer-description {
  text-align: center;
  color: $main;
}
</style>
