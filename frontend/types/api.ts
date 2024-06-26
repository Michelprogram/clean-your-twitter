export type OAuthParameters = {
  redirect_uri: string;
  client_id: string;
  state: string;
  response_type: string;
  code_challenge: string;
  code_challenge_method: string;
  scope: string;
};

export type Tweet = {
  text: string;
  created_at: string;
  id: string;
  deleted: boolean;
};

export interface TweetClean {
  tweet_removed: number;
  tweet_stored: number;
}
