const formElem = document.getElementById("login-form");
const submitElem = document.getElementById("login-form-submit");
const errMsgElem = document.getElementById("login-error-msg");
const apiUrl = `/api/v1/auth/login`;

submitElem.addEventListener("click", (e) => {
  e.preventDefault();
  let id = formElem.id.value;
  let password = formElem.password.value;
  if (id && password) {
    loginApi(id, password);
  } else {
    errMsgElem.style.opacity = 1;
  }
});

function loginApi(id, password) {
  axios
    .post(apiUrl, {
      id: id,
      password: password,
    })
    .catch((err) => {
      switch (err.response.status) {
        case HttpStatusCode.FOUND:
          window.location.replace(err.response.data.domain);
          break;
        case HttpStatusCode.BAD_REQUEST:
          console.error("login error: input error");
          errMsgElem.style.opacity = 1;
          break;
        case HttpStatusCode.INTERNAL_SERVER_ERROR:
          alert("server has some problem");
          break;
      }
    });
}
