import { defineStore } from "pinia";

export const useUserStore = defineStore({
  id: "user-store",
  state: () => {
    return {
      picture: String,
      username: String,
      pseudo: String,
    };
  },
  actions: {
    newStore(picture: string, username: string, pseudo: string) {
      this.picture = picture;
      this.username = username;
      this.pseudo = pseudo;
    },
  },
  getters: {},
});
