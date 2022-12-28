export type OAuthParameters = {
  redirect_uri: string;
  client_id: string;
  state: string;
  response_type: string;
  code_challenge: string;
  code_challenge_method: string;
  scope: string;
};