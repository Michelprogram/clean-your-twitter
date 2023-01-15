<template>
  <div>
    <Cursor />
    <Header />
    <slot />
    <Footer />
  </div>
</template>

<script setup>
import Header from "@/Components/header.vue";
import Footer from "@/Components/footer.vue";
import Cursor from "@/Components/cursor.vue";
</script>
