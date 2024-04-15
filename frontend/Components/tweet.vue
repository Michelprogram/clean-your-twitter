<template>
  <div class="tweet">
    <div class="header">
      <div class="name">
        <img :src="user.picture" alt="" srcset="" class="profile-picture" />
        <div>
          <p>{{ user.username }}</p>
          <p>@{{ user.pseudo }}</p>
        </div>
      </div>
      <img
        @click="updateLink"
        class="twitter-logo"
        :src="link"
        alt=""
        srcset=""
      />
    </div>
    <div class="content">
      <p>{{ tweet.text }}</p>
    </div>
    <div class="footer">
      <p>{{ createdAt }}</p>
      <a :href="tweetLink" target="_blank">Go to the tweet</a>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Tweet } from "@/types/api";
import type { User } from "@/types/store";
import Formatter from "@/utils/dates";

interface TweetProps {
  user: User;
  tweet: Tweet;
  deleted: boolean;
}

const props = defineProps<TweetProps>();

const link = ref("/images/svg/twitter-green.svg");

const updateLink = () => {
  props.tweet.deleted = !props.tweet.deleted;
  link.value = props.tweet.deleted
    ? "/images/svg/twitter-green.svg"
    : "/images/svg/twitter-blue.svg";
};

const createdAt = computed((): string => {
  const date = new Date(props.tweet.created_at);
  return Formatter.full.format(date);
});

const tweetLink = computed((): string => {
  return `https://twitter.com/${props.user.username}/status/${props.tweet.id}`;
});
</script>

<style lang="scss" scoped>
$size-img: 55px;

.tweet {
  background-color: white;
  border-radius: 20px;
  width: 80%;
  color: black;
  .header {
    width: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1em;
    .name {
      display: flex;
      align-items: center;
      .profile-picture {
        width: $size-img;
        height: $size-img;
        border-radius: 100%;
        object-fit: cover;
      }

      div {
        margin-left: 0.8em;
        p:nth-child(1) {
          font-size: 1.5em;
          font-weight: 100;
        }
        p:nth-child(2) {
          opacity: 0.5;
        }
      }
    }
    .twitter-logo {
      cursor: pointer;
    }
  }

  .content {
    padding: 0 1em;
    font-size: 1.3em;
    font-weight: lighter;

    word-wrap: break-word;
  }

  .footer {
    padding: 1em;
    opacity: 0.5;

    display: inline-flex;
    gap: 1em;

    a {
      text-decoration: underline;
      color: #1da1f2;
    }
  }
}
</style>
