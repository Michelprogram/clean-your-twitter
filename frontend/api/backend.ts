const authBackend = async () => {
  const url = "http://localhost:3021/auth/backend";

  const request = await fetch(url, {
    credentials: "include",
  });

  return request;
};
