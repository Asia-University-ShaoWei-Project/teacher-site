<?php include('PcWeb/Login_v16/Login_v16/db_conn.inc.php');
?>
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <style>
    .index-news-list:hover,
    .index-news-list-box .index-news-list:last-child:hover {
      border: 1px solid #f6ac00;
    }

    .index-news-list:hover,
    .index-news-list-box .index-news-list:last-child:hover {
      border-left: 0;
      border-right: 0;
    }

    .news-list-box .index-news-list:last-child:hover {
      border-bottom: 1px solid #f6ac00;
    }

    /*.index-news-list:hover .knowledge-list-read{*/
    /*background-color: #f6ac00;*/
    /*}*/
    .index-banner-list span {
      width: 100%;
      display: block;
      height: 100%;
      background-repeat: no-repeat;
      background-size: cover;
      background-position: center;
    }

    .index-link a {
      width: 100%;
      display: -webkit-flex;
      display: flex;
      -webkit-align-items: center;
      align-items: center;
      -webkit-justify-content: center;
      justify-content: center;
      flex-direction: column;
      height: 100%;
      padding: 20px;
      box-sizing: border-box;
      text-align: justify;
    }

    .index-link-mask span {
      width: 100%;
      height: 100%;
      border: 1px solid #f6ac00;
      display: block;
      background-color: rgba(0, 0, 0, 0.7);
      transition: all .3s linear;
      opacity: 0;
    }

    .index-news-title-box {
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
    .index-news-icon {
      display: none;
      position: absolute;
      bottom: 0;
      left: 0;
      pointer-events: none;
      max-height: 151px;
    }

    .index-news-icon img {
      vertical-align: top;
    }

    .index-news-list-box {
      width: 100%;
      display: block;
    }

    .index-news-list {
      display: block;
      position: relative;
      border-top: 1px solid #EEEEEE;
      border-bottom: 0;
    }

    .index-news-list-box .index-news-list:last-child {
      border-bottom: 1px solid #EEE;
    }

    .index-news-list-info-box {
      width: 1200px;
      max-width: 100%;
      display: block;
      margin: 0 auto;
      font-size: 0;
      padding: 20px 0;
    }

    .index-news-list-date {
      font-size: 18px;
      color: #555555;
      display: inline-block;
      vertical-align: middle;
      width: 180px;
      max-width: 100%;
      letter-spacing: 1px;
    }

    .index-news-list-title {
      font-size: 18px;
      overflow: auto;
      color: #555555;
      width: calc(100% - 360px);
      min-width: 200px;
      display: inline-block;
      vertical-align: middle;
      letter-spacing: 1px;
      height: 22px;
      /*overflow: hidden;*/
    }

    .index-news-list-read {
      width: 90px;
      height: 40px;
      display: inline-block;
      vertical-align: middle;
      max-width: 100%;
      text-align: center;
      border: 1px solid #333333;
      box-sizing: border-box;
      position: relative;
      transition: all .3s linear;
    }

    .index-news-list-read span {
      font-size: 16px;
      font-weight: bold;
      letter-spacing: 15px;
      position: absolute;
      left: 48px;
      top: calc(50% - 12.5px);
    }

    .index-news-list-read span :hover {
      background: black;
    }

    #index-news {
      padding-bottom: 50px;
    }

    #PDF_Button a {
      border: none;
      border-radius: 10px;
      padding: 15px 15px;
    }

    #PDF_Button a i:hover {
      opacity: 0.9;
    }

    #PDF_Button a:nth-child(1) {
      background-color: DodgerBlue
    }

    #PDF_Button a:nth-child(2) {
      background-color: #e2e6ff
    }

    .modal {
      display: none;
      /* Hidden by default */
      position: fixed;
      /* Stay in place */
      z-index: 1;
      /* Sit on top */
      left: 0;
      top: 0;
      width: 100%;
      /* Full width */
      height: auto;
      /* Full height */
      padding: auto 5% auto 5%;
      overflow: auto;
      /* Enable scroll if needed */
      background-color: rgb(0, 0, 0);
      /* Fallback color */
      background-color: rgba(0, 0, 0, 0.3);
      /* Black w/ opacity */
    }

    .modal-content {
      background-color: #1d1d1d;
      border: 2px solid #ff7810;
      border-radius: 2px;
      top: 10%;
      width: 26%;
      height: 70%;
    }

    .close {
      position: absolute;
      right: 25px;
      top: 0;
      color: #ffc274;
      font-size: 35px;
      font-weight: bold;
    }

    .close:hover {
      color: #ffffff;

    }

    .animate {
      -webkit-animation: animatezoom 0.6s;
      animation: animatezoom 0.6s
    }

    @-webkit-keyframes animatezoom {
      from {
        -webkit-transform: scale(0)
      }

      to {
        -webkit-transform: scale(1)
      }
    }

    @keyframes animatezoom {
      from {
        transform: scale(0)
      }

      to {
        transform: scale(1)
      }
    }
  </style>

