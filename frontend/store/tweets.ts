import BackendApi from "@/api/backend";
import { Tweet } from "@/types/api";
import { defineStore } from "pinia";

export const useTweetStore = defineStore({
  id: "tweet-store",
  state: () => {
    return {
      tweets: new Array<Tweet>(),
      state: Number(0),
    };
  },
  actions: {
    UpState() {
      this.state++;
    },
    async getTweets(startDate: string, endDate: string) {
      this.tweets = await BackendApi.tweets(startDate, endDate);
    },
  },
  getters: {
    size(): number {
      return this.tweets.length;
    },
    deleted(): number {
      return this.tweets.filter((t) => t.deleted).length;
    },
    idsDeleted(): Array<string> {
      return this.tweets.filter((t) => t.deleted).map((t) => t.id);
    },
    isWaiting(): boolean {
      return this.state == 0;
    },
    isFetching(): boolean {
      return this.state == 1;
    },
    isLookingTweets(): boolean {
      return this.state == 2;
    },
    isDeleting(): boolean {
      return this.state == 3;
    },
    tweetByWord() {
      return (word: string) => {
        return this.tweets.filter((tweet: Tweet) =>
          tweet.text.toLowerCase().includes(word.toLowerCase())
        );
      };
    },
  },
});
