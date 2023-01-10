<template>
  <div class="container">
    <div class="dashboard">
      <div class="select-date">
        <p class="title">Select</p>
        <div>
          <div class="inputs">
            <Input title="From" data="" disabled="true" />
            <Input title="To" data="" disabled="true" />
          </div>
          <p class="sub-title">Suggestions</p>
          <div class="info">
            <p class="suggestion">
              Most of your tweets are behind <span id="start"></span> and
              <span id="end"></span>
            </p>
          </div>
        </div>
        <Button
          text="Find tweets"
          :action="() => {}"
          :fill="true"
          class="find-tweets"
        />
      </div>
      <div class="preview">
        <div class="title">
          <p>Preview</p>
          <hr />
        </div>

        <div class="content">
          <div v-if="!finded" class="waiting">
            <p>Waiting to choose dates ...</p>
            <img src="/images/svg/waiting.svg" alt="" />
          </div>
          <div v-else class="dashboard-preview">
            <div class="filter">
              <p class="sub-title">Filters</p>
              <div class="inputs">
                <Input title="Words" data="" :input-event="filterTweets" />
              </div>
              <Button
                text="Apply"
                :action="() => {}"
                :fill="true"
                class="apply"
              />
            </div>
            <div class="tweets">
              <Tweet
                v-for="(tweet, i) in copyTweets"
                :key="i"
                :tweet="tweet"
                :user="user"
              />
            </div>
            <div class="info">
              <p>
                <span class="sub-title">Info :</span>By default all tweets with
                green bird will be remove, you can select which tweets to keep
                by clicking on birds.
              </p>
            </div>
          </div>
        </div>
      </div>
      <div class="clean-info">
        <Button text="Clean" :action="() => {}" :fill="true" />
      </div>
    </div>
    <div class="shapes">
      <img src="/images/svg/shapes/fraise.svg" alt="" />
      <img src="/images/svg/shapes/patate.svg" alt="" />
      <img src="/images/svg/shapes/pill-2.svg" alt="" />
      <img src="/images/svg/shapes/rectangle.svg" alt="" />
    </div>
  </div>
</template>

<script setup lang="ts">
import Input from "@/Components/input.vue";
import Button from "@/Components/button.vue";
import Tweet from "@/Components/tweet.vue";
import BackendApi from "@/api/backend";
import { useUserStore } from "@/store/user";
import { Tweet as typedTweet } from "@/types/api";
import { User } from "@/types/store";

const user = useUserStore() as User;

const finded = ref(false);

const tweets = ref<Array<typedTweet>>([]);

const copyTweets = ref<Array<typedTweet>>([]);

onMounted(async () => {
  tweets.value = await BackendApi.tweets();
  copyTweets.value = tweets.value;
  finded.value = true;
});

const filterTweets = (event: Event) => {
  const target = event.target as HTMLInputElement;
  const value = target.value.toLowerCase();
  copyTweets.value = tweets.value.filter((tweet: any) =>
    tweet.text.toLowerCase().includes(value)
  );
};
</script>

<style lang="scss" scoped>
@forward "@/assets/scss/variables/fonts";
@import "@/assets/scss/colors";

.container {
  overflow: hidden;
  position: relative;
  padding-top: 3em;

  .shapes {
    img {
      position: absolute;

      &:first-child {
        right: -3em;
        bottom: 1em;
      }

      &:nth-child(2) {
        top: -2em;
        right: 0em;
      }

      &:nth-child(3) {
        top: 8em;
      }

      &:last-child {
        left: 19em;
        bottom: 2em;
      }
    }
  }
}

.dashboard {
  position: relative;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  grid-template-rows: repeat(2, 1fr) 0.4fr;
  grid-column-gap: 2em;
  grid-row-gap: 1em;

  width: 90%;
  margin-left: 5%;

  .preview,
  .select-date,
  .clean-info {
    background-color: $highlighted;
    border-radius: 15px;
    color: $main;
    z-index: 1;
  }

  .title {
    font-size: 3em;
    font-weight: 600;
    margin: 0.5em;
  }
  .sub-title {
    font-size: 2em;
    font-weight: 500;
    margin: 0.5em;
  }
}

.select-date {
  grid-area: 1 / 1 / 3 / 2;
  height: 70vh;
  display: flex;
  flex-direction: column;
  align-items: baseline;
  row-gap: 1em;

  .inputs {
    width: 80%;
    margin: 0.5em;

    div:first-child {
      margin-bottom: 1em;
    }
  }

  .info {
    font-size: 1.2em;
    width: 90%;
  }
}
.preview {
  grid-area: 1 / 2 / 3 / 5;

  hr {
    height: 2px;
    color: $main;
    border: 1px solid $main;
  }

  .content {
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
}
.clean-info {
  height: 10vh;
  grid-area: 3 / 1 / 4 / 5;

  display: flex;
  align-items: center;
  justify-content: center;

  div {
    font-size: 1.5em;
    width: 10%;
  }
}
</style>
