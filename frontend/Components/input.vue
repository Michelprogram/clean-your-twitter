<template>
  <div class="input">
    <label :for="inputName">{{ title }}</label>
    <input
      type="text"
      :name="inputName"
      :id="inputName"
      :value="modelValue"
      @input="
        $emit('update:modelValue', ($event.target as HTMLInputElement).value)
      "
    />
  </div>
</template>

<script setup lang="ts">
interface InputProps {
  title: string;
  modelValue: string;
  inputEvent?: Function;
}

const props = defineProps({
  title: String,
  modelValue: String,
});

const inputName = computed((): string => {
  return `${props.title}-name`;
});
</script>

<style lang="scss" scoped>
.input {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  width: 100%;

  label {
    font-size: 1.5em;
    margin-bottom: 0.3em;
  }

  input {
    border: none;
    padding: 0.5em;
    font-size: 1.2em;
    border-radius: 100px;
    text-decoration: none;
    outline: none;
    font-family: "TT commons";
  }
}
</style>
