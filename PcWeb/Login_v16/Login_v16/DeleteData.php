<?php
include ('db_conn.inc.php');
$FileName=$_GET['DeleteFileName'];
$ID=$_GET['DeleteID'];
mysqli_query($link,"DELETE FROM `$FileName` WHERE `id`='$ID' ");
//echo "DELETE FROM `test1` WHERE `id`=1 ";
echo "<meta http-equiv=\"refresh\" content=\"0.1;url=./index1.php\" />";
