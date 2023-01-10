<template>
  <div>
    <Header />
    <slot />
    <Footer />
  </div>
</template>

<script setup>
import Header from "@/Components/header.vue";
import Footer from "@/Components/footer.vue";
</script>
