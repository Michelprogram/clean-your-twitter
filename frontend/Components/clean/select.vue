<template>
  <div>
    <p class="title">Select</p>
    <div class="select-container">
      <div class="inputs">
        <Input title="From" v-model="startDate" />
        <Input title="To" v-model="endDate" />
      </div>
      <p class="informations">Format required : YYYY/MM/DD</p>
      <p class="informations">Your account was created {{ createdAt }}</p>
      <p class="informations error">{{ error }}</p>
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
import { useUserStore } from "@/store/user";
import Formatter from "@/utils/dates";
import Input from "@/Components/utils/input.vue";
import Button from "@/Components/utils/button.vue";

const startDate = ref("2021/01/01");
const endDate = ref("2023/01/01");
const error = ref("");

const twitter = useTweetStore();
const user = useUserStore();

const cookie = useCookie("token-twitter");

const now = new Date(Date.now());

const regex = /\d{4}\/\d{2}\/\d{2}/gm;

const findTweets = computed(async () => {
  const [start, end] = [startDate.value, endDate.value];
  if (!start.match(regex) || !end.match(regex)) {
    error.value = "Date format invalid";
    return;
  }
  if (new Date(start) < new Date(user.created_at)) {
    error.value = "Date from should be before your account creation date";
    return;
  }
  if (new Date(end) > now) {
    error.value = "You can't see in the futur";
    return;
  }
  twitter.getTweets(cookie, start, end);
});

const createdAt = computed(() => {
  const date = new Date(user.created_at || now);
  return Formatter.short.format(date);
});
</script>

<style lang="scss" scoped>
.select-container {
  margin: 1em;

  .informations {
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
