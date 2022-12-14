<template>
  <div>
    <div :class="styleClass" @click="action()">
      <p>{{ text }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps({
  text: { type: String, required: true },
  action: { type: Function, required: true },
  fill: { type: Boolean, required: true },
});

const styleClass = computed((): String => {
  const defaultStyle = "button-container ";
  return props.fill ? defaultStyle + "fill" : defaultStyle + "outline";
});
</script>

<style lang="scss" scoped>
@import "@/assets/scss/colors.scss";

.button-container {
  font-size: 1.3em;
  width: 100%;
  border-radius: 50px;
  cursor: pointer;
  text-align: center;
  p {
    padding: 0.3em 1em;
    color: $main;
  }
  &.fill {
    background-color: $main;
    color: $background;
  }

  &.outline {
    border: 1.5px solid $main;
  }
}
</style>
