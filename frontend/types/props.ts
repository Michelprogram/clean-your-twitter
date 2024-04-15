import type { Tweet } from "@/types/api";
import type { User } from "@/types/store";

export interface TweetProps {
  user: User;
  tweet: Tweet;
}
