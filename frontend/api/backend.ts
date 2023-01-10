import { Tweet } from "@/types/api";
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
    };

    return res;
  };

  static tweets = async (): Promise<Array<Tweet>> => {
    const url = apiEndpoint + "/backend/tweets";

    const request = await fetch(url, {
      credentials: "include",
    });

    const json = await request.json();

    return [...json.data].map((el) => {
      el.deleted = true;
      return el;
    });
  };
}
