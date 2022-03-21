<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="keywords" content="Rikki">
    <title>RiKKi</title>
    <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
    <!--===============================================================================================-->
    <link rel="stylesheet" href="PcWeb/css/left-layout.css">
    <link rel="stylesheet" href="PcWeb/css/top-layout.css">
    <link rel="stylesheet" href="PcWeb/css/menu.css">
    <link rel="stylesheet" href="PcWeb/css/center-layout.css">
    <link rel="stylesheet" href="PcWeb/css/divBottom.css">
    <link rel="stylesheet" href="PcWeb/css/LoginLayout.css">
    <link rel="stylesheet" href="PcWeb/css/LineDecorate.css">
    <!--===============================================================================================-->
    <link rel="icon" type="image/png" href="PcWeb/Login_v16/Login_v16/images/icons/login.png"/>
    <!--===============================================================================================-->
    <link rel="stylesheet" type="text/css" href="PcWeb/Login_v16/Login_v16/vendor/bootstrap/css/bootstrap.min.css">
    <!--===============================================================================================-->
    <link rel="stylesheet" type="text/css" href="PcWeb/Login_v16/Login_v16/fonts/font-awesome-4.7.0/css/font-awesome.min.css">
    <!--===============================================================================================-->
    <link rel="stylesheet" type="text/css" href="PcWeb/Login_v16/Login_v16/fonts/Linearicons-Free-v1.0.0/icon-font.min.css">
    <!--===============================================================================================-->
    <link rel="stylesheet" type="text/css" href="PcWeb/Login_v16/Login_v16/vendor/animate/animate.css">
    <!--===============================================================================================-->
    <link rel="stylesheet" type="text/css" href="PcWeb/Login_v16/Login_v16/vendor/css-hamburgers/hamburgers.min.css">
    <!--===============================================================================================-->
    <link rel="stylesheet" type="text/css" href="PcWeb/Login_v16/Login_v16/vendor/animsition/css/animsition.min.css">
    <!--===============================================================================================-->
    <link rel="stylesheet" type="text/css" href="PcWeb/Login_v16/Login_v16/vendor/select2/select2.min.css">
    <!--===============================================================================================-->
    <link rel="stylesheet" type="text/css" href="PcWeb/Login_v16/Login_v16/vendor/daterangepicker/daterangepicker.css">
    <!--===============================================================================================-->
    <link rel="stylesheet" type="text/css" href="PcWeb/Login_v16/Login_v16/css/util.css">
    <link rel="stylesheet" type="text/css" href="PcWeb/Login_v16/Login_v16/css/copyMain.css">
    <!--===============================================================================================-->
    <script type="text/javascript" src="jQuery.js"></script>


    <style>
        #SignIn_i:hover{color: #f9830d;}
        /*          編輯頁面設定            */
        .index-news-list:hover,
        .index-news-list-box .index-news-list:last-child:hover{
            border: 2px solid #f6ac00;
        }
        .index-news-list:hover, .index-news-list-box .index-news-list:last-child:hover{
            border-left: 0;
            border-right: 0;
        }
        .index-news-list:hover .index-news-list-read{
            background-color: #f6ac00;
            border: 1px solid  #f6ac00;
        }
        .index-news-list:hover .index-news-list-read span{
            color: #FFF;
        }
        .news-list-box .index-news-list:last-child:hover{
            border-bottom: 2px solid #f6ac00;
        }
        /*.index-news-list:hover .knowledge-list-read{*/
        /*background-color: #f6ac00;*/
        /*}*/
        .index-banner-list span{
            width: 100%;
            display: block;
            height: 100%;
            background-repeat: no-repeat;
            background-size: cover;
            background-position: center;
        }
        .index-link a{
            width: 100%;
            display: -webkit-flex;
            display:         flex;
            -webkit-align-items: center;
            align-items: center;
            -webkit-justify-content: center;
            justify-content: center;
            flex-direction:column;
            height: 100%;
            padding: 20px;
            box-sizing: border-box;
            text-align: justify;
        }
        .index-link-mask span{
            width: 100%;
            height: 100%;
            border: 1px solid #f6ac00;
            display: block;
            background-color: rgba(0, 0, 0, 0.7);
            transition: all .3s linear;
            opacity: 0;
        }
        .index-news-title-box{
            max-width: 100%;
            display: block;
            margin: 0 auto;
            position: relative;
            height: 200px;
            line-height: 200px;
        }
        /*.index-news-title{*/
        /*font-size: 16px;*/
        /*font-weight: bold;*/
        /*text-align: center;*/
        /*}*/
        .index-news-icon{
            /*display: none;*/
            position: absolute;
            /*position: relative;*/
            bottom: 0;
            left: 40%;
            pointer-events: none;
            max-height: 151px;
        }
        .index-news-icon img{
            vertical-align: top;
        }

        .index-news-list-box{
            width: 100%;
            display: block;
        }
        .index-news-list{
            display: block;
            position: relative;
            border-top: 1px solid #dddddd;
            border-bottom: 0;
        }
        .index-news-list-box .index-news-list:last-child{
            border-bottom: 1px solid #dddddd;
        }
        .index-news-list-info-box{
            /*width: 1200px;*/
            width: 100%;
            max-width: 100%;
            display: block;
            margin: 0 auto;
            font-size: 0;
            padding: 20px 0 20px 10%;
        }
        .index-news-list-date{
            font-size: 20px;
            color: #555555;
            display: inline-block;
            vertical-align: middle;
            width: 10%;
            min-width: 110px;
            /*max-width: 100%;*/
            letter-spacing: 1px;
        }
        .index-news-list-title{
            font-size: 20px;
            /*overflow: auto;*/
            color: #555555;
            width: 60%;
            min-width: 500px;
            padding: 0 5%;
            display: inline-block;
            vertical-align: middle;
            letter-spacing: 1px;
            height: auto;
            /*overflow: ;*/
        }
        .index-news-list-personal{
            font-size: 20px;
            /*overflow: auto;*/
            color: #b47d2d;
            min-width: 600px;
            max-width: 600px;
            padding: 0 5%;
            display: inline-block;
            vertical-align: middle;
            letter-spacing: 1px;
            height: auto;

        }
        .index-news-list-read{
            width: 15%;
            min-width: 180px;
            height: 40px;
            display: inline-block;
            vertical-align: middle;
            text-align: center;
            border: 1px solid #333333;
            box-sizing: border-box;
            position: relative;
            transition: all .3s linear;
        }
        .index-news-list-read span{
            font-size: 16px;
            font-weight: bold;
            letter-spacing: 15px;
            /*position: absolute;*/
            /*left: 48px;*/
            /*top: calc(50% - 12.5px);*/
        }
        .index-news-list-date-edit{
            text-align: center;
            cursor: pointer;
            font-size: 16px;
            color: #f98312;
            display: inline-block;
            vertical-align: middle;
            min-width: 80px;
            /*max-width: 100%;*/
            letter-spacing: 1px;
        }
        .index-news-list-add{
            font-size: 20px;
            font-weight: bold;
            /*overflow: auto;*/
            color: #828282;
            width: 15%;
            /*padding: 0 10%;*/
            margin: 0 auto;
            display: inline-block;
            vertical-align: middle;
            letter-spacing: 1px;
            height: auto;
            /*overflow: ;*/
        }

        #index-news{
            padding-bottom: 50px;
        }
        .login100-form-btn-exit{
            font-family: Ubuntu-Bold;
            font-size: 16px;
            color: #fff;
            line-height: 1.2;
            text-transform: uppercase;
            background: transparent;
            display: -webkit-box;
            display: -webkit-flex;
            display: -moz-box;
            display: -ms-flexbox;
            display: flex;
            justify-content: center;
            align-items: center;
            padding: 0 20px;
            min-width: 160px;
            height: 35px;
            position: relative;
            z-index: 1;
        }
        .login100-form-btn-exit:hover{
            color: orange;
        }

        /*          一般設定            */
        a{
            text-decoration: none;
        }
        a:hover{
            text-decoration: none;
        }
        html{height:100%;

        } /* 頁面的高度設置 */
        * html body{height:100%;}
        body {
            margin: 0 auto;
            padding: 0;
            position: relative; /* 相對的位置 */
            background-color: #858585;
        }

        div #MessageAnnouncement,#PersonalInformation,
        #ComputerNetworks,#ComputerOrganization,
        #Profile,#Journal,#ConferencePapers,#ResearchGrant{
            display: none;
        }


        /*input[type=text], input[type=password] {*/
        /*width: 100%;*/
        /*padding: 12px 20px;*/
        /*margin: 8px 0;*/
        /*display: inline-block;*/
        /*border: 1px solid #ccc;*/
        /*box-sizing: border-box;*/
        /*}*/

        /* Set a style for all buttons */
        /*button {*/
        /*background-color: #4CAF50;*/
        /*color: white;*/
        /*padding: 14px 20px;*/
        /*margin: 8px 0;*/
        /*border: none;*/
        /*cursor: pointer;*/
        /*width: 100%;*/
        /*}*/

        a:hover {
            opacity: 0.8;
        }
        @keyframes AddWIDTH{
            from{width: 0}
            to{width: 100%}
        }
        div .Line{
            position: absolute;
            border: none;
            display: inline-block;
            vertical-align: middle;
            background-color: #f9830d;
        }
        div .Line:nth-child(2){
            height: 3px;width: 100%;max-width: 100%;
            top: 0;right: -1px;
            animation-name: AddWIDTH;
            animation-duration: 2s;
        }
        div .Line:nth-child(3){
            height: 4px;width: 100%;max-width: 100%;
            bottom: 2px;left: -1px;
            animation-name: AddWIDTH;
            animation-duration: 2s;
        }
        @keyframes top1 {
            from{top: -60px}
            to{top: -3px}
        }
        @keyframes top2 {
            from{bottom: -60px}
            to{bottom: 2px}
        }
        div #a1,#a2{
            position: absolute;
            border: none;
            display: inline-block;
            vertical-align: middle;
            background-color: #f9830d;
        }
        div #a1{height: 63px;width: 4px;max-height: 100%;
            top: -1px;right: 65%;
            animation-name: top1;
            animation-duration: 2s;}
        div #a2{height: 65px;width: 4px;max-height: 100%;
            bottom: 2px;left: 65%;
            animation-name: top2;
            animation-duration: 2s;
            z-index: 1;
        }
        @keyframes home {
            from{left: 0}
            to{left: 47%}
        }
        div #HOME_ICON{
            animation-name: home;
            animation-duration: 2s;
        }

        /*登入介面AND PDF介面*/
        #PDF_Button a{
            border: none;
            border-radius: 10px;
            padding: 15px 15px;
        }
        #PDF_Button a i:hover{
            opacity: 0.9;
        }
        #PDF_Button a:nth-child(1){background-color: DodgerBlue}
        #PDF_Button a:nth-child(2){background-color: #e2e6ff
        }
        .modal{
            display: none; /* Hidden by default */
            position: fixed; /* Stay in place */
            z-index: 1; /* Sit on top */
            left: 0;
            top: 0;
            width: 100%; /* Full width */
            height: auto; /* Full height */
            padding: auto 5% auto 5%;
            overflow: auto; /* Enable scroll if needed */
            background-color: rgb(0,0,0); /* Fallback color */
            background-color: rgba(0,0,0,0.3); /* Black w/ opacity */
        }


        .modal-content{
            position: absolute;
            background-color: #000000;
            border: 2px solid #ff7810;
            border-radius: 2px;
            top: 10%;
            width: 26%;
            height: 70%;
        }
        .modal-content-pdf {
            background-color: #1d1d1d;
            margin: 0 auto; /* 5% from the top, 15% from the bottom and centered */
            border: 2px solid #ff7810;
            border-radius: 2px;
            width: 95%;
            height: 100%;
        }
        .close {
            position: absolute;
            right: 25px;
            top: 0;
            color: #ffc274;
            font-size: 35px;
            font-weight: bold;
        }
        .close:hover{
            color: #ffffff;

        }
        .animate {
            -webkit-animation: animatezoom 0.6s;
            animation: animatezoom 0.6s
        }
        @-webkit-keyframes animatezoom {
            from {-webkit-transform: scale(0)}
            to {-webkit-transform: scale(1)}
        }
        @keyframes animatezoom {
            from {transform: scale(0)}
            to {transform: scale(1)}
        }


    </style>

