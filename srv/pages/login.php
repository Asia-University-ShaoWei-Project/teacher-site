<?php require_once('route.php'); ?>
<!-- ref: https://medium.com/swlh/how-to-create-your-first-login-page-with-html-css-and-javascript-602dd71144f1 -->
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Login</title>
  <link rel="stylesheet" type="text/css" href="../static/css/index/login.css">
</head>

<body>
  <main id="main-holder">
    <h1 id="login-header">Login</h1>

    <div id="login-error-msg-holder">
      <p id="login-error-msg">Invalid ID <span id="error-msg-second-line">and/or password</span></p>
    </div>

    <form id="login-form">
      <input type="text" name="id" id="id-field" class="login-form-field" placeholder="ID">
      <input type="password" name="password" id="password-field" class="login-form-field" placeholder="Password">
      <input type="submit" value="Login" id="login-form-submit">
    </form>

  </main>
</body>

</html>
<script src="../static/js/jquery/jquery-3.6.0.min.js"></script>
<script>
const loginForm = document.getElementById("login-form");
const loginButton = document.getElementById("login-form-submit");
const loginErrorMsg = document.getElementById("login-error-msg");

loginButton.addEventListener("click", (e) => {
  e.preventDefault();
  const id = loginForm.id.value;
  const password = loginForm.password.value;
  if (id && password) {
    auth_api(id, password);
  } else {
    loginErrorMsg.style.opacity = 1;
  }
})
const login_url = "http://<?php echo $Domain . $Route['login'] ?>";

function auth_api(id, password) {
  $.ajax({
      method: "POST",
      url: login_url,
      data: {
        id: id,
        password: password
      },
      dataType: 'json'
    })
    .done(function(response) {
      if (response.status == 200) {
        window.location.replace("http://<?php echo $Domain ?>");
      }
    })
    .fail(function(msg) {
      loginErrorMsg.style.opacity = 1;
      console.log(msg);
    });
}
</script>