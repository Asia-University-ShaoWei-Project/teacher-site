<?php
session_start();

require_once('auth/db_conn.php');
require_once('pages/head.php');
?>

<body onload="WebSize()">
  <!------- Login DIV ------->
  <?php require_once('pages/login.php'); ?>
  <header id="main-header"></header>
  <section id="main-section"></section>
  <article id="main-article"></article>
  <aside id="main-aside"></aside>
  <footer id="main-footer"></footer>

  <?php
    require_once('pages/header.php'); ?>
  <main>
    <?php
        require_once('pages/side_bar.php');
        require_once('pages/content.php');
        ?>
  </main>

  <!-- 
  <main>
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
              <input id="edit_FileName" type="hidden" name="edit_FileName">
              <input id=" edit_ID" type="hidden" name="edit_ID">
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
              <input id="insert_FileName" type="hidden" name="insert_FileName">
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
              <input id="insert_FileName" type="hidden" name="insert_FileName">
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
 -->
</body>
<?php
include_once('pages/login.php');
include_once('pages/pdf.php');
$db->close();

?>

</html>