<template>
  <div class="home-container">
    <div class="description">
      <p class="title">When it’s time to clean your account !</p>
      <p class="presentation">
        Using the web application, you can easily clean up your Twitter account
        and do your part in the fight against climate change..
      </p>
      <div class="btns">
        <Button v-if="!isLogged" text="Connect" :action="login" :fill="true" />
        <Button text="Learn More" :action="about" :fill="false" />
      </div>
    </div>
    <div class="preview">
      <img
        src="/images/svg/illustration.svg"
        alt="illustration-cleaning"
        srcset=""
      />
      <img
        class="element-next-img img-1"
        src="/images/svg/shapes/curve.svg"
        alt="funky-icon"
        srcset=""
      />
      <img
        class="element-next-img img-2"
        src="/images/svg/shapes/pill.svg"
        alt="funky-icon"
        srcset=""
      />
      <img
        class="element-next-img img-3"
        src="/images/svg/shapes/triangle.svg"
        alt="funky-icon"
        srcset=""
      />
      <p class="element-next-img text">Cleaning ...</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import Button from "@/components/utils/button.vue";
import { generateTwitterOAuth } from "@/api/twitter";
import { useUserStore } from "@/store/user";

const router = useRouter();

const user = useUserStore();

const login = (): void => {
  location.href = generateTwitterOAuth();
};

const about = (): void => {
  router.push("/about");
};

const isLogged = computed((): boolean => {
  return user.username != "";
});
</script>

<style lang="scss" scoped>
@forward "@/assets/scss/variables/fonts";
@import "@/assets/scss/colors.scss";

.home-container {
  display: grid;
  grid-template-columns: 1fr 1fr;
  grid-template-rows: 1fr;
  grid-column-gap: 1em;
  grid-row-gap: 0px;
  padding-top: 3em;

  width: 90%;
  margin-left: 5%;

  div {
    height: 70vh;
  }
}

.description {
  grid-area: 1 / 1 / 2 / 2;
  color: $main;
  display: flex;
  flex-direction: column;
  gap: 3%;

  .title {
    font-size: 6vw;
    font-family: "Faustina", serif;
    font-weight: 700;
    font-style: normal;
    letter-spacing: 0.02em;
  }

  .presentation {
    font-size: 1.5vw;
  }

  .btns {
    display: flex;
    align-items: baseline;
    justify-content: space-between;
    width: 50%;
  }
}
.preview {
  margin: auto;
  position: relative;
  grid-area: 1 / 2 / 2 / 3;

  .element-next-img {
    position: absolute;

    &.text {
      top: -55px;
      right: 0px;

      transform: rotate(10deg);
      color: $main;

      font-family: "Faustina", serif;
      font-size: 3vw;
      letter-spacing: 0.05em;
      font-weight: 700;
    }

    &.img-1 {
      top: 5em;
      right: -5em;
    }

    &.img-2 {
      top: 1em;
    }

    &.img-3 {
      top: 10em;
      left: -10em;
    }
  }
}
</style>
