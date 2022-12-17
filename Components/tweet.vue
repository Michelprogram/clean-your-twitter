<template>
  <div class="tweet">
    <div class="header">
      <div class="name">
        <img :src="avatar" alt="" srcset="" class="profile-picture" />
        <div>
          <p>{{ name }}</p>
          <p>@{{ pseudo }}</p>
        </div>
      </div>
      <img
        @click="updateLogo()"
        class="twitter-logo"
        :src="logo"
        alt=""
        srcset=""
      />
    </div>
    <div class="content">
      <p>{{ text }}</p>
    </div>
    <div class="footer">
      <p>{{ date }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps({
  name: {
    type: String,
    required: true,
  },
  pseudo: {
    type: String,
    required: true,
  },
  text: {
    type: String,
    required: true,
  },
  date: {
    type: String,
    required: true,
  },
  avatar: {
    type: String,
    required: true,
  },
});

const remove = ref(true);

const logo = computed(() => {
  const path = "/images/svg/";

  return remove ? path + "twitter-blue.svg" : path + "twitter-green.svg";
});

const updateLogo = () => {
  remove.value = !remove.value;
};
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
  }
}
</style>
