const loginForm = document.getElementById("login-form");
const loginButton = document.getElementById("login-form-submit");
const loginErrorMsg = document.getElementById("login-error-msg");
const loginUrl = `/auth/login`;

loginButton.addEventListener("click", (e) => {
  e.preventDefault();
  const id = loginForm.id.value;
  const password = loginForm.password.value;
  if (id && password) {
    login_api(id, password);
  } else {
    loginErrorMsg.style.opacity = 1;
  }
});

function login_api(id, password) {
  axios
    .post(loginUrl, {
      id: id,
      password: password,
    })
    .then((res) => {
      console.log("login api success");
      console.log("data=", res.data);
      console.log("status=", res.status);
      if (res.status == HTTP_STATUS_CODE.ok) {
        console.log("login success!!!");
        // todo: get domain -> to domain url
        // window.location.replace("http://<?php echo $Domain ?>");
      }
    })
    .catch((err) => {
      console.error("login error:", err);
      loginErrorMsg.style.opacity = 1;
      switch (err.response.status) {
        case HTTP_STATUS_CODE.notFound:
          console.error("login error: path not found");
          break;
        case HTTP_STATUS_CODE.badRequest:
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
