const loginForm = document.getElementById("login-form");
const loginButton = document.getElementById("login-form-submit");
const loginErrorMsg = document.getElementById("login-error-msg");

loginButton.addEventListener("click", (e) => {
  e.preventDefault();
  let id = loginForm.id.value;
  let password = loginForm.password.value;
  if (id && password) {
    login_api(id, password);
  } else {
    loginErrorMsg.style.opacity = 1;
  }
});

// todo: stored session(token), but can't get value for other path
function login_api(id, password) {
  let loginUrl = `/api/v1/auth/login`;

  axios
    .post(loginUrl, {
      id: id,
      password: password,
    })
    .then((res) => {})
    .catch((err) => {
      console.error("login error:", err);
      switch (err.response.status) {
        case HTTP_STATUS_CODE.found:
          console.log(err.response);
          window.location.replace(err.response.data.domain);
          break;
        case HTTP_STATUS_CODE.badRequest:
          loginErrorMsg.style.opacity = 1;
          console.error("login error: input error");
          break;
        case HTTP_STATUS_CODE.statusInternalServerError:
          alert("server has some problem");
          break;
        default:
          alert("unknown code");
          break;
      }
    });
}

function is_logged_in() {
  axios
    .get("/api/v1/test", { headers: { Authorization: getAuthorization() } })
    .then((res) => {
      console.log("login api success");
      console.log("data=", res.data);
      console.log("status=", res.status);
      if (res.status == HTTP_STATUS_CODE.ok) {
        console.log("login success!!!");
      }
    })
    .catch((err) => {
      console.error(err);
    });
}
function logout() {
  axios
    .post(
      "/api/v1/auth/logout",
      {},
      { headers: { Authorization: getAuthorization() } }
    )

    .then((res) => {
      console.log("logout status code=", res.status);
      if (res.status == HTTP_STATUS_CODE.noContent) {
        console.log("login success!!!");
        sessionStorage.removeItem("Authorization");
        // todo: get domain -> to domain url
        // window.location.replace("http://<?php echo $Domain ?>");
      }
    })
    .catch((err) => {
      console.error(err);
    });
}
function getAuthorization() {
  return sessionStorage.getItem("Authorization");
}
