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
      <p>{{ displayDate }}</p>
      <a :href="tweetLink" target="_blank">Go to the tweet</a>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Tweet } from "@/types/api";
import { User } from "@/types/store";

interface TweetProps {
  user: User;
  tweet: Tweet;
  deleted: boolean;
}

const props = defineProps<TweetProps>();

const remove = ref(false);
const link = ref("/images/svg/twitter-green.svg");

const updateLink = () => {
  remove.value = !remove.value;
  link.value = remove.value
    ? "/images/svg/twitter-blue.svg"
    : "/images/svg/twitter-green.svg";
};

const displayDate = computed((): string => {
  const options: Intl.DateTimeFormatOptions = {
    hour: "numeric",
    minute: "numeric",
    month: "long",
    year: "numeric",
    day: "numeric",
  };
  const formatter = Intl.DateTimeFormat("en-US", options);
  const date = new Date(props.tweet.created_at);

  return formatter.format(date);
});

const tweetLink = computed((): string => {
  if (
    props.tweet.entities == undefined ||
    props.tweet.entities.urls == undefined
  )
    return `https://twitter.com/${props.user.username}/status/${props.tweet.id}`;
  return props.tweet.entities.urls[0];
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
