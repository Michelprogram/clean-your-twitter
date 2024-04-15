<template>
  <div>
    <div class="header-container">
      <div class="img-container">
        <NuxtLink to="/">
          <img
            src="/images/logo_header.png"
            alt="clean your twitter logo"
            class="icon-app"
          />
        </NuxtLink>
      </div>
      <div class="menu">
        <NuxtLink to="/" class="cursor-text">Clean your twitter</NuxtLink>
        <NuxtLink to="/clean" class="cursor-text">Dashboard</NuxtLink>
        <NuxtLink to="about" class="cursor-text">About</NuxtLink>
      </div>
      <div v-if="user.isLogged" class="connect">
        <div class="images-container">
          <img
            class="picture-profile"
            :src="user.picture"
            alt="profile picture"
          />
          <img
            class="logout"
            src="/images/deconnection.png"
            @click="() => logout"
          />
        </div>
      </div>
      <div v-else class="connect">
        <Button class="btn" text="Connect" :action="login" :fill="false" />
      </div>
    </div>
    <hr />
  </div>
</template>

<script setup lang="ts">
import Button from "@/components/utils/button.vue";
import { generateTwitterOAuth } from "@/api/twitter";
import BackendApi from "@/api/backend";
import { useUserStore } from "@/store/user";

const cookie = useCookie("token-twitter");
const url = useRoute();
const user = useUserStore();

const login = (): void => {
  location.href = generateTwitterOAuth();
};

const logout = computed((): void => {
  cookie.value = "";
  const url = new URL("/", window.location.origin);
  window.location.href = url.toString();
});

onMounted(async () => {
  //Regarde si l'url contient le jwt puis set le cookie
  if (url.query.info) {
    const jwtToken = url.query.info as string;
    cookie.value = jwtToken;
  }
  //Cookie existe verifie l'identité
  if (cookie.value) {
    const u = await BackendApi.infoUser();
    user.newStore(u);
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

    .images-container {
      display: flex;
      position: relative;

      &:hover {
        .picture-profile {
          opacity: 0;
        }
        .logout {
          opacity: 1;
        }
      }

      .picture-profile,
      .logout {
        transition: opacity 0.5s ease-in-out;
      }

      .picture-profile {
        border-radius: 5px;
      }
      .logout {
        position: absolute;
        padding: 0.3em;
        opacity: 0;
        cursor: pointer;
      }
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
