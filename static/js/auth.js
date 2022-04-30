function login() {
  window.open("/login", "_blank");
}
function is_logged_in() {
  let url = `/api/v1/test`;
  axios
    .get(url, { headers: { Authorization: getAuthorization() } })
    .then((res) => {
      if (res.status == HttpStatusCode.OK) {
        console.warn("is logged in");
      }
    })
    .catch((err) => {
      console.error(err);
    });
}
function logout() {
  let url = "/api/v1/auth/logout";
  axios
    .post(url, {}, axiosConfig)
    .then((res) => {
      console.log("logout status code=", res.status);
      if (res.status == HttpStatusCode.NO_CONTENT) {
        location.reload();
      }
    })
    .catch((err) => {
      console.error(err);
    });
}
