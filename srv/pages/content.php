<!-- https://getbootstrap.com/docs/5.1/content/typography/ -->
<style>
  main {
    flex: 9;
    display: flex;
    flex-direction: row;
  }

  /* Side bar CSS */
  #select {
    flex: 2;
    display: flex;
    flex-direction: column;
  }

  select hr {
    margin: 10px auto;
    width: 70%;
    border-color: #f9830d;
  }

  #info {
    /* text-align: left; */
    flex: 1;
    margin: 10%;
    font-size: 16px;
  }

  #sub-info {
    /* text-align: left; */
    flex: 1;
    margin: 10%;
    font-size: 14px;
    text-align: center;
  }

  #option-box {
    flex: 8;
    padding-right: 10px;
    display: flex;
    flex-direction: column;
    justify-self: center;
  }

  .option-item {
    margin-top: 10px;
    text-transform: capitalize;
  }

  .sub-info {
    font-size: 14px;
    text-align: center;
    color: #ffffff;
  }

  /* Content CSS */
  content {
    flex: 8;

  }

  /* item have id */
  .content-page {
    margin: 5%;
  }

  .item-box {
    display: flex;
    flex-direction: column;
  }

  .item-title {
    flex: 2;
    font-size: 32 px;
  }

  .item-content {
    flex: 8;
    padding: 0 5%;
  }

  .bulletin-board-box {
    display: flex;
    flex-direction: row;
  }

  .bulletin-board-setting {
    flex: 1
  }

  .bulletin-board-date {
    flex: 2;
  }

  .bulletin-board-title {
    flex: 7;
    padding-left: 5%;
  }

  .file-item {
    display: flex;
    flex-direction: row;
  }

  /** Loading */
  #loading {
    display: none;
  }
