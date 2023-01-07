import { OAuthParameters } from "../types/api";

const REDIRECT_URI =
  process.env.NODE_ENV == "production"
    ? "http://www.clean-your-tw.online:3021/twitter/auth"
    : "http://www.localhost:3021/twitter/auth";

const CLIENT_ID = "UGVlUFB0N0pxeXA2UWhxX0tiZlI6MTpjaQ";

export function generateTwitterOAuth(): string {
  const options: OAuthParameters = {
    redirect_uri: REDIRECT_URI,
    client_id: CLIENT_ID,
    state: "state_against_crsf_attack",
    response_type: "code",
    code_challenge: "y_SfRG4BmOES02uqWeIkIgLQAlTBggyf_G7uKT51ku8",
    code_challenge_method: "S256",
    scope: "tweet.read tweet.write users.read offline.access",
  };
  const qs = new URLSearchParams(options).toString();
  return `https://twitter.com/i/oauth2/authorize?${qs}`;
}
//
//
