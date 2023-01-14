<template>
  <div class="clean">
    <p>{{ twitter.deleted }} tweets will be deleted on {{ twitter.size }}</p>
    <p>That represent : {{ pollution }} gram{{ plurals }} of CO2</p>
    <Button text="Clean" :action="() => {}" :fill="true" />
  </div>
</template>

<script setup lang="ts">
import { useTweetStore } from "@/store/tweets";

const twitter = useTweetStore();

const pollution = computed((): number => {
  return twitter.tweets.length * 0.02;
});

const plurals = computed((): string => {
  return pollution.value > 1 ? "s" : "";
});
</script>

<style lang="scss" scoped>
.clean > div {
  font-size: 1.5em;
  width: 10%;
}

p {
  font-size: 1.2em;
}
</style>
