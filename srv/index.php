<?php
session_start();

require_once('auth/db_conn.php');
require_once('pages/head.php');
?>

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



  <?php require_once('pages/header.php'); ?>



  <div id="DIV_BODY" style="width: 100%;position: relative;padding-top: 66px;z-index: 8;background: rgb(0,0,0)">

    <?php require_once('pages/sid_bar.php'); ?>
    <?php require_once('pages/content.php'); ?>



    <script type="text/javascript">
    function test(FileName, Num, Chose) {
      var a = document.getElementById('PDF');
      var a1 = document.getElementById('iframe-pdf');
      try {
        //0是上 1下
        switch (Num) {
          case 0: {
            a1.src = "PcWeb/CourseSlides/" + Chose + "/" + FileName;
            break
          }
          case 1: {
            a1.src = "PcWeb/Homeworks/" + Chose + "/Hw" + FileName;
            break
          }
        }
      } catch (e) {
        alert("?" + e.getErrorMessage())
      }
      a.style.display = 'block';
    }
    </script>
    <script type="text/javascript">
    function Edit_BB(FileName, ID, date, content) {
      var a = document.getElementById('Edit_BB');
      document.getElementById('edit_FileName').value = FileName;
      document.getElementById('edit_ID').value = ID;
      document.getElementById('edit_date').value = date;
      document.getElementById('edit_content').value = content;
      a.style.display = 'block';
    }

    function Insert_BB(FileName) {
      var a = document.getElementById('Insert_BB');
      document.getElementById('insert_FileName').value = FileName;
      a.style.display = 'block';
    }

    function Edit_Personal(FileName, ID, date, content, partner, type, location, typeView, locationView) {
      // alert("Mass:"+FileName+ID+date+content+partner+type+location+typeView+locationView)
      var a = document.getElementById('Edit-Personal-JournalPapers');
      document.getElementById('Personal_FileName').value = FileName;
      document.getElementById('Personal_ID').value = ID;
      document.getElementById('Personal_date').value = date;
      document.getElementById('Personal_content').value = content;
      document.getElementById('Personal_type').value = type;
      document.getElementById('Personal_location').value = location;
      document.getElementById('Personal_partner').value = partner;
      // <!--=========================================================-->
      // document.getElementById('h4-type').style.display=typeView;
      // document.getElementById('h4-location').style.display=locationView;
      // <!--=========================================================-->
      // document.getElementById('edit-type').style.display=typeView;
      // document.getElementById('edit_location').style.display=locationView;
      a.style.display = 'block';
    }

    function Insert_Personal(FileName) {
      var a = document.getElementById('Insert_BB');
      document.getElementById('insert_FileName').value = FileName;
      a.style.display = 'block';
    }
    </script>
    <script>
    $(".Button").click(function() {
      $(".index-news-icon").fadeIn(1400);
      $(".index-news-list-info-box").fadeIn(1950);
      // $("#CENTER").fadeIn(1950);

    });
    </script>
    <script type="text/javascript">
    $(document).ready(function() {
      $("#bt_Exit").click(function() {
        $("#SignIn").fadeOut(1000);
      });
      $("#SignIn_i").click(function() {
        $("#SignIn").fadeIn(1200);
      });
    });
    </script>
    <script>
    window.onresize = function() {
      WebSize()
    };

    function deBUG(value) {
      try {
        // var VAR=document.getElementById("");
        alert(value);
      } catch (e) {
        alert("Error" + e.getErrorMessage());
      }
    }

    function WebSize() {
      try {
        var Width = $(window).width();
        if (Width >= 1049) {
          if (Width >= 1501) {
            ReSet_size(Width * 0.26, Width * 0.74);
          } else {
            ReSet_size(Width * 0.3, Width * 0.7);
          }
        } else {
          ReSet_size(0, Width);
        }
      } catch (e) {
        alert("Error" + e.getErrorMessage());

      }
    }

    function ReSet_size(LEFT_W, RIGHT_W) {
      var DIV_Left = document.getElementById("LEFT");
      DIV_Left.style.width = LEFT_W + "px";
      var DIV_Center = document.getElementById("CENTER");
      // var DIV_Center1 = document.getElementById("content");
      DIV_Center.style.width = RIGHT_W + "px";
      DIV_Center.style.maxWidth = RIGHT_W + "px";
      // DIV_Center1.style.width=RIGHT_W+"px";
      // DIV_Center1.style.maxWidth=RIGHT_W+"px";
      // DIV_Center.style.maxWidth=center_width+"px";
      DIV_Center.style.left = DIV_Left.style.width;
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
      var a = document.getElementById(who);
      $(".Chose").css("display", "none");
      $("#" + who).fadeIn(1300);
      a.style.display = 'block';
      // $(".index-news-icon").fadeIn(1400);
      // $(".index-news-list-info-box").fadeIn(1950);
    }
    </script>


    <!--BB編輯-->
    <div id="Edit_BB" class="modal" style="z-index: 99">
      <div class="modal-content animate">
        <div style="height: 90%;position: relative;">
          <span style="position: absolute;right: 2.5%;top: 30px"
            onclick="document.getElementById('Edit_BB').style.display='none'" class="close" title="Close Modal">X</span>
          <div id="Title" style="position: absolute;"></div>
          <div id="InputContent" style="position: relative;">
            <h1>Edit</h1>
            <form action="PcWeb/Login_v16/Login_v16/UpData.php" method="GET" enctype="multipart/form-data"
              style="position: absolute;">
              <input id="edit_FileName" type="hidden" name="edit_FileName" ">
                        <input id=" edit_ID" type="hidden" name="edit_ID" ">
                        <h4 style=" color: orange;">Date:</h4>
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
          <span style="position: absolute;right: 2.5%;top: 30px"
            onclick="document.getElementById('Insert_BB').style.display='none'" class="close" title="Close Modal">X
          </span>
          <div id="Title" style="position: absolute;">
          </div>
          <div id="InputContent" style="margin: 0 auto">
            <form action="PcWeb/Login_v16/Login_v16/InsertData.php" method="GET" enctype="multipart/form-data"
              style="position: absolute;">
              <input id="insert_FileName" type="hidden" name="insert_FileName"">
                        <h4 style=" color: orange;">Date:</h4>
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
          <span style="position: absolute;right: 2.5%;top: 30px"
            onclick="document.getElementById('Edit-Personal-JournalPapers').style.display='none'" class="close"
            title="Close Modal">X</span>
          <div style="position: relative;">
            <h1 style="color: white;">Edit</h1>
            <form action="PcWeb/Login_v16/Login_v16/UpData.php" method="GET" style="position: absolute;">
              <input id="Personal_FileName" type="hidden" name="edit_FileName" ">
                        <input id=" Personal_ID" type="hidden" name="edit_ID" ">
                        <h4 style=" color: orange;">Date:</h4>
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
              <button onclick="document.getElementById('Edit-Personal-JournalPapers').style.display='none'"
                style="color: white;">退出</button>

            </form>
          </div>
        </div>
      </div>
    </div>

    <!--Personal新增-->
    <div id="Insert_Personal" class="modal" style="z-index: 99">
      <div class="modal-content animate">
        <div style="height: 90%;position: relative;">
          <span style="position: absolute;right: 2.5%;top: 30px"
            onclick="document.getElementById('Insert_BB').style.display='none'" class="close" title="Close Modal">X
          </span>
          <div id="Title" style="position: absolute;">
          </div>
          <div id="InputContent" style="margin: 0 auto">
            <form action="PcWeb/Login_v16/Login_v16/InsertData.php" method="GET" enctype="multipart/form-data"
              style="position: absolute;">
              <input id="insert_FileName" type="hidden" name="insert_FileName"">
                        <h4 style=" color: orange;">Date:</h4>
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


    <?php
        require_once('pages/login.php');
        require_once('pages/pdf.php');
        mysqli_close($link);
        ?>

</body>

</html>