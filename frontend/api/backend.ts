export default class BackendApi {
  static infoUser = async (): Promise<User> => {
    const url = "http://localhost:3021/backend/auth";

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

  static tweets = async () => {
    const url = "http://localhost:3021/backend/tweets";

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
