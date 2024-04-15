<template>
  <div>
    <div class="title">
      <p>Preview</p>
      <p class="info-rate-limit">
        {{ twitter.rate }} on {{ twitter.limit }} requests limit, fixed by
        twitter.
      </p>
      <hr />
    </div>

    <div class="content">
      <transition name="fade" mode="out-in">
        <div v-if="twitter.isWaiting" class="waiting">
          <p>Waiting to choose dates ...</p>
          <img src="/images/svg/waiting.svg" alt="" />
        </div>
        <Loader v-else-if="twitter.isFetching" message="Data fetching" />
        <div v-else-if="!haveTweets" class="no-tweets">
          <p>
            You have no tweets between {{ twitter.from }} and {{ twitter.to }}.
          </p>
          <p>Pick others days</p>
        </div>
        <div v-else-if="twitter.isLookingTweets" class="dashboard-preview">
          <div class="filter">
            <p class="sub-title">Filters</p>
            <div class="inputs">
              <Input title="Words" v-model="filter" />
            </div>
          </div>
          <div class="tweets">
            <Tweet
              v-for="(tweet, i) in filterTweet"
              :key="i"
              :tweet="tweet"
              :user="user"
              :deleted="true"
            />
          </div>
          <div class="info">
            <p>
              By default all tweets with green bird will be remove, you can
              select which tweets to keep by clicking on birds.
            </p>
          </div>
        </div>
        <Loader v-else-if="twitter.isDeleting" message="Tweets deleting" />
        <Removed v-else />
      </transition>
    </div>
  </div>
</template>

<script setup lang="ts">
import Loader from "@/components/utils/loader.vue";
import Input from "@/components/utils/input.vue";
import Tweet from "@/components/tweet.vue";
import Removed from "@/components/clean/removed.vue";
import { useTweetStore } from "@/store/tweets";
import { useUserStore } from "@/store/user";

const twitter = useTweetStore();
const user = useUserStore();

const filter = ref("");

const filterTweet = computed(() => twitter.tweetByWord(filter.value));

const haveTweets = computed(() => twitter.size > 1);
</script>

<style lang="scss" scoped>
@import "@/assets/scss/colors";

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter,
.fade-leave-to {
  opacity: 0;
}

hr {
  height: 2px;
  color: $main;
  border: 1px solid $main;
}

.title {
  position: relative;
  .info-rate-limit {
    position: absolute;
    right: 0;
    top: 1em;
    font-size: 1vw;
    margin-top: 1em;
    font-style: italic;
    font-weight: 500;
  }
}

.content {
  height: 100%;
  .waiting {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 90%;
    margin: 2em;

    p {
      font-size: 3em;
      font-family: "Faustina", serif;
      font-weight: 700;
      font-style: normal;
      letter-spacing: 0.02em;
    }

    img {
      width: 50%;
    }
  }

  .no-tweets {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;

    gap: 1.5em;

    margin: auto;

    height: 80%;

    p:nth-child(1) {
      font-size: 2em;
    }

    p:nth-child(2) {
      font-size: 1.5em;
    }
  }

  .dashboard-preview {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    grid-template-rows: repeat(2, 1fr) 0.4fr;
    row-gap: 2em;
    margin: 1em;

    .filter {
      grid-area: 1 / 1 / 3 / 2;
      display: flex;
      flex-direction: column;
      gap: 1em;

      margin: 0.5em;

      .sub-title {
        margin: 0;
      }
    }

    .tweets {
      grid-area: 1 / 2 / 3 / 5;
      height: 350px;
      display: flex;
      flex-direction: column;
      align-items: center;
      gap: 1em;
      overflow-y: scroll;
    }

    .info {
      grid-area: 3 / 1 / 4 / 5;

      p:not(span) {
        font-size: 1em;
        font-weight: 600;
      }
    }
  }
}
</style>
