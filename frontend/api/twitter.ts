import { OAuthParameters } from "../types/api";

export function generateTwitterOAuth(): string {
  const runtimeConfig = useRuntimeConfig();

  const [REDIRECT_URI, CLIENT_ID, TWITTER_CODE_CHALLENGE] = [
    runtimeConfig.public.API_ENDPOINT,
    runtimeConfig.public.TWITTER_CLIENT_ID,
    runtimeConfig.public.TWITTER_CODE_CHALLENGE,
  ];

  const options: OAuthParameters = {
    redirect_uri: REDIRECT_URI,
    client_id: CLIENT_ID,
    state: "state_against_crsf_attack",
    response_type: "code",
    code_challenge: TWITTER_CODE_CHALLENGE,
    code_challenge_method: "S256",
    scope: "tweet.read tweet.write users.read dm.write offline.access",
  };
  const qs = new URLSearchParams(options).toString();
  return `https://twitter.com/i/oauth2/authorize?${qs}`;
}
//
//
