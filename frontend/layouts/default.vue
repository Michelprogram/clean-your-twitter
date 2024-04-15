<template>
  <div>
    <Cursor />
    <Header />
    <slot />
    <Footer />
  </div>
</template>

<script setup>
import Header from "@/components/header.vue";
import Footer from "@/components/footer.vue";
import Cursor from "@/components/utils/cursor.vue";
</script>
