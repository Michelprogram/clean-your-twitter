import { defineStore } from "pinia";
import { User } from "@/types/store";

export const useUserStore = defineStore({
  id: "user-store",
  state: () => {
    return {
      picture: "",
      username: "",
      pseudo: "",
      created_at: "",
    };
  },
  actions: {
    newStore(user: User) {
      this.picture = user.picture;
      this.username = user.username;
      this.pseudo = user.pseudo;
      this.created_at = user.created_at;
    },
  },
  getters: {},
});
