import { Tweet } from "~~/types/api";

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

    const user = json.user.profile.data;

    const res: User = {
      picture: user.profile_image_url,
      username: user.username,
      pseudo: user.name,
    };

    return res;
  };

  //TODO : ADD promise type and update vue files
  static tweets = async () => {
    const url = apiEndpoint + "/backend/tweets";

    const request = await fetch(url, {
      credentials: "include",
    });

    const json = await request.json();

    return json.data;
  };
}

type User = {
  picture: string;
  username: string;
  pseudo: string;
};

//
//
