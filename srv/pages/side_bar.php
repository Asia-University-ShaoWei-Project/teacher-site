<?php
$mock_course = array(
  array(
    "tag" => "info",
    "title-us" => "Information",
    "title-zh" => "公告"
  ),
  array(
    "tag" => "CN",
    "title-us" => "Computer Network",
    "title-zh" => "計算機網路"
  ),
  array(
    "tag" => "MWT",
    "title-us" => "Multimedia Web Tech",
    "title-zh" => "多媒體網路技術"
  ),
  array(
    "tag" => "pub",
    "title-us" => "Publisher",
    "title-zh" => "學術"
  ),
);
?>
<section>
  <div>
    <img id="divImage2" src="static/img/avatar.jpg" style="border-radius: 2%">
  </div>
  <hr>

  <div style="margin-left: 20%;text-align: left;font-size: 18px;">
    <p style="color: white">學歷： 國立中興大學資訊科學博士</p>
    <p style="color: white">辦公室： HB13</p>
    <p style="color: white">分機： 20013</p>
    <p style="color: white">E-mail： rikki@asia.edu.tw</p>
  </div>
  <hr>
  <table border="0">
    <?php
    foreach ($mock_course as $mock) { ?>
    <tr>
      <td class="side-bar-option" onclick="opt_ctrl_frame_display(<?php echo $mock['tag'] ?>)">
        <span>
          <?php echo $mock['title-zh'] . "<br />" . $mock['title-us'] ?>
        </span>
      </td>
    </tr>
    <?php

    } ?>
  </table>
  <br>
  <hr>
  <p class="BottomText">任何建議請寄: EMAIL rikki@asia.edu.tw</p>
  <p class="BottomText">亞洲大學資工系 陳瑞奇(Rikki Chen, CSIE, Asia Univ.)</p>
  <p class="BottomText">感謝您！</p>
</section>
<script>
function opt_ctrl_content_display(tag_elem) {
  var content_elem = document.getElementById(tag_elem);
  $(".Chose").css("display", "none");
  $("#" + tag_elem).fadeIn(1300);
  content_elem.style.display = "block";
  // $(".index-news-icon").fadeIn(1400);
  // $(".index-news-list-info-box").fadeIn(1950);
}
</script>