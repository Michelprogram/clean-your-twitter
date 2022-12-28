<template>
  <div>
    <div class="header-container">
      <div class="img-container">
        <img
          src="/images/logo_header.png"
          alt="clean your twitter logo"
          class="icon-app"
        />
      </div>
      <div class="menu">
        <NuxtLink to="/">Clean your twitter</NuxtLink>
        <NuxtLink v-if="isConnected" to="/clean">Dashboard</NuxtLink>
        <NuxtLink to="about">About</NuxtLink>
      </div>
      <div v-if="isConnected" class="connect">
        <img class="picture-profile" :src="profilePicture" alt="" srcset="" />
      </div>
      <div v-else class="connect">
        <Button class="btn" text="Connect" :action="login" :fill="false" />
      </div>
    </div>
    <hr />
  </div>
</template>

<script setup lang="ts">
import Button from "@/Components/button.vue";
import { generateTwitterOAuth } from "../api/twitter";

const cookie = useCookie("token-twitter");
const url = useRoute();

const profilePicture = ref("");

const login = (): void => {
  location.href = generateTwitterOAuth();
};

const authBackend = async () => {
  const url = "http://localhost:3021/auth/backend";

  const request = await fetch(url, {
    credentials: "include",
  });

  const data = await request.json();

  return data.user.profile.data.profile_image_url;
};

const isConnected = computed((): boolean => {
  return profilePicture.value != "";
});

onMounted(async () => {
  //Regarde si l'url contient le jwt puis set le cookie
  if (url.query.info) {
    const jwtToken = url.query.info as string;
    cookie.value = jwtToken;
  }
  //Cookie existe verifie l'identité
  if (cookie.value) {
    profilePicture.value = await authBackend();
  }
});
</script>

<style lang="scss" scoped>
@import "@/assets/scss/colors.scss";

$size-width: 90%;
$size-edge: 40%;
$size-center: 40%;

.header-container {
  display: flex;
  align-items: center;

  margin: auto;
  padding: 1.5em 1em;

  width: $size-width;
  color: $main;

  .img-container {
    display: flex;
    gap: 1em;
    align-items: center;

    font-weight: 400;
    font-size: 1.5em;

    width: $size-edge;

    img {
      width: 50px;
      height: 50px;
    }
  }

  .menu {
    display: flex;
    justify-content: space-around;

    width: $size-center;
    font-size: 1.3em;
    cursor: pointer;

    a {
      position: relative;
      &::after {
        position: absolute;
        content: "";

        width: 0%;
        margin-left: 5%;
        height: 2px;

        left: 0;
        bottom: -5px;

        background-color: $main;

        transition: width 0.5s ease-in-out;
      }
    }

    .router-link-active {
      &::after {
        width: 90%;
      }
    }
  }

  .connect {
    width: $size-edge;
    display: flex;
    flex-direction: row-reverse;

    .btn {
      width: fit-content;
    }

    .picture-profile {
      border-radius: 5px;
    }
  }
}

hr {
  margin-left: 5%;
  margin-bottom: 2em;
  height: 1px;

  border: 0px;
  border-radius: 50%;

  width: $size-width;
  background-color: $main;
}
</style>