</head>
<?php
session_start() ;
include('PcWeb/Login_v16/Login_v16/db_conn.inc.php');
if(isset($_SESSION['TeacherKey'])){
    $key=0;
}else{
    $key=1;
}?>

<body onload="WebSize()">
<!--Login DIV-->
<header id="main-header">

</header>
<section id="main-section">

</section>
<article id="main-article">

</article>
<aside id="main-aside">

</aside>
<footer id="main-footer">

</footer>





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
            <?php if($key==0){?>
                <div>
                    <form action="PcWeb/Login_v16/Login_v16/SignOut.php">
                        <button type="submit" title="SignOut">
                            <i id="SignOut_i" class="fa fa-sign-out"  style="position: absolute;top: 10px;font-size: 30px"></i>
                        </button>
                    </form>
                </div>

            <?php }else{?>
                <i id="SignIn_i" class="fa fa-sign-in" title="SignIn" style="position: absolute;top: 10px;font-size: 30px" ></i>
            <?php }?>
        </span>
</div>

<div id="DIV_BODY"
     style="width: 100%;position: relative;padding-top: 66px;z-index: 8;background: rgb(0,0,0)">

    <div id="LEFT" class="leftNav"
         style="position:fixed;z-index: 2">
        <br>
        <div>
            <img id="divImage2" src="Teacher.jpg" style="border-radius: 5%">
        </div>
        <br>
        <hr style="width: 70%;margin: 0 auto">
        <br>
        <div style="margin-left: 20%;text-align: left;font-size: 18px;">
            <p style="color: white">學歷： 國立中興大學資訊科學博士</p>
            <p style="color: white">辦公室： HB13</p>
            <p style="color: white">分機： 20013</p>
            <p style="color: white">E-mail： rikki@asia.edu.tw</p>
        </div>
        <br>

        <hr style="width: 70%;margin: 0 auto">
        <br>
        <table border="0" style="width: 100%;text-align:left;margin-left: 10%">
            <tr>
                <td class="Button" onclick="f('MessageAnnouncement')">
                    <span>訊   息   公   告</span>
                </td>
            </tr>
            <tr>
                <td class="Button" onclick="f('PersonalInformation')">
                    <span>個   人   資   料</span>
                </td>
            </tr>
            <tr>
                <td class="Button" onclick="f('ComputerNetworks')">
                    <span>計   算   機   網   路</span>
                </td>
            </tr>
            <tr>
                <td class="Button" onclick="f('ComputerOrganization')">
                    <span>計   算   機   組   織</span>
                </td>
            </tr>
            <tr>
                <td class="Button" onclick="f('A1')">
                    <span>微   處   理   器   系   統</span>
                </td>
            </tr>
            <tr>
                <td class="Button" onclick="f('A2')">
                    <span>多  媒  體  網  站  技  術</span>
                </td>
            </tr>
        </table>
        <br>

        <hr style="width: 70%;margin: 0 auto">
        <br>
        <p class="BottomText">任何建議請寄:  EMAIL rikki@asia.edu.tw</p>
        <p class="BottomText">亞洲大學資工系 陳瑞奇(Rikki Chen, CSIE, Asia Univ.)</p>
        <p class="BottomText">感謝您！</p>
    </div>

    <div id="CENTER" style="position: absolute">
        <div id="content" style="background-color: white">
            <!--  公告 列表區  -->
            <?php
            $Word = array("MessageAnnouncement","PersonalInformation","ComputerNetworks"
            ,"ComputerOrganization");
            $Chose = array("Bulletin_board","JournalPapers","ComputerNetworks"
            ,"ComputerOrganization");
            for($i=0;$i<4;$i++ ){?>
                <div id="<?php echo $Word[$i]?>" class="Chose">
                    <div id="index-news">
                        <div class="index-news-title-box">
                            <div class="index-news-icon">
                                <img src="PcWeb/image/<?php echo $Chose[$i]?>.jpg">
                            </div>
                        </div>
                        <?php
                        switch ($i){
                            case 0:
                                $Information = mysqli_query($link, "SELECT * FROM `Bulletin_board`");
                                break;
                            case 1:
                                $Information = mysqli_query($link, "SELECT * FROM `JournalPapers`");
                                break;
                            case 2:
                                $Information = mysqli_query($link, "SELECT * FROM `ComputerNetworks`");
                                break;
                            case 3:
                                $Information = mysqli_query($link, "SELECT * FROM `ComputerOrganization`");
                                break;
                        }
                        for($i=0;$i<mysqli_num_rows($Information);$i++){
                            $rs=mysqli_fetch_row($Information);
                            ?>
                            <a class="index-news-list">
                                <div class="index-news-list-info-box" style="display: none">
                                    <?php if($key==0){?>
                                        <h4 class="index-news-list-date-edit" onclick="Edit_BB('BulletinBoard','<?php echo $rs[0]?>','<?php echo $rs[1]?>','<?php echo $rs[2] ?>')">編輯</h4>
                                    <?php }?>
                                    <h4 class="index-news-list-date"><?php echo $rs[1]?></h4>
                                    <h3 class="index-news-list-title" style="width: 75%"><?php echo $rs[2] ?></h3>
                                    <?php if($key==0){?>
                                        <div class="index-news-list-date-edit">
                                            <form  action="PcWeb/Login_v16/Login_v16/DeleteData.php" method="get">
                                                <input type="hidden" name="DeleteFileName" value="BulletinBoard">
                                                <input type="hidden" name="DeleteID" value='<?php echo $rs[0]?>'>
                                                <button type="submit"><i class="fa fa-trash" style="color: black;"></i></button>
                                            </form>
                                        </div>
                                    <?php }?>
                                </div>
                            </a>
                            <?php
                        }
                        ?>
                        <?php
                        if($key==0) {
                            ?>
                            <a href="#" onclick="Insert_BB('BulletinBoard')" class="index-news-list">
                                <div class="index-news-list-info-box" style="display: none">
                                    <h3 class="index-news-list-add">新增(+)</h3>
                                </div>
                            </a>
                        <?php }?>
                    </div>
                </div>


                <?php
                if($i>0){

                }

            }
            ?>

            <!--  個人資料 列表區  -->
            <div id="PersonalInformation" class="Chose">
                <div id="index-news">
                    <div class="index-news-title-box">
                        <div class="index-news-icon">
                            <img src="PcWeb/image/JournalPapers.jpg">
                        </div>
                    </div>
                    <?php
                    $Information = mysqli_query($link, "SELECT * FROM `JournalPapers`");
                    for($i=0;$i<mysqli_num_rows($Information);$i++){
                        $rs=mysqli_fetch_row($Information);
                        ?>
                        <a class="index-news-list">
                            <div class="index-news-list-info-box" style="display: none">
                                <?php if($key==0){?>
                                    <h4 class="index-news-list-date-edit" onclick="Edit_Personal(1,'<?php echo $rs[0]?>','<?php echo $rs[1]?>','<?php echo $rs[2] ?>','<?php echo $rs[3]?>','<?php echo $rs[4]?>','','block','none')">編輯</h4>
                                <?php }?>
                                <h4 class="index-news-list-date"><?php echo $rs[1]?></h4>
                                <h3 class="index-news-list-personal" style="width: 75%"><?php echo $rs[2] ?></h3>
                                <h4 class="index-news-list-date"><?php echo $rs[3]?></h4>
                                <h4 class="index-news-list-date"><?php echo $rs[4]?></h4>
                                <?php if($key==0){?>
                                    <div class="index-news-list-date-edit">
                                        <form  action="PcWeb/Login_v16/Login_v16/DeleteData.php" method="get">
                                            <input type="hidden" name="DeleteFileName" value="JournalPapers">
                                            <input type="hidden" name="DeleteID" value='<?php echo $rs[0]?>'>
                                            <button type="submit"><i class="fa fa-trash" style="color: black;"></i></button>
                                        </form>
                                    </div>
                                <?php }?>
                            </div>
                        </a>
                        <?php
                    }?>
                    <?php
                    if($key==0) {
                        ?>
                        <a href="#" onclick="Insert_BB('JournalPapers')" class="index-news-list">
                            <div class="index-news-list-info-box" style="display: none">
                                <h3 class="index-news-list-add">新增(+)</h3>
                            </div>
                        </a>
                    <?php }?>
                </div>
                <hr style="background-color: #f9830d;width: 85%;margin: 0 auto">
                <div id="index-news">
                    <div class="index-news-title-box">
                        <div class="index-news-icon">
                            <img src="PcWeb/image/ConferencePapers.jpg">
                        </div>
                    </div>
                    <?php
                    $Information = mysqli_query($link, "SELECT * FROM `ConferencePapers`");
                    for($i=0;$i<mysqli_num_rows($Information);$i++){
                        $rs=mysqli_fetch_row($Information);
                        ?>
                        <a class="index-news-list">
                            <div class="index-news-list-info-box" style="display: none">
                                <?php if($key==0){?>
                                    <h4 class="index-news-list-date-edit" onclick="Edit_Personal(2,'<?php echo $rs[0]?>','<?php echo $rs[1]?>','<?php echo $rs[2]?>','<?php echo $rs[4]?>','','<?php echo $rs[3]?>','none','block')">編輯</h4>
                                <?php }?>
                                <h4 class="index-news-list-date"><?php echo $rs[1]?></h4>
                                <h3 class="index-news-list-personal" style="width: 75%"><?php echo $rs[2] ?></h3>
                                <h4 class="index-news-list-date"><?php echo $rs[3]?></h4>
                                <h4 class="index-news-list-date"><?php echo $rs[4]?></h4>
                                <?php if($key==0){?>
                                    <div class="index-news-list-date-edit">
                                        <form  action="PcWeb/Login_v16/Login_v16/DeleteData.php" method="get">
                                            <input type="hidden" name="DeleteFileName" value="ConferencePapers">
                                            <input type="hidden" name="DeleteID" value='<?php echo $rs[0]?>'>
                                            <button type="submit"><i class="fa fa-trash" style="color: black;"></i></button>
                                        </form>
                                    </div>
                                <?php }?>
                            </div>
                        </a>
                        <?php
                    }?>
                    <?php
                    if($key==0) {
                        ?>
                        <a href="#" onclick="Insert_BB('ConferencePapers')" class="index-news-list">
                            <div class="index-news-list-info-box" style="display: none">
                                <h3 class="index-news-list-add">新增(+)</h3>
                            </div>
                        </a>
                    <?php }?>
                </div>
                <hr style="background-color: #f9830d;width: 85%;margin: 0 auto">
                <div id="index-news">
                    <div class="index-news-title-box">
                        <div class="index-news-icon">
                            <img src="PcWeb/image/ResearchGrant.jpg">
                        </div>
                    </div>
                    <?php
                    $Information = mysqli_query($link, "SELECT * FROM `ResearchGrant`");
                    for($i=0;$i<mysqli_num_rows($Information);$i++){
                        $rs=mysqli_fetch_row($Information);
                        ?>
                        <a class="index-news-list">
                            <div class="index-news-list-info-box" style="display: none">
                                <?php if($key==0){?>
                                    <h4 class="index-news-list-date-edit" onclick="Edit_Personal(3,'<?php echo $rs[0]?>','<?php echo $rs[1]?>','<?php echo $rs[2] ?>','<?php echo $rs[3] ?>','','','none','none')">編輯</h4>
                                <?php }?>
                                <h4 class="index-news-list-date"><?php echo $rs[1]?></h4>
                                <h3 class="index-news-list-personal" style="max-width: 800px"><?php echo $rs[2] ?></h3>
                                <h4 class="index-news-list-date"><?php echo $rs[3]?></h4>
                                <?php if($key==0){?>
                                    <div class="index-news-list-date-edit">
                                        <form  action="PcWeb/Login_v16/Login_v16/DeleteData.php" method="get">
                                            <input type="hidden" name="DeleteFileName" value="ResearchGrant">
                                            <input type="hidden" name="DeleteID" value='<?php echo $rs[0]?>'>
                                            <button type="submit"><i class="fa fa-trash" style="color: black;"></i></button>
                                        </form>
                                    </div>
                                <?php }?>
                            </div>
                        </a>
                        <?php
                    }?>
                    <?php
                    if($key==0) {
                        ?>
                        <a href="#" onclick="Insert_BB('ResearchGrant')" class="index-news-list">
                            <div class="index-news-list-info-box" style="display: none">
                                <h3 class="index-news-list-add">新增(+)</h3>
                            </div>
                        </a>
                    <?php }?>
                </div>
            </div>
            <!--  計算機網路 列表區  -->
            <div id="ComputerNetworks" class="Chose">
                <div id="index-news">
                    <div class="index-news-title-box">
                        <div class="index-news-icon">
                            <img src="PcWeb/image/bulletin_board.jpg">
                        </div>
                    </div>
                    <?php
                    $Information = mysqli_query($link, "SELECT * FROM `IntroductionToComputerNetworks`");
                    $Course_slides = mysqli_query($link, "SELECT * FROM `Course_slides_1`");
                    $Homeworks = mysqli_query($link, "SELECT * FROM `Homeworks_1`");
                    for($i=0;$i<mysqli_num_rows($Information);$i++){
                        $rs=mysqli_fetch_row($Information);
                        ?>
                        <a class="index-news-list">
                            <div class="index-news-list-info-box" style="display: none">
                                <?php if($key==0){?>
                                    <h4 class="index-news-list-date-edit" onclick="Edit_BB('IntroductionToComputerNetworks','<?php echo $rs[0]?>','<?php echo $rs[1]?>','<?php echo $rs[2] ?>')">編輯</h4>
                                <?php }?>
                                <h4 class="index-news-list-date"><?php echo $rs[1]?></h4>
                                <h3 class="index-news-list-title" style="width: 75%"><?php echo $rs[2] ?></h3>
                                <?php if($key==0){?>
                                    <div class="index-news-list-date-edit">
                                        <form  action="PcWeb/Login_v16/Login_v16/DeleteData.php" method="get">
                                            <input type="hidden" name="DeleteFileName" value="IntroductionToComputerNetworks">
                                            <input type="hidden" name="DeleteID" value='<?php echo $rs[0]?>'>
                                            <button type="submit"><i class="fa fa-trash" style="color: black;"></i></button>
                                        </form>
                                    </div>
                                <?php }?>
                            </div>
                        </a>
                        <?php
                    }?>
                    <?php
                    if($key==0) {
                        ?>
                        <a href="#" onclick="Insert_BB('IntroductionToComputerNetworks')" class="index-news-list">
                            <div class="index-news-list-info-box" style="display: none">
                                <h3 class="index-news-list-add">新增(+)</h3>
                            </div>
                        </a>
                    <?php }?>
                </div>
                <hr style="background-color: #f9830d;width: 85%;margin: 0 auto">

                <div id="index-news">
                    <div class="index-news-title-box">
                        <div class="index-news-icon">
                            <img src="PcWeb/image/course_slides.jpg">
                        </div>
                    </div>

                    <?php
                    for($i=0;$i<mysqli_num_rows($Course_slides);$i++){
                        $rs=mysqli_fetch_row($Course_slides);
                        ?>
                        <?php $test=$rs[2]?>
                        <a onclick="test('<?php echo $test ?>',0,'Computerworks')" class="index-news-list">
                            <div class="index-news-list-info-box" style="display: none">
                                <h4 class="index-news-list-date">Ch:<?php echo $rs[0]?></h4>
                                <h3 class="index-news-list-title"><?php echo $rs[1] ?></h3>
                                <div class="index-news-list-read">
                                    <span >READ</span>
                                </div>
                            </div>
                        </a>
                        <?php
                    }
                    ?>
                </div>
                <hr style="background-color: #f9830d;width: 85%;margin: 0 auto">
                <div id="index-news">
                    <div class="index-news-title-box">
                        <div class="index-news-icon">
                            <img src="PcWeb/image/homeworks.jpg">
                            <!--            <img src="bulletin-board-ch.jpg">-->
                        </div>
                    </div>

                    <!--  C的列表區  -->
                    <?php
                    for($i=0;$i<mysqli_num_rows($Homeworks);$i++){
                        $rs=mysqli_fetch_row($Homeworks);
                        ?>
                        <?php $test=$rs[2]?>
                        <a onclick="test('<?php echo $test ?>',1,'Computerworks')" class="index-news-list">
                            <div class="index-news-list-info-box" style="display: none">
                                <h4 class="index-news-list-date"><?php echo $rs[0]?></h4>
                                <h3 class="index-news-list-title"><?php echo $rs[1] ?></h3>
                                <div class="index-news-list-read">
                                    <span >READ</span>
                                </div>
                            </div>
                        </a>
                        <?php
                    }
                    ?>
                </div>
            </div>
            <!--  計算機組織 列表區  -->
            <div id="ComputerOrganization" class="Chose">
                <div id="index-news">
                    <div class="index-news-title-box">
                        <div class="index-news-icon">
                            <img src="PcWeb/image/bulletin_board.jpg">
                        </div>
                    </div>
                    <?php
                    $Information = mysqli_query($link, "SELECT * FROM `ComputerOrganization`");
                    $Course_slides = mysqli_query($link, "SELECT * FROM `Course_slides_2`");
                    $Homeworks = mysqli_query($link, "SELECT * FROM `Homeworks_1`");
                    for($i=0;$i<mysqli_num_rows($Information);$i++){
                        $rs=mysqli_fetch_row($Information);
                        ?>
                        <a class="index-news-list">
                            <div class="index-news-list-info-box" style="display: none">
                                <?php if($key==0){?>
                                    <h4 class="index-news-list-date-edit" onclick="Edit_BB('ComputerOrganization','<?php echo $rs[0]?>','<?php echo $rs[1]?>','<?php echo $rs[2] ?>')">編輯</h4>
                                <?php }?>
                                <h4 class="index-news-list-date"><?php echo $rs[1]?></h4>
                                <h3 class="index-news-list-title" style="width: 75%"><?php echo $rs[2] ?></h3>
                                <?php if($key==0){?>
                                    <div class="index-news-list-date-edit">
                                        <form  action="PcWeb/Login_v16/Login_v16/DeleteData.php" method="get">
                                            <input type="hidden" name="DeleteFileName" value="ComputerOrganization">
                                            <input type="hidden" name="DeleteID" value='<?php echo $rs[0]?>'>
                                            <button type="submit"><i class="fa fa-trash" style="color: black;"></i></button>
                                        </form>
                                    </div>
                                <?php }?>
                            </div>
                        </a>
                        <?php
                    }?>
                    <?php
                    if($key==0) {
                        ?>
                        <a href="#" onclick="Insert_BB('ComputerOrganization')" class="index-news-list">
                            <div class="index-news-list-info-box" style="display: none">
                                <h3 class="index-news-list-add">新增(+)</h3>
                            </div>
                        </a>
                    <?php }?>
                </div>
                <hr style="background-color: #f9830d;width: 85%;margin: 0 auto">

                <div id="index-news">
                    <div class="index-news-title-box">
                        <div class="index-news-icon">
                            <img src="PcWeb/image/course_slides.jpg">
                        </div>
                    </div>
                    <?php
                    for($i=0;$i<mysqli_num_rows($Course_slides);$i++){
                        $rs=mysqli_fetch_row($Course_slides);
                        ?>
                        <?php $test=$rs[2]?>
                        <a onclick="test('<?php echo $test ?>',0,'ComputerOrganization')" class="index-news-list">
                            <div class="index-news-list-info-box" style="display: none">
                                <h4 class="index-news-list-date">Ch:<?php echo $rs[0]?></h4>
                                <h3 class="index-news-list-title"><?php echo $rs[1] ?></h3>
                                <div class="index-news-list-read">
                                    <span >READ</span>
                                </div>
                            </div>
                        </a>
                        <?php
                    }
                    ?>
                </div>
                <hr style="background-color: #f9830d;width: 85%;margin: 0 auto">
                <div id="index-news">
                    <div class="index-news-title-box">
                        <div class="index-news-icon">
                            <img src="PcWeb/image/homeworks.jpg">
                            <!--            <img src="bulletin-board-ch.jpg">-->
                        </div>
                    </div>
                    <!--  C的列表區  -->
                    <?php
                    for($i=0;$i<mysqli_num_rows($Homeworks);$i++){
                        $rs=mysqli_fetch_row($Homeworks);
                        ?>
                        <?php $test=$rs[2]?>
                        <a onclick="test('<?php echo $test ?>',1,'ComputerOrganization')" class="index-news-list">
                            <div class="index-news-list-info-box" style="display: none">
                                <h4 class="index-news-list-date"><?php echo $rs[0]?></h4>
                                <h3 class="index-news-list-title"><?php echo $rs[1] ?></h3>
                                <div class="index-news-list-read">
                                    <span >READ</span>
                                </div>
                            </div>
                        </a>
                        <?php
                    }
                    ?>
                </div>
            </div>
            <div id="A1" class="Chose" style="display: none;"><h1>還沒做sorry</h1></div>
            <div id="A2" class="Chose" style="display: none;"><h1>還沒做sorry</h1></div>
        </div>
    </div>
    <script type="text/javascript">
        function test(FileName,Num,Chose) {
            var a= document.getElementById('PDF');
            var a1= document.getElementById('iframe-pdf');
            try{
                //0是上 1下
                switch (Num) {
                    case 0:{
                        a1.src="PcWeb/CourseSlides/"+Chose+"/"+FileName;
                        break
                    }
                    case 1:{
                        a1.src="PcWeb/Homeworks/"+Chose+"/Hw"+FileName;
                        break
                    }
                }
            }catch (e) {
                alert("?"+e.getErrorMessage())
            }
            a.style.display='block';
        }

    </script>
    <script type="text/javascript">
        function Edit_BB(FileName,ID,date,content) {
            var a= document.getElementById('Edit_BB');
            document.getElementById('edit_FileName').value=FileName;
            document.getElementById('edit_ID').value=ID;
            document.getElementById('edit_date').value=date;
            document.getElementById('edit_content').value=content;
            a.style.display='block';
        }
        function Insert_BB(FileName) {
            var a= document.getElementById('Insert_BB');
            document.getElementById('insert_FileName').value=FileName;
            a.style.display='block';
        }
        function Edit_Personal(FileName,ID,date,content,partner,type,location,typeView,locationView) {
            // alert("Mass:"+FileName+ID+date+content+partner+type+location+typeView+locationView)
            var a= document.getElementById('Edit-Personal-JournalPapers');
            document.getElementById('Personal_FileName').value=FileName;
            document.getElementById('Personal_ID').value=ID;
            document.getElementById('Personal_date').value=date;
            document.getElementById('Personal_content').value=content;
            document.getElementById('Personal_type').value=type;
            document.getElementById('Personal_location').value=location;
            document.getElementById('Personal_partner').value=partner;
            // <!--=========================================================-->
            // document.getElementById('h4-type').style.display=typeView;
            // document.getElementById('h4-location').style.display=locationView;
            // <!--=========================================================-->
            // document.getElementById('edit-type').style.display=typeView;
            // document.getElementById('edit_location').style.display=locationView;
            a.style.display='block';
        }
        function Insert_Personal(FileName) {
            var a= document.getElementById('Insert_BB');
            document.getElementById('insert_FileName').value=FileName;
            a.style.display='block';
        }
    </script>
    <script>
        $(".Button").click(function(){
            $(".index-news-icon").fadeIn(1400);
            $(".index-news-list-info-box").fadeIn(1950);
            // $("#CENTER").fadeIn(1950);

        });
    </script>
    <script type="text/javascript">
        $(document).ready(function(){
            $("#bt_Exit").click(function () {
                $("#SignIn").fadeOut(1000);
            });
            $("#SignIn_i").click(function(){
                $("#SignIn").fadeIn(1200);
            });
        });
    </script>
    <script>
        window.onresize=function () {WebSize()};
        function deBUG(value) {
            try{
                // var VAR=document.getElementById("");
                alert(value);
            }catch (e) {
                alert("Error"+e.getErrorMessage());
            }
        }
        function WebSize() {
            try{
                var Width = $(window).width();
                if(Width>=1049){
                    if(Width>=1501){
                        ReSet_size(Width*0.26,Width*0.74);
                    }else {
                        ReSet_size(Width*0.3,Width*0.7);}
                }
                else {
                    ReSet_size(0,Width);
                }
            }catch (e) {
                alert("Error"+e.getErrorMessage());

            }
        }
        function ReSet_size(LEFT_W,RIGHT_W){
            var DIV_Left = document.getElementById("LEFT");
            DIV_Left.style.width=LEFT_W+"px";
            var DIV_Center = document.getElementById("CENTER");
            // var DIV_Center1 = document.getElementById("content");
            DIV_Center.style.width=RIGHT_W+"px";
            DIV_Center.style.maxWidth=RIGHT_W+"px";
            // DIV_Center1.style.width=RIGHT_W+"px";
            // DIV_Center1.style.maxWidth=RIGHT_W+"px";
            // DIV_Center.style.maxWidth=center_width+"px";
            DIV_Center.style.left=DIV_Left.style.width;
            // alert(DIV_Center.style.left);
        }
        //滾輪移動
        // function scrollFunction() {
        //     if(document.body.scrollTop>20 || document.documentElement.scrollTop>20){
        //         document.getElementById("header").style.height="80px";
        //     }else {
        //         document.getElementById("header").style.height="65px";
        //     }
        // }
        function f(who) {
            var a =document.getElementById(who);
            $(".Chose").css("display","none");
            $("#"+who).fadeIn(1300);
            a.style.display='block';
            // $(".index-news-icon").fadeIn(1400);
            // $(".index-news-list-info-box").fadeIn(1950);
        }
    </script>


    <!--BB編輯-->
    <div id="Edit_BB" class="modal" style="z-index: 99">
        <div class="modal-content animate">
            <div style="height: 90%;position: relative;">
                    <span style="position: absolute;right: 2.5%;top: 30px" onclick="document.getElementById('Edit_BB').style.display='none'"
                          class="close"  title="Close Modal">X</span>
                <div id="Title" style="position: absolute;"></div>
                <div id="InputContent" style="position: relative;">
                    <h1>Edit</h1>
                    <form action="PcWeb/Login_v16/Login_v16/UpData.php" method="GET" enctype="multipart/form-data" style="position: absolute;">
                        <input id="edit_FileName" type="hidden" name="edit_FileName" ">
                        <input id="edit_ID" type="hidden" name="edit_ID" ">
                        <h4 style="color: orange;">Date:</h4>
                        <input id="edit_date" type="text" name="edit_date" style="width: 200px;left: 30%;top: 5%;">
                        <h4 style="color: orange;">Content:</h4>
                        <textarea id="edit_content" name="edit_content" style="width: 400px;height: 400px;"></textarea>
                        <br>
                        <input type="submit" name="submit" value="修改">
                    </form>
                </div>
            </div>
        </div>
    </div>
    <!--BB新增-->
    <div id="Insert_BB" class="modal" style="z-index: 99">
        <div class="modal-content animate">
            <div style="height: 90%;position: relative;">
            <span style="position: absolute;right: 2.5%;top: 30px" onclick="document.getElementById('Insert_BB').style.display='none'"
                  class="close"  title="Close Modal">X
            </span>
                <div id="Title" style="position: absolute;">
                </div>
                <div id="InputContent" style="margin: 0 auto">
                    <form action="PcWeb/Login_v16/Login_v16/InsertData.php" method="GET" enctype="multipart/form-data" style="position: absolute;">
                        <input id="insert_FileName" type="hidden" name="insert_FileName"">
                        <h4 style="color: orange;">Date:</h4>
                        <input id="insert_date" type="text" name="insert_date" style="width: 200px;left: 30%;top: 5%;">
                        <h4 style="color: orange;">Content:</h4>
                        <textarea id="insert_content" name="insert_content" style="width: 400px;height: 400px;"></textarea>
                        <br>
                        <input type="submit" name="submit" value="新增">
                    </form>
                </div>
            </div>
        </div>
    </div>
    <!--Personal編輯-->
    <div id="Edit-Personal-JournalPapers" class="modal" style="z-index: 99">
        <div class="modal-content animate">
            <div style="height: 90%;position: relative;">
                    <span style="position: absolute;right: 2.5%;top: 30px" onclick="document.getElementById('Edit-Personal-JournalPapers').style.display='none'"
                          class="close"  title="Close Modal">X</span>
                <div style="position: relative;">
                    <h1 style="color: white;">Edit</h1>
                    <form action="PcWeb/Login_v16/Login_v16/UpData.php" method="GET" style="position: absolute;">
                        <input id="Personal_FileName" type="hidden" name="edit_FileName" ">
                        <input id="Personal_ID" type="hidden" name="edit_ID" ">
                        <h4 style="color: orange;">Date:</h4>
                        <input id="Personal_date" type="text" name="edit_date" style="width: 200px">
                        <h4 style="color: orange;">Content:</h4>
                        <textarea id="Personal_content" name="edit_content" style="width: 400px;height: 200px;"></textarea>
                        <br>


                        <h4 id="h4-type" style="color: orange;">Type:</h4>
                        <input id="Personal_type" type="text" name="edit_type" style="width: 200px">
                        <h4 id="h4-location" style="color: orange;">Location:</h4>
                        <input id="Personal_location" type="text" name="edit_location" style="width: 200px">
                        <h4 id="h4-partner" style="color: orange;">Partner:</h4>
                        <input id="Personal_partner" type="text" name="edit_partner" style="width: 400px;height:70px">
                        <input type="submit" name="submit" value="修改">
                        <button onclick="document.getElementById('Edit-Personal-JournalPapers').style.display='none'" style="color: white;">退出</button>

                    </form>
                </div>
            </div>
        </div>
    </div>

    <!--Personal新增-->
    <div id="Insert_Personal" class="modal" style="z-index: 99">
        <div class="modal-content animate">
            <div style="height: 90%;position: relative;">
            <span style="position: absolute;right: 2.5%;top: 30px" onclick="document.getElementById('Insert_BB').style.display='none'"
                  class="close"  title="Close Modal">X
            </span>
                <div id="Title" style="position: absolute;">
                </div>
                <div id="InputContent" style="margin: 0 auto">
                    <form action="PcWeb/Login_v16/Login_v16/InsertData.php" method="GET" enctype="multipart/form-data" style="position: absolute;">
                        <input id="insert_FileName" type="hidden" name="insert_FileName"">
                        <h4 style="color: orange;">Date:</h4>
                        <input id="insert_date" type="text" name="insert_date" style="width: 200px;left: 30%;top: 5%;">
                        <h4 style="color: orange;">Content:</h4>
                        <textarea id="insert_content" name="insert_content" style="width: 400px;height: 400px;"></textarea>
                        <br>
                        <input type="submit" name="submit" value="新增">
                    </form>
                </div>
            </div>
        </div>
    </div>




    <div id="SignIn" class="limiter">
        <div class="container-login100" style="background: transparent">
            <div class="wrap-login100 p-t-30 p-b-50" style="width: 500px">
				<span class="login100-form-title p-b-41">
					Rikki  Web  Login
				</span>
                <form class="login100-form validate-form p-b-33 p-t-5" action="PcWeb/Login_v16/Login_v16/LoginChack.php" method="post">
                    <div class="wrap-input100 validate-input" data-validate = "Enter username">
                        <input class="input100" type="email" name="email" placeholder="email">
                        <span class="focus-input100" data-placeholder="&#xe82a;"></span>
                    </div>
                    <div class="wrap-input100 validate-input" data-validate="Enter password">
                        <input class="input100" type="password" name="password" placeholder="Password">
                        <span class="focus-input100" data-placeholder="&#xe80f;"></span>
                    </div>
                    <div class="container-login100-form-btn m-t-32">
                        <button id="bt_SignIn" type="submit" class="login100-form-btn">
                            Sign in
                        </button>
                    </div>
                </form>
                <div class="container-login100-form-btn m-t-32">
                    <button id="bt_Exit" type="submit" class="login100-form-btn-exit">
                        Exit
                    </button>
                </div>
            </div>
        </div>
    </div>




    <div id="PDF" class="modal" style="z-index: 99">
        <form class="modal-content-pdf animate">
            <div style="height: 90%;position: relative;">
                <!--            <iframe src=""></iframe>-->
                <iframe id="iframe-pdf" style="height: 100%;width: 90%;border:1px solid DodgerBlue;margin: 2% 3% auto 3%"></iframe>
                <span style="position: absolute;right: 2.5%;top: 30px" onclick="document.getElementById('PDF').style.display='none'"
                      class="close"  title="Close Modal">X</span>
                <span id="PDF_Button" style="line-height:100px;position: absolute;margin: 80px auto">
                <a href="#" onclick="document.getElementById('PDF')
                .style.display='none'"><i class="fa fa-download" title="Download"></i></a>
                <br>
                <a href="#" onclick="document.getElementById('PDF')
                .style.display='none'"><i class="fa fa-expand" title="Amplification"></i></a>
                <br>
            </span>
            </div>
        </form>
    </div>
    <?php
    mysqli_close($link)
    ?>

</body>
</html>

