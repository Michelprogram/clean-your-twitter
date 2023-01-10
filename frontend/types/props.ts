import { Tweet } from "@/types/api";
import { User } from "@/types/store";

export interface TweetProps {
  user: User;
  tweet: Tweet;
}
