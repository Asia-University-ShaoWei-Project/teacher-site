<?php require_once('route.php'); ?>
<div id="header">
  <div id="HOME_ICON" style="width: 150px;margin: 0 auto">
    <a id="Teacher" href="http://dns2.asia.edu.tw/~rikki/" target="_blank">
      <img src="../static/img/rikki.png">
    </a>
    <?php require_once('login.php'); ?>
  </div>
  <div class="Line"></div>
  <div class="Line"></div>
  <div id="a1"></div>
  <div id="a2"></div>
  <span style="position: absolute;top: 5px;right: 0;color: white;text-align: center;width: 130px;height: 65px">
    <?php if (!empty($_SESSION['token'])) { ?>
    <div>
      <form action=<?php $Route['logout'] ?>>
        <button type="submit" title="sign-out">
          <i class="fa fa-sign-out" style="position: absolute;top: 10px;font-size: 30px"></i>
        </button>
      </form>
    </div>
    <?php } else { ?>
    <i id="sign-in-header-button" class="fa fa-sign-in" title="sign-in"
      style="position: absolute;top: 10px;font-size: 30px"></i>
    <?php } ?>
  </span>
</div>