<?php
require_once('route.php');
$is_login = $_SESSION['token'];
$action_url = '';
if ($is_login) {
  $action_url = $Route['logout'];
} else {
  $action_url = $Route['login'];
}
?>
<header>
  <!-- TODO -->
  <div style="width: 150px;margin: 0 auto">
    <a id="Teacher" href="http://dns2.asia.edu.tw/~teacher_name/" target="_blank">
      <img src="../static/img/teacher_name.png">
    </a>
  </div>
  <!-- TODO -->
  <span style="position: absolute;top: 5px;right: 0;color: white;text-align: center;width: 130px;height: 65px">
    <form action="<?php echo $action_url ?>">
      <button type="submit">
        <?php
        if ($is_login) {
          echo `<i title="logout" class="sign-header-button fa-solid fa-arrow-right-from-bracket"></i>`;
        } else {
          echo `<i title="login" class="sign-header-button fa-solid fa-address-card"></i>`;
        }
        ?>
      </button>
    </form>
  </span>
</header>