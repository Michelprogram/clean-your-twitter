<template>
  <div>
    <div :class="styleClass + ' ' + disabled" @click="action()">
      <p>{{ text }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
type ButtonProps = {
  text: string;
  action: Function;
  fill: boolean;
  disable?: boolean;
};

const props = defineProps<ButtonProps>();

const styleClass = computed((): String => {
  const defaultStyle = "button-container ";
  return props.fill ? defaultStyle + "fill" : defaultStyle + "outline";
});

const disabled = computed((): String => {
  return props.disable ? "disable" : "";
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
  }
  &.disable {
    cursor: none;
    pointer-events: none;
    filter: grayscale(70%);
  }
  &.fill {
    background-color: $main;
    color: $background;
  }

  &.outline {
    color: $main;
    border: 1.5px solid $main;
  }
}
</style>
