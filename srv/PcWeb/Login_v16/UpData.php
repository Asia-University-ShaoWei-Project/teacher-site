<?php
include ('db_conn.inc.php');
$FileName=$_GET['edit_FileName'];
$ID=$_GET['edit_ID'];
$DATE=$_GET['edit_date'];
$CONTENT=$_GET['edit_content'];
mysqli_query($link,"UPDATE `$FileName` SET `date`='$DATE',`content`='$CONTENT' WHERE `id`='$ID'");
//echo "DELETE FROM `test1` WHERE `id`=1 ";
echo "<meta http-equiv=\"refresh\" content=\"0.1;url=./index.php\" />";
