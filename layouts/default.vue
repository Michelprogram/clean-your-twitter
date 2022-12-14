<template>
  <div>
    <Header />
    <p>testtt</p>
    <slot />
    <Footer />
  </div>
</template>

<script setup>
import Header from "@/Components/header.vue";
import Footer from "@/Components/footer.vue";
</script>

<style lang="scss" scoped></style>