</head>
<?php
include('Login_v16/Login_v16/db_conn.inc.php');
$value = $_GET['value'];
$InformationA = '';
$InformationB = '';
$InformationC = '';
$test = '';
$Who_1 = '';
if (empty($_GET['value'])) {
  $value = 1;
}
?>

<body>


  <div id="edit-box" class="modal" style="z-index: 99">
    <form class="modal-content animate">
      <div style="height: 90%;position: relative;">
        <span style="position: absolute;right: 2.5%;top: 30px" onclick="document.getElementById('Edit').style.display='none'" class="close" title="Close Modal">X</span>
        <div id="Title" style="position: absolute;">
        </div>
        <div id="InputContent" style="position: absolute;">
          <h1>1123456</h1>
          <form action="testUPdata.php" method="POST" enctype="multipart/form-data">
            <input type="file" name="file" id="file">
            <br>
            <input type="submit" name="submit" value="上傳">
          </form>

        </div>
        <div style="position: absolute;">

        </div>

      </div>
    </form>
  </div>

  <?php
  switch ($value) {
    case 0: {
        $InformationA = mysqli_query($link, "SELECT * FROM `BulletinBoard`");
        break;
      }
    case 1: {
        $InformationA = mysqli_query($link, "SELECT * FROM `IntroductionToComputerNetworks`");
        $InformationB = mysqli_query($link, "SELECT * FROM `Course_slides_1`");
        $InformationC = mysqli_query($link, "SELECT * FROM `Homeworks_1`");
        $Who_2 = 'Computerworks';
        break;
      }
    case 2: {
        $InformationA = mysqli_query($link, "SELECT * FROM `ComputerOrganization` ");
        $Who_2 = "ComputerOrg";
        break;
      }
    case 3: {
        $InformationA = mysqli_query($link, "SELECT * FROM `MicroprocessorSystems` ");
        break;
      }
    case 4: {
        $InformationA = mysqli_query($link, "SELECT * FROM `MicroprocessorSystems` ");
        break;
      }
  }
  ?>


  <script type="text/javascript">
    function Edit(Who, Num, Select) {
      var elem_edit_box = document.getElementById('edit-box');
      var elem_iframe_pdf = document.getElementById('iframe-pdf');
      try {
        switch (Select) {
          case 1: {
            break
          }
          case 2: {
            break
          }
          case 3: {
            break
          }
        }
      } catch (e) {
        alert("?" + e.getErrorMessage())
      }
      elem_edit_box.style.display = 'block';
    }
  </script>

  <div id="index-news">
    <div class="index-news-title-box">
      <div class="index-news-icon">
        <img src="bulletin_board_En.jpg">
      </div>
    </div>
    <!--  A的列表區  -->
    <?php
    if (!empty($InformationA)) {
      for ($i = 0; $i < mysqli_num_rows($InformationA); $i++) {
        $rs = mysqli_fetch_row($InformationA);
    ?>
        <a class="index-news-list">
          <div class="index-news-list-info-box" style="display: none">
            <h4 class="index-news-list-date"><?php echo $rs[0] ?></h4>
            <h3 class="index-news-list-title"><?php echo $rs[1] ?></h3>
            <div class="index-news-list-read">
              <span>Edit</span>
            </div>
            <div class="index-news-list-read">
              <span>Delete</span>
            </div>
          </div>
        </a>
    <?php
      }
    }
    ?>
    <a href="#" onclick="Edit(1,1,1)" class="index-news-list">
      <div class="index-news-list-info-box" style="display: none">
        <h3 class="index-news-list-title">+ 新增</h3>
      </div>
    </a>
  </div>

  <hr style="background-color: #00d6b2;width: 85%">
  <div id="index-news">
    <div class="index-news-title-box">
      <div class="index-news-icon">
        <img src="Course_slides.jpg">
        <!--            <img src="bulletin-board-ch.jpg">-->
      </div>
    </div>
    <!--  B的列表區  -->
    <?php
    if (!empty($InformationB)) {
      for ($i = 0; $i < mysqli_num_rows($InformationB); $i++) {
        $rs = mysqli_fetch_row($InformationB);
    ?>
        <?php $test = $rs[2] ?>
        <a onclick="test('<?php echo $test ?>',0)" class="index-news-list">
          <div class="index-news-list-info-box" style="display: none">
            <h4 class="index-news-list-date">Ch:<?php echo $rs[0] ?></h4>
            <h3 class="index-news-list-title"><?php echo $rs[1] ?></h3>
          </div>
        </a>
    <?php
      }
      mysqli_close($link);
    }
    ?>
    <p>+ 新增</p>

  </div>

  <hr style="background-color: #00d6b2;width: 85%">
  <div id="index-news">
    <div class="index-news-title-box">
      <div class="index-news-icon">
        <img src="Course_slides.jpg">
        <!--            <img src="bulletin-board-ch.jpg">-->
      </div>
    </div>
    <!--  C的列表區  -->
    <?php
    if (!empty($InformationC)) {
      for ($i = 0; $i < mysqli_num_rows($InformationC); $i++) {
        $rs = mysqli_fetch_row($InformationC);
    ?>
        <?php $test = $rs[2] ?>
        <a onclick="test('<?php echo $test ?>',1)" class="index-news-list">
          <div class="index-news-list-info-box" style="display: none">
            <h4 class="index-news-list-date"><?php echo $rs[0] ?></h4>
            <h3 class="index-news-list-title"><?php echo $rs[1] ?></h3>
          </div>
        </a>
    <?php
      }
      mysqli_close($link);
    }
    ?>
    <p>+ 新增</p>
  </div>

  <!--    </div>-->
  <!--</div>-->
  <!--//    $result = mysqli_query($sql) ;-->
  <!--//    while($row = mysqli_fetch_array($result)){//印出資料-->
  <!--//        echo $row['查詢的欄位1']." ";-->
  <!--//        echo $row['查詢的欄位2']."<br>";-->
  <!--//    }-->

  <!--<div id="index-news">-->
  <!--    <div class="index-news-title-box">-->
  <!--        <div class="index-news-icon">-->
  <!--            <img src="Course_slides.jpg">-->
  <!--        </div>-->
  <!--    </div>-->
  <!--    <div class="index-news-list-box">-->
  <!--        <a href="#" onclick="document.getElementById('PDF').style.display='block'" class="index-news-list">-->
  <!--            <div class="index-news-list-info-box">-->
  <!--                <h4 class="index-news-list-date">2018.12.17</h4>-->
  <!--                <h3 class="index-news-list-title"> 歡迎澳華國際法律事務所成為本所之合作夥伴，共同為客戶提供跨國性的服務！</h3>-->
  <!--                <div class="index-news-list-read"><span>READ</span></div>-->
  <!--            </div>-->
  <!--        </a>-->
  <!--        <a href="#" class="index-news-list">-->
  <!--            <div class="index-news-list-info-box">-->
  <!--                <h4 class="index-news-list-date">2018.12.10</h4>-->
  <!--                <h3 class="index-news-list-title">如何設計一個雙贏的藝術經紀合約?</h3>-->
  <!--                <div class="index-news-list-read"><span>READ</span></div>-->
  <!--            </div>-->
  <!--        </a>-->
  <!--    </div>-->
  <!--</div>-->
  <?php
  mysqli_close($link);
  ?>
  <script>
    $(document).ready(function() {
      $(".index-news-icon").fadeIn(1400);
      $(".index-news-list-info-box").fadeIn(1950);
    });
  </script>
</body>

</html>