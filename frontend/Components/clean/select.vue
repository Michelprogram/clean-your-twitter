<template>
  <div>
    <p class="title">Select</p>
    <div class="select-container">
      <div class="inputs">
        <Input title="From" v-model="startDate" />
        <Input title="To" v-model="endDate" />
      </div>
      <p class="format-date">Format : YYYY/MM/DD</p>
    </div>
    <Button
      text="Find tweets"
      :action="() => findTweets"
      :fill="true"
      class="find-tweets"
    />
  </div>
</template>

<script setup lang="ts">
import { useTweetStore } from "@/store/tweets";

const startDate = ref("2022/12/21");
const endDate = ref("2023/01/01");

const twitter = useTweetStore();

const regex = /\d{4}\/\d{2}\/\d{2}/gm;

const findTweets = computed(async () => {
  const [start, end] = [startDate.value, endDate.value];
  if (start.match(regex) || end.match(regex)) {
    twitter.UpState();
    await twitter.getTweets(start, end);
    twitter.UpState();
  }
});
</script>

<style lang="scss" scoped>
.select-container {
  margin: 1em;

  .format-date {
    font-size: 1vw;
    margin-top: 1em;
    font-style: italic;
  }
}

.find-tweets {
  width: 80%;
  margin: 1em;
}

.inputs {
  width: 80%;

  div:first-child {
    margin-bottom: 1em;
  }
}
</style>
