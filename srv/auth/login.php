<html>

<body style="
width: 100%;height: 100%;
">
  <img src="signIn.gif" style="width: 30%;margin: 0 auto">
</body>

</html>
<?php session_start(); ?>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<?php
include('db_conn.inc.php');

$Email = $_POST['email'];
$Password = sha1($_POST['password']);
$sql = "SELECT `ID`, `PASSWORD` FROM `TeacherUser` WHERE `ID` ='$Email' AND `PASSWORD` ='$Password'";
$result = mysqli_query($link, $sql);
$row = mysqli_fetch_row($result);
if ($row[0] == $Email and $row[1] == $Password) {
  $_SESSION['TeacherKey'] = $Email;
  echo "<meta http-equiv=\"refresh\" content=\"1;url=./index1.php\" />";
} else {
  echo "<meta http-equiv=\"refresh\" content=\"0.2;url=./index1.php\" />";
}