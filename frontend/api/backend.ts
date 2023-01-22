import { Tweet, TweetClean } from "@/types/api";
import { User } from "@/types/store";

let apiEndpoint =
  process.env.NODE_ENV == "production"
    ? "http://clean-your-tw.online:3021"
    : "http://localhost:3021";

export default class BackendApi {
  static infoUser = async (): Promise<User> => {
    const url = apiEndpoint + "/backend/auth";

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
    const url = apiEndpoint + "/backend/tweets";

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
    const url = apiEndpoint + "/backend/clean";

    const request = await fetch(url, {
      credentials: "include",
      method: "POST",
      body: JSON.stringify({ tweets_id: tweetsID }),
    });

    const json = await request.json();

    return json;
  };
}
