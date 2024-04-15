import type { Tweet, TweetClean } from "@/types/api";
import type { User } from "@/types/store";

export default class BackendApi {
  static API_ENDPOINT = () => useRuntimeConfig().public.API_ENDPOINT;

  static infoUser = async (): Promise<User> => {
    const url = this.API_ENDPOINT + "/backend/auth";

    const request = await fetch(url, {
      credentials: "include",
    });

    const json = await request.json();

    const res: User = {
      picture: json.profile_image_url,
      username: json.username,
      pseudo: json.name,
      created_at: json.created_at,
    };

    return res;
  };

  static tweets = async (
    dateStart: string,
    dateEnd: string
  ): Promise<Array<Tweet>> => {
    const url = this.API_ENDPOINT + "/backend/tweets";

    const [start, end] = [new Date(dateStart), new Date(dateEnd)];

    const request = await fetch(url, {
      credentials: "include",
      method: "POST",
      body: JSON.stringify({ date_start: start, date_end: end }),
    });

    const json = await request.json();

    return json.map((el: Tweet) => {
      el.deleted = true;
      return el;
    });
  };

  static clean = async (tweetsID: Array<string>): Promise<TweetClean> => {
    const url = this.API_ENDPOINT + "/backend/clean";

    const request = await fetch(url, {
      credentials: "include",
      method: "POST",
      body: JSON.stringify({ tweets_id: tweetsID }),
    });

    const json = await request.json();

    return json;
  };
}
