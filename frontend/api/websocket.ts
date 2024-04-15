import { CookieRef } from "nuxt/dist/app/composables";

export default class WebSocketB {
  static tweets = (
    cookie: CookieRef<string>,
    start: string,
    end: string
  ): WebSocket => {
    const runtimeConfig = useRuntimeConfig();

    const API_WEBSOCKET = runtimeConfig.public.API_WEBSOCKET;

    if (cookie.value == null) return new WebSocket("");

    const options = new URLSearchParams({
      "token-twitter": cookie.value,
      start: new Date(start).toISOString(),
      end: new Date(end).toISOString(),
    }).toString();

    const url = `${API_WEBSOCKET}?${options}`;
    const socket = new WebSocket(url);

    return socket;
  };
}
