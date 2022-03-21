<?php
include ('db_conn.inc.php');
$FileName=$_GET['insert_FileName'];
$DATE=$_GET['insert_date'];
$CONTENT=$_GET['insert_content'];
mysqli_query($link,"INSERT INTO `$FileName`(`date`, `content`) VALUES ('$DATE','$CONTENT')");
//echo "DELETE FROM `test1` WHERE `id`=1 ";
echo "<meta http-equiv=\"refresh\" content=\"0.1;url=./index1.php\" />";
