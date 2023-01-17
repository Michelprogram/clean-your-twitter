<template>
  <div>
    <h1 class="title">How does it’s works ?</h1>
    <div>
      <Paragraphe
        class="paragraphe"
        v-for="paragraphe in data"
        :title="paragraphe.title"
        :description="paragraphe.descritpion"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import Paragraphe from "@/Components/about/paragraphe.vue";

const data = [
  {
    title: "Did i just hack your account ?",
    descritpion: `Unfornutality i’m not a hacker i used OAuth 2.0 which is an authentication protocol that allows an application to access a user's information on another website, while keeping the user's identification information (such as their username and password) private.\n 
      This means that you can use your account information on one website to log in to other apps, without having to share your actual identification information with those apps.\n
      When a user wants to log in to the application with their identification information, the application redirects the user to the identity provider so they can enter their identification information. Once the user enters their identification information and grants permission, the identity provider sends a token to the application, which then uses it to access the user's information.\n
      By using OAuth 2.0, apps can obtain access to user information securely, without having to ask the user to enter their identification information each time. It also makes the user experience smoother, as users don't have to remember a lot of usernames and passwords for different apps.
    `,
  },
  {
    title: "How your tweets are deleted ?",
    descritpion: `Once you connect to the application, I communicate with the Twitter API using your access tokens. Unfortunately, for security reasons, Twitter limits the number of API calls per user and per type of request.\n
       A user can make up to 900 requests to the Twitter API every 15 minutes, and can only delete 50 tweets every 15 minutes. That's why the application saves the ID of the next tweets to delete.\n
       Then a program regularly checks the stored tweets and delete them every 15 minutes. On the clean page, you can find a counter that shows you how many requests are available, and how much application needs to delete all your tweets.
       `,
  },
  {
    title: "Informations",
    descritpion: `You can find more informations about oauth2 here <a style="text-decoration:underline;" target="_blank" href="https://oauth.net/2/">OAuth2</a>, and how it's work with twitter here <a style="text-decoration:underline;" target="_blank" href="https://developer.twitter.com/en/docs/authentication/oauth-2-0/user-access-token">Twitter</a>, if you love schema you can find here at the bottom of the page <a style="text-decoration:underline;" target="_blank" href="https://developer.twitter.com/en/docs/authentication/oauth-2-0/authorization-code">Schema</a> have fun 🧙‍♂️\n
    `,
  },
];
</script>

<style lang="scss" scoped>
@import "@/assets/scss/colors.scss";

.title {
  font-size: 6vw;
  font-family: "Faustina", serif;
  font-weight: 700;
  font-style: normal;
  letter-spacing: 0.02em;
  text-align: center;

  color: $main;
}

.paragraphe {
  width: 80%;
  margin-left: 10%;
}
</style>
