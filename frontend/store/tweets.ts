import BackendApi from "@/api/backend";
import { Tweet } from "@/types/api";
import { CookieRef } from "nuxt/dist/app/composables";
import { defineStore } from "pinia";
import WebSocketB from "~~/api/websocket";

export const useTweetStore = defineStore({
  id: "tweet-store",
  state: () => {
    return {
      tweets: new Array<Tweet>(),
      from: String(""),
      to: String(""),
      state: Number(0),
      limit: Number(0),
      rate: Number(0),
    };
  },
  actions: {
    UpState() {
      this.state++;
    },
    setFrom(from: string) {
      this.from = from;
    },
    setTo(to: string) {
      this.to = to;
    },
    getTweets(cookie: CookieRef<string>, start: string, end: string) {
      this.UpState();

      const socket = WebSocketB.tweets(cookie, start, end);

      socket.onmessage = (event: MessageEvent) => {
        const data = JSON.parse(event.data);
        const tweets = data.data as Array<Tweet>;
        const rate = data.rate;

        this.limit = rate.limit
        this.rate = rate.remaining

        this.setFrom(tweets[0].created_at);
        this.setTo(tweets.at(-1)!.created_at);

        tweets.forEach((t) => (t.deleted = true));

        this.tweets.push(...tweets);
      };

      socket.onclose = () => this.UpState();
    },
  },
  getters: {
    getFrom(): string {
      return this.from;
    },
    getTo(): string {
      return this.to;
    },
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