</style>
<main>
  <section id="select">
    <div id="info">
      <p>辦公室: HB13</p>
      <p>分機: 20013</p>
      <p>E-mail: xxx@asia.edu.tw</p>
    </div>
    <hr>
    <div id="option-box"></div>
    <hr>
    <div id="sub-info">
      <p> 亞洲大學資工系 陳瑞奇</p>
      <p>(Rikki Chen, CSIE, Asia Univ.)</p>
    </div>

  </section>
  <content>
    <div id="loading">
      <p class="card-text placeholder-glow">
        <span class="placeholder col-7"></span>
        <span class="placeholder col-4"></span>
        <span class="placeholder col-4"></span>
        <span class="placeholder col-6"></span>
        <span class="placeholder col-8"></span>
      </p>
    </div>
    <div id="content-switch">

    </div>
  </content>
  <script>
    const option_box_elem = document.getElementById("option-box");
    const teacher_domain = "/teacher_name";
    const api_url = "http://localhost:9000/v1" + teacher_domain;

    const api_url_init = api_url + "/init"
    const api_url_content = api_url + "/course";

    const content_switch_elem = document.getElementById("content-switch");
    const br_tag = "<br />";
    const option_links = {
      "info": "/info",
      "course": "/course",
      "profile": "/profile"
    }
    var options_switch = 0;
    var options_arr = [];
    var options_last_updated_arr = [];
    var contents_arr = [];
    const COURSE = "course";
    const INFO = "info";
    const PROFILE = "profile";

    var init_response_data;

    function init() {
      get_init_api_and_create_view();
    }

    function get_init_api_and_create_view() {
      $.ajax({
          method: "GET",
          url: api_url_init,
          dataType: 'json'
        })
        .done(function(response) {
          create_init_view(response.data)
        })
        .fail(function(msg) {
          console.log(msg);
        });
    }

    function create_init_view(data) {
      // Information button
      options_arr.push(create_option_button(INFO, options_switch, data.information));
      contents_arr.push(create_information_content(data.information))
      options_last_updated_arr.push("0")
      options_switch += 1;
      // courses button
      for (let i = 0; i < data.courses.length; i++) {
        options_arr.push(create_option_button(COURSE, options_switch, data.courses[i]))
        options_last_updated_arr.push("0")
        options_switch += 1;
      }
      // todo profile button
      // options_arr.push(create_option_button(PROFILE, options_switch, option_links["info"], data.information));

      show_option_buttons();
      // first page(information content)
      show_content(0)
    }


    function create_option_button(type, options_switch, data) {

      var btn = document.createElement('button');
      var span = document.createElement('span');
      var click_even;
      btn.className = "option-item button button--anthe";
      switch (type) {
        case COURSE:
          click_even = function() {
            get_course_api(options_switch, data.id);
          }
          span.innerHTML = data.name_zh + br_tag + data.name_us;
          break;
        case INFO:
          click_even = function() {
            get_info_api(options_switch, data);
          }
          span.innerHTML = "公布欄" + br_tag + "Information";
          break;
        case PROFILE:
          // todo
          console.log('non');
          break;
        default:
          console.log('create option default???');
          break;
      }
      btn.onclick = click_even
      btn.appendChild(span);
      return btn
    }
    // todo
    function get_option_api() {
      $.ajax({
          method: "GET",
          url: api_url_content,
          dataType: 'json',
          data: {
            // "last_updated": data.last_updated
          }
        })
        .done(function(response) {
          // if (response.updated){
          //   contents_arr[options_switch] = create_content(response)
          // }

        })
        .fail(function(msg) {
          console.log(msg);
          return null
        });
    }


    function get_info_api(options_switch, id, last_updated) {
      console.log('not have get_info_api');
    }

    function get_course_api(options_switch, id) {
      $.ajax({
          method: "GET",
          url: api_url_content + "/" + id + "/" + options_last_updated_arr[options_switch],
          dataType: 'json',
        })
        .done(function(response) {
          if (response.need_update) {
            console.log('should update your course content');
            options_last_updated_arr[options_switch] = response.data.last_updated
            create_course_content(options_switch, response.data);
          }
          show_content(options_switch);
          // contents_arr
          // if (response.updated){
          //   contents_arr[options_switch] = create_content(response)
          // }

        })
        .fail(function(msg) {
          console.log(msg);
          return null
        });
    }


    function create_information_content(data) {
      var info_content = create_bulletin_board_card(data)
      return info_content
    }
    // todo
    function create_profile_content(response) {}

    function create_course_content(options_switch, data) {
      bulletin_board_card_elem = create_bulletin_board_card(data.bulletin_board)
      slide_card_elem = create_slide_card(data.slide)
      homework_card_elem = create_homework_card(data.homework)
      // todo bulletin + slide + homework
      contents_arr[options_switch] = bulletin_board_card_elem + slide_card_elem + homework_card_elem
    }


    function create_bulletin_board_card(bulletin_boards) {
      var thead = create_table_field_tags(["Date", "Information"])
      var tbody = "";
      for (let i = 0; i < bulletin_boards.length; i++) {
        tbody += tbody += create_table_field_tags([
          bulletin_boards[i].created_date,
          bulletin_boards[i].info,
        ])
      }
      var bulletin_board_card = create_card("Bulletin Board", thead, tbody)
      return bulletin_board_card
    }

    function create_slide_card(slides) {
      var thead = create_table_field_tags(["Chapter", "Type", "Title"])
      var tbody = "";
      for (let i = 0; i < slides.length; i++) {
        tbody += create_table_field_tags(
          [`Ch  ${slides[i].chapter}`,
            slides[i].file.type,
            slides[i].file.title
          ]
        )
      }
      var slide_card = create_card("Slides", thead, tbody)
      return slide_card
    }

    function create_homework_card(homeworks) {
      var thead = create_table_field_tags(["#", "Type", "Title"])
      var tbody = "";
      for (let i = 0; i < homeworks.length; i++) {
        tbody += create_table_field_tags([
          `Hw ${homeworks[i].number}`,
          homeworks[i].file.type,
          homeworks[i].file.title,
        ])
      }
      var homework_card = create_card("Homework", thead, tbody)
      return homework_card
    }

    function create_card(title, thead, tbody) {
      var card = `
      <div class="content-page">
        <div class="item-box">
          <div class="item-title h2">
            <div class="card">
              <div class="card-body">${title}</div>
            </div>
          </div>
          <div class = "item-content">
          <table class = "table table-dark table-striped">
        <thead>${thead}</thead> 
        <tbody>${tbody}</tbody>
        </table>
          </div>
        </div>
      </div>`
      return card
    }
    init();


    // ? I don't know what it is?
    // function opt_ctrl_content_display(tag_elem) {
    //   var content_switch_elem = document.getElementById(tag_elem);
    //   $(".Chose").css("display", "none");
    //   $("#" + tag_elem).fadeIn(1300);
    //   content_switch_elem.style.display = "block";
    //   // $(".index-news-icon").fadeIn(1400);
    //   // $(".index-news-list-info-box").fadeIn(1950);
    // }

    function create_table_field_tags(texts) {
      var elems = "";
      for (let i = 0; i < texts.length; i++) {
        elems += `<td>${texts[i]}</td>`;
      }
      return `<tr>${elems}</tr>`
    }

    function show_option_buttons() {
      for (let i = 0; i < options_arr.length; i++) {
        option_box_elem.appendChild(options_arr[i])
      }
    }

    function show_content(options_switch) {
      content_switch_elem.innerHTML = contents_arr[options_switch]
    }
  </script>
</main>