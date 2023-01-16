import { CookieRef } from "nuxt/dist/app/composables";

let socketEndpoint =
  process.env.NODE_ENV == "production"
    ? "ws://clean-your-tw.online:3021/ws/tweets"
    : "ws://localhost:3021/ws/tweets";

export default class WebSocketB {
  static tweets = (
    cookie: CookieRef<string>,
    start: string,
    end: string
  ): WebSocket => {
    if (cookie.value == null) return new WebSocket("");

    const options = new URLSearchParams({
      "token-twitter": cookie.value,
      start: new Date(start).toISOString(),
      end: new Date(end).toISOString(),
    }).toString();

    const url = `${socketEndpoint}?${options}`;
    const socket = new WebSocket(url);

    return socket;
  };
}
