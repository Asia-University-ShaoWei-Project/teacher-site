<?php
require_once('route.php');
?>


<div id="sign-in-frame" class="limiter">
  <div class="container-login100" style="background: transparent">
    <div class="wrap-login100 p-t-30 p-b-50" style="width: 500px">
      <span class="login100-form-title p-b-41">
        Rikki Web Login
      </span>
      <form class="login100-form validate-form p-b-33 p-t-5" action=<?php echo $Route['login'] ?> method="POST">
        <div class="wrap-input100 validate-input" data-validate="Enter username">
          <input class="input100" type="text" name="id" placeholder="ID">
          <span class="focus-input100" data-placeholder="&#xe82a;"></span>
        </div>
        <div class="wrap-input100 validate-input" data-validate="Enter password">
          <input class="input100" type="password" name="password" placeholder="Password">
          <span class="focus-input100" data-placeholder="&#xe80f;"></span>
        </div>
        <div class="container-login100-form-btn m-t-32">
          <button id="sign-in-submit-button" type="submit" class="login100-form-btn">
            Login
          </button>
        </div>
      </form>
      <div class="container-login100-form-btn m-t-32">
        <button id="sign-in-exit-button" type="submit" class="login100-form-btn-exit">
          Exit
        </button>
      </div>
    </div>
  </div>
</div>