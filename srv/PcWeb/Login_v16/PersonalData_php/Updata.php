<?php
include ('PcWeb/Login_v16/Login_v16/db_conn.inc.php');
$FileName=$_GET['edit_FileName'];
$ID=$_GET['edit_ID'];
$DATE=$_GET['edit_date'];
$CONTENT=$_GET['edit_content'];
switch ($FileName){
    case 1:
        $Type=$_GET['edit_type'];
        mysqli_query($link,"UPDATE `JournalPapers` 
SET `date`='$DATE',`content`='$CONTENT',`type`='$Type' WHERE `id`='$ID'");
        break;
    case 2:
        $Location=$_GET['edit_location'];
        $Partner=$_GET['edit_partner'];
        mysqli_query($link,"UPDATE `JournalPapers` 
SET `date`='$DATE',`content`='$CONTENT',`Location`='$Location',`partner`='$Partner' WHERE `id`='$ID'");
        break;
    case 3:
        $Partner=$_GET['edit_partner'];
        mysqli_query($link,"UPDATE `ConferencePapers` 
SET `date`='$DATE',`content`='$CONTENT',`partner`='$Partner' WHERE `id`='$ID'");
        break;
}

//mysqli_query($link,"UPDATE `$FileName` SET `date`='$DATE',`content`='$CONTENT' WHERE `id`='$ID'");
//echo "DELETE FROM `test1` WHERE `id`=1 ";
echo "<meta http-equiv=\"refresh\" content=\"0.1;url=./index.php\" />";
