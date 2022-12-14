<template>
  <div>
    <div class="header-container">
      <div class="img-container">
        <img
          src="/images/logo_header.png"
          alt="clean your twitter logo"
          class="icon-app"
        />
        <NuxtLink to="/">Clean your twitter</NuxtLink>
      </div>
      <div class="menu">
        <NuxtLink to="/">Clean your twitter</NuxtLink>
        <NuxtLink to="about">About</NuxtLink>
      </div>
      <div class="connect">
        <Button class="btn" text="Connect" :action="login" :fill="false" />
      </div>
    </div>
    <hr />
  </div>
</template>

<script setup lang="ts">
import Button from "@/Components/button.vue";

const login = (): void => {
  this.$router.push({ path: "/login" });
};
</script>

<style lang="scss" scoped>
@import "@/assets/scss/colors.scss";

$size-width: 90%;
$size-edge: 40%;
$size-center: 20%;

.header-container {
  display: flex;
  align-items: center;

  margin: auto;
  padding: 1em;

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
  }
}

hr {
  margin: auto;
  height: 1px;

  border: 0px;
  border-radius: 50%;

  width: $size-width;
  background-color: $main;
}
</style>
