<template>
  <div class="removed">
    <div v-if="twitter.stored == 0">
      <p>All your tweets have been deleted</p>
    </div>
    <div v-else>
      <p>
        <span class="number-info">{{ twitter.removed }}</span> tweets have been
        deleted and <span class="number-info">{{ twitter.stored }}</span> tweets
        have been stored.
      </p>
      <p>
        This will take
        <span class="number-info">{{ twitter.timeDelete }}</span> to remove
        <span class="number-info">{{ twitter.stored }}</span> tweets
      </p>
    </div>

    <Button text="Back home" :fill="true" :action="backHome" />
  </div>
</template>

<script setup lang="ts">
import { useTweetStore } from "@/store/tweets";
import Button from "@/Components/utils/button.vue";

const twitter = useTweetStore();
const router = useRouter();

const backHome = () => {
  twitter.initAttributes();
  router.push("/");
};
</script>

<style lang="scss" scoped>
.removed {
  height: 70%;

  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  gap: 1em;

  p {
    font-size: 1.5em;
  }

  .number-info {
    font-weight: 700;
  }
}
</style>
