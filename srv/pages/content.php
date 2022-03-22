<div id="CENTER" style="position: absolute">
  <div id="content" style="background-color: white">
    <!--  公告 列表區  -->
    <?php
    $Word = array(
      "MessageAnnouncement", "PersonalInformation", "ComputerNetworks", "ComputerOrganization"
    );
    $Chose = array(
      "Bulletin_board", "JournalPapers", "ComputerNetworks", "ComputerOrganization"
    );
    for ($i = 0; $i < 4; $i++) { ?>
    <div id="<?php echo $Word[$i] ?>" class="Chose">
      <div id="index-news">
        <div class="index-news-title-box">
          <div class="index-news-icon">
            <img src="PcWeb/image/<?php echo $Chose[$i] ?>.jpg">
          </div>
        </div>
        <?php
          switch ($i) {
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
          for ($i = 0; $i < mysqli_num_rows($Information); $i++) {
            $rs = mysqli_fetch_row($Information);
          ?>
        <a class="index-news-list">
          <div class="index-news-list-info-box" style="display: none">
            <?php if ($key == 0) { ?>
            <h4 class="index-news-list-date-edit"
              onclick="Edit_BB('BulletinBoard','<?php echo $rs[0] ?>','<?php echo $rs[1] ?>','<?php echo $rs[2] ?>')">
              編輯</h4>
            <?php } ?>
            <h4 class="index-news-list-date"><?php echo $rs[1] ?></h4>
            <h3 class="index-news-list-title" style="width: 75%"><?php echo $rs[2] ?></h3>
            <?php if ($key == 0) { ?>
            <div class="index-news-list-date-edit">
              <form action="PcWeb/Login_v16/Login_v16/DeleteData.php" method="get">
                <input type="hidden" name="DeleteFileName" value="BulletinBoard">
                <input type="hidden" name="DeleteID" value='<?php echo $rs[0] ?>'>
                <button type="submit"><i class="fa fa-trash" style="color: black;"></i></button>
              </form>
            </div>
            <?php } ?>
          </div>
        </a>
        <?php
          }
          ?>
        <?php
          if ($key == 0) {
          ?>
        <a href="#" onclick="Insert_BB('BulletinBoard')" class="index-news-list">
          <div class="index-news-list-info-box" style="display: none">
            <h3 class="index-news-list-add">新增(+)</h3>
          </div>
        </a>
        <?php } ?>
      </div>
    </div>


    <?php
      if ($i > 0) {
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
        for ($i = 0; $i < mysqli_num_rows($Information); $i++) {
          $rs = mysqli_fetch_row($Information);
        ?>
        <a class="index-news-list">
          <div class="index-news-list-info-box" style="display: none">
            <?php if ($key == 0) { ?>
            <h4 class="index-news-list-date-edit"
              onclick="Edit_Personal(1,'<?php echo $rs[0] ?>','<?php echo $rs[1] ?>','<?php echo $rs[2] ?>','<?php echo $rs[3] ?>','<?php echo $rs[4] ?>','','block','none')">
              編輯</h4>
            <?php } ?>
            <h4 class="index-news-list-date"><?php echo $rs[1] ?></h4>
            <h3 class="index-news-list-personal" style="width: 75%"><?php echo $rs[2] ?></h3>
            <h4 class="index-news-list-date"><?php echo $rs[3] ?></h4>
            <h4 class="index-news-list-date"><?php echo $rs[4] ?></h4>
            <?php if ($key == 0) { ?>
            <div class="index-news-list-date-edit">
              <form action="PcWeb/Login_v16/Login_v16/DeleteData.php" method="get">
                <input type="hidden" name="DeleteFileName" value="JournalPapers">
                <input type="hidden" name="DeleteID" value='<?php echo $rs[0] ?>'>
                <button type="submit"><i class="fa fa-trash" style="color: black;"></i></button>
              </form>
            </div>
            <?php } ?>
          </div>
        </a>
        <?php
        } ?>
        <?php
        if ($key == 0) {
        ?>
        <a href="#" onclick="Insert_BB('JournalPapers')" class="index-news-list">
          <div class="index-news-list-info-box" style="display: none">
            <h3 class="index-news-list-add">新增(+)</h3>
          </div>
        </a>
        <?php } ?>
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
        for ($i = 0; $i < mysqli_num_rows($Information); $i++) {
          $rs = mysqli_fetch_row($Information);
        ?>
        <a class="index-news-list">
          <div class="index-news-list-info-box" style="display: none">
            <?php if ($key == 0) { ?>
            <h4 class="index-news-list-date-edit"
              onclick="Edit_Personal(2,'<?php echo $rs[0] ?>','<?php echo $rs[1] ?>','<?php echo $rs[2] ?>','<?php echo $rs[4] ?>','','<?php echo $rs[3] ?>','none','block')">
              編輯</h4>
            <?php } ?>
            <h4 class="index-news-list-date"><?php echo $rs[1] ?></h4>
            <h3 class="index-news-list-personal" style="width: 75%"><?php echo $rs[2] ?></h3>
            <h4 class="index-news-list-date"><?php echo $rs[3] ?></h4>
            <h4 class="index-news-list-date"><?php echo $rs[4] ?></h4>
            <?php if ($key == 0) { ?>
            <div class="index-news-list-date-edit">
              <form action="PcWeb/Login_v16/Login_v16/DeleteData.php" method="get">
                <input type="hidden" name="DeleteFileName" value="ConferencePapers">
                <input type="hidden" name="DeleteID" value='<?php echo $rs[0] ?>'>
                <button type="submit"><i class="fa fa-trash" style="color: black;"></i></button>
              </form>
            </div>
            <?php } ?>
          </div>
        </a>
        <?php
        } ?>
        <?php
        if ($key == 0) {
        ?>
        <a href="#" onclick="Insert_BB('ConferencePapers')" class="index-news-list">
          <div class="index-news-list-info-box" style="display: none">
            <h3 class="index-news-list-add">新增(+)</h3>
          </div>
        </a>
        <?php } ?>
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
        for ($i = 0; $i < mysqli_num_rows($Information); $i++) {
          $rs = mysqli_fetch_row($Information);
        ?>
        <a class="index-news-list">
          <div class="index-news-list-info-box" style="display: none">
            <?php if ($key == 0) { ?>
            <h4 class="index-news-list-date-edit"
              onclick="Edit_Personal(3,'<?php echo $rs[0] ?>','<?php echo $rs[1] ?>','<?php echo $rs[2] ?>','<?php echo $rs[3] ?>','','','none','none')">
              編輯</h4>
            <?php } ?>
            <h4 class="index-news-list-date"><?php echo $rs[1] ?></h4>
            <h3 class="index-news-list-personal" style="max-width: 800px"><?php echo $rs[2] ?></h3>
            <h4 class="index-news-list-date"><?php echo $rs[3] ?></h4>
            <?php if ($key == 0) { ?>
            <div class="index-news-list-date-edit">
              <form action="PcWeb/Login_v16/Login_v16/DeleteData.php" method="get">
                <input type="hidden" name="DeleteFileName" value="ResearchGrant">
                <input type="hidden" name="DeleteID" value='<?php echo $rs[0] ?>'>
                <button type="submit"><i class="fa fa-trash" style="color: black;"></i></button>
              </form>
            </div>
            <?php } ?>
          </div>
        </a>
        <?php
        } ?>
        <?php
        if ($key == 0) {
        ?>
        <a href="#" onclick="Insert_BB('ResearchGrant')" class="index-news-list">
          <div class="index-news-list-info-box" style="display: none">
            <h3 class="index-news-list-add">新增(+)</h3>
          </div>
        </a>
        <?php } ?>
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
        for ($i = 0; $i < mysqli_num_rows($Information); $i++) {
          $rs = mysqli_fetch_row($Information);
        ?>
        <a class="index-news-list">
          <div class="index-news-list-info-box" style="display: none">
            <?php if ($key == 0) { ?>
            <h4 class="index-news-list-date-edit"
              onclick="Edit_BB('IntroductionToComputerNetworks','<?php echo $rs[0] ?>','<?php echo $rs[1] ?>','<?php echo $rs[2] ?>')">
              編輯</h4>
            <?php } ?>
            <h4 class="index-news-list-date"><?php echo $rs[1] ?></h4>
            <h3 class="index-news-list-title" style="width: 75%"><?php echo $rs[2] ?></h3>
            <?php if ($key == 0) { ?>
            <div class="index-news-list-date-edit">
              <form action="PcWeb/Login_v16/Login_v16/DeleteData.php" method="get">
                <input type="hidden" name="DeleteFileName" value="IntroductionToComputerNetworks">
                <input type="hidden" name="DeleteID" value='<?php echo $rs[0] ?>'>
                <button type="submit"><i class="fa fa-trash" style="color: black;"></i></button>
              </form>
            </div>
            <?php } ?>
          </div>
        </a>
        <?php
        } ?>
        <?php
        if ($key == 0) {
        ?>
        <a href="#" onclick="Insert_BB('IntroductionToComputerNetworks')" class="index-news-list">
          <div class="index-news-list-info-box" style="display: none">
            <h3 class="index-news-list-add">新增(+)</h3>
          </div>
        </a>
        <?php } ?>
      </div>
      <hr style="background-color: #f9830d;width: 85%;margin: 0 auto">

      <div id="index-news">
        <div class="index-news-title-box">
          <div class="index-news-icon">
            <img src="PcWeb/image/course_slides.jpg">
          </div>
        </div>

        <?php
        for ($i = 0; $i < mysqli_num_rows($Course_slides); $i++) {
          $rs = mysqli_fetch_row($Course_slides);
        ?>
        <?php $test = $rs[2] ?>
        <a onclick="test('<?php echo $test ?>',0,'Computerworks')" class="index-news-list">
          <div class="index-news-list-info-box" style="display: none">
            <h4 class="index-news-list-date">Ch:<?php echo $rs[0] ?></h4>
            <h3 class="index-news-list-title"><?php echo $rs[1] ?></h3>
            <div class="index-news-list-read">
              <span>READ</span>
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
        for ($i = 0; $i < mysqli_num_rows($Homeworks); $i++) {
          $rs = mysqli_fetch_row($Homeworks);
        ?>
        <?php $test = $rs[2] ?>
        <a onclick="test('<?php echo $test ?>',1,'Computerworks')" class="index-news-list">
          <div class="index-news-list-info-box" style="display: none">
            <h4 class="index-news-list-date"><?php echo $rs[0] ?></h4>
            <h3 class="index-news-list-title"><?php echo $rs[1] ?></h3>
            <div class="index-news-list-read">
              <span>READ</span>
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
        for ($i = 0; $i < mysqli_num_rows($Information); $i++) {
          $rs = mysqli_fetch_row($Information);
        ?>
        <a class="index-news-list">
          <div class="index-news-list-info-box" style="display: none">
            <?php if ($key == 0) { ?>
            <h4 class="index-news-list-date-edit"
              onclick="Edit_BB('ComputerOrganization','<?php echo $rs[0] ?>','<?php echo $rs[1] ?>','<?php echo $rs[2] ?>')">
              編輯</h4>
            <?php } ?>
            <h4 class="index-news-list-date"><?php echo $rs[1] ?></h4>
            <h3 class="index-news-list-title" style="width: 75%"><?php echo $rs[2] ?></h3>
            <?php if ($key == 0) { ?>
            <div class="index-news-list-date-edit">
              <form action="PcWeb/Login_v16/Login_v16/DeleteData.php" method="get">
                <input type="hidden" name="DeleteFileName" value="ComputerOrganization">
                <input type="hidden" name="DeleteID" value='<?php echo $rs[0] ?>'>
                <button type="submit"><i class="fa fa-trash" style="color: black;"></i></button>
              </form>
            </div>
            <?php } ?>
          </div>
        </a>
        <?php
        } ?>
        <?php
        if ($key == 0) {
        ?>
        <a href="#" onclick="Insert_BB('ComputerOrganization')" class="index-news-list">
          <div class="index-news-list-info-box" style="display: none">
            <h3 class="index-news-list-add">新增(+)</h3>
          </div>
        </a>
        <?php } ?>
      </div>
      <hr style="background-color: #f9830d;width: 85%;margin: 0 auto">

      <div id="index-news">
        <div class="index-news-title-box">
          <div class="index-news-icon">
            <img src="PcWeb/image/course_slides.jpg">
          </div>
        </div>
        <?php
        for ($i = 0; $i < mysqli_num_rows($Course_slides); $i++) {
          $rs = mysqli_fetch_row($Course_slides);
        ?>
        <?php $test = $rs[2] ?>
        <a onclick="test('<?php echo $test ?>',0,'ComputerOrganization')" class="index-news-list">
          <div class="index-news-list-info-box" style="display: none">
            <h4 class="index-news-list-date">Ch:<?php echo $rs[0] ?></h4>
            <h3 class="index-news-list-title"><?php echo $rs[1] ?></h3>
            <div class="index-news-list-read">
              <span>READ</span>
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
        for ($i = 0; $i < mysqli_num_rows($Homeworks); $i++) {
          $rs = mysqli_fetch_row($Homeworks);
        ?>
        <?php $test = $rs[2] ?>
        <a onclick="test('<?php echo $test ?>',1,'ComputerOrganization')" class="index-news-list">
          <div class="index-news-list-info-box" style="display: none">
            <h4 class="index-news-list-date"><?php echo $rs[0] ?></h4>
            <h3 class="index-news-list-title"><?php echo $rs[1] ?></h3>
            <div class="index-news-list-read">
              <span>READ</span>
            </div>
          </div>
        </a>
        <?php
        }
        ?>
      </div>
    </div>
    <div id="A1" class="Chose" style="display: none;">
      <h1>還沒做sorry</h1>
    </div>
    <div id="A2" class="Chose" style="display: none;">
      <h1>還沒做sorry</h1>
    </div>
  </div>
</div>