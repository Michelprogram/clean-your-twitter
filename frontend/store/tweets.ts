import type { CookieRef } from "#app";
import BackendApi from "@/api/backend";
import type { Tweet } from "@/types/api";
import { defineStore } from "pinia";
import WebSocketB from "~~/api/websocket";

export const useTweetStore = defineStore({
  id: "tweet-store",
  state: () => {
    return {
      tweets: new Array<Tweet>(),
      from: String("01/01/2020"),
      to: String("02/01/2020"),
      state: Number(0),
      limit: Number(0),
      rate: Number(0),
      removed: Number(0),
      stored: Number(0),
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
    initAttributes() {
      this.tweets = new Array<Tweet>();
      this.from = String("01/01/2020");
      this.to = String("02/01/2020");
      this.state = Number(0);
      this.limit = Number(0);
      this.rate = Number(0);
      this.removed = 0;
      this.stored = 0;
    },
    getTweets(cookie: CookieRef<string>, start: string, end: string) {
      this.initAttributes();

      this.UpState();

      const socket = WebSocketB.tweets(cookie, start, end);

      socket.onmessage = (event: MessageEvent) => {
        const data = JSON.parse(event.data);
        const tweets = data.data as Array<Tweet>;
        const rate = data.rate;

        this.limit = rate.limit;
        this.rate = rate.remaining;

        this.setFrom(tweets[0].created_at);
        this.setTo(tweets.at(-1)!.created_at);

        tweets.forEach((t) => (t.deleted = true));

        this.tweets.push(...tweets);
      };

      socket.onclose = () => this.UpState();
    },
    async cleanTweets(tweets: Array<string>) {
      this.UpState();
      const res = await BackendApi.clean(tweets);
      this.removed = res.tweet_removed;
      this.stored = res.tweet_stored;
      this.UpState();
    },
  },
  getters: {
    timeDelete(): string {
      const factor = (this.stored - 50) / 50;

      const minutes = Math.floor(15 * factor);

      if (minutes < 60) return `0h${minutes}`;

      const hours = minutes / 60;
      return `${hours}`;
    },
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
