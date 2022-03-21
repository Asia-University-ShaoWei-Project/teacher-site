<?php session_start();

unset($_SESSION['TeacherKey']);
echo "登出中...";
echo "<meta http-equiv=\"refresh\" content=\"0.1;url=./index1.php\" />";