<div id="header">
  <div id="HOME_ICON" style="width: 150px;margin: 0 auto">
    <a id="Teacher" href="http://dns2.asia.edu.tw/~rikki/" target="_blank">
      <img src="PcWeb/image/rikki.png">
    </a>
  </div>
  <div class="Line"></div>
  <div class="Line"></div>
  <div id="a1"></div>
  <div id="a2"></div>
  <span style="position: absolute;top: 5px;right: 0;color: white;text-align: center;width: 130px;height: 65px">
    <?php if ($key == 0) { ?>
    <div>
      <form action="PcWeb/Login_v16/Login_v16/SignOut.php">
        <button type="submit" title="SignOut">
          <i id="SignOut_i" class="fa fa-sign-out" style="position: absolute;top: 10px;font-size: 30px"></i>
        </button>
      </form>
    </div>

    <?php } else { ?>
    <i id="SignIn_i" class="fa fa-sign-in" title="SignIn" style="position: absolute;top: 10px;font-size: 30px"></i>
    <?php } ?>
  </span>
</div>