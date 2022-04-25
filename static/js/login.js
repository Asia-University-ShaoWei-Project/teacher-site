const loginForm = document.getElementById("login-form");
const loginButton = document.getElementById("login-form-submit");
const loginErrorMsg = document.getElementById("login-error-msg");
const loginApiUrl = `/api/v1/auth/login`;

loginButton.addEventListener("click", (e) => {
  e.preventDefault();
  let id = loginForm.id.value;
  let password = loginForm.password.value;
  if (id && password) {
    loginApi(id, password);
  } else {
    loginErrorMsg.style.opacity = 1;
  }
});

function loginApi(id, password) {
  axios
    .post(loginApiUrl, {
      id: id,
      password: password,
    })
    .then((res) => {})
    .catch((err) => {
      switch (err.response.status) {
        case HTTP_STATUS_CODE.found:
          window.location.replace(err.response.data.domain);
          break;
        case HTTP_STATUS_CODE.badRequest:
          console.error("login error: input error");
          loginErrorMsg.style.opacity = 1;
          break;
        case HTTP_STATUS_CODE.statusInternalServerError:
          alert("server has some problem");
          break;
      }
    });
}
