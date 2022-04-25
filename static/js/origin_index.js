var form_modal = document.getElementById("form-modal-bulletin");
form_modal.addEventListener("show.bs.modal", function (event) {
  var btn = event.relatedTarget;
  // todo: add attribute in generate html function
  var date = btn.getAttribute("data-bs-date");
  var info = btn.getAttribute("data-bs-info");
  // If necessary, you could initiate an AJAX request here
  // and then do the updating in a callback.
  //
  // Update the modal's content.
  var modal_title = form_modal.querySelector(".modal-title");
  // var modal_input_date = form_modal.querySelector('.modal-body #bulletin-date')

  var modal_input_date = form_modal.querySelector("#bulletin-date");
  // var modal_input_info = form_modal.querySelector('.modal-body #bulletin-info')
  var modal_input_info = form_modal.querySelector("#bulletin-info");
  modal_title.textContent = "New message to ";
  modal_input_date.value = date;
  modal_input_info.value = info;
  var modal_input_submit_btn = form_modal.querySelector("#submit");
  modal_input_submit_btn.addEventListener("click", function name() {
    var url = api.url + "/test/1";
    var data = {
      create_date: modal_input_date.value,
      info: modal_input_info.value,
    };
    console.log(data);

    testSendForm(url, data);
  });
});

function testSendForm(url, params) {
  axios
    .delete(url, params)
    .then((res) => {
      console.log(res);
    })
    .catch((err) => {
      console.error(err);
    });
}

// url

// todo: use axios instead of ajax
// https://github.com/axios/axios
// const axios = require('axios').default;

const option_box_elem = document.getElementById("option-box");
const content_switch_elem = document.getElementById("content-switch");
const br_tag = "<br>";

class API {
  constructor(origin, version, teacherDomain, resources) {
    this.origin = origin;
    this.version = version;
    this.teacherDomain = teacherDomain;
    this.resources = resources;
    this.url = this.origin + "/api/" + this.version + "/" + this.teacherDomain;
  }
  getResourceUrl(type) {
    return this.url + "/" + this.resources[type];
  }
  getCourseUrl(id, lastUpdateTime) {
    return this.resources.course + "/" + id + "/" + lastUpdateTime;
  }
}
const COURSE = "course";
const INFO = "info";
const INIT = "init";

var api = new API(
  // origin(e.g. http://domain)
  window.location.origin,
  // version
  "v1",
  // teacher domain
  window.location.pathname.replace("/", ""),
  // resource name
  {
    init: INIT,
    info: INFO,
    course: COURSE,
  }
);
var options = [];
// options[index].showContent() in button click event
class Item {
  constructor(apiUrl, nameZh, nameUs, bulletin, slide, homework, lastUpdated) {
    this.apiUrl = apiUrl;
    this.id = id;
    this.nameZh = nameZh;
    this.nameUs = nameUs;
    this.bulletin = bulletin;
    this.slide = slide;
    this.homework = homework;
    this.lastUpdated = lastUpdated;
    this.content = "";
  }
  // todo: info or course
  createOptionButton() {
    return;
  }
  showContent(rebuild) {
    let _content;
    if (rebuild) {
      _content = createContent(this.bulletin, this.slide, this.homework);
    } else {
      _content = this.content;
    }
    content_switch_elem.innerHTML = _content;
  }
  getData() {
    // apiUrl = [ api.resources.info | api.resources.course]
    let url = this.apiUrl + "/" + id + "/" + lastUpdateTime;
    axios
      .get(url)
      .then((res) => {
        let rebuild = true;
        switch (res.status) {
          // the information is up to date
          case HTTP_STATUS_CODE.noContent:
            rebuild = false;
            alert("the data is up to date!");
            break;
          // need to update information
          case HTTP_STATUS_CODE.ok:
            this.lastUpdated = res.data.data.last_updated;
            this.bulletin = newBulletin(res.data.data.bulletin_board);
            // todo
            // this.slide = newSlide()
            // this.homework = newHomework()
            break;
          // todo:
          default:
            break;
        }
        this.showContent(rebuild);
      })
      .catch((err) => {
        switch (err.response.status) {
          case HTTP_STATUS_CODE.badRequest:
            console.error("bad request");
            break;
          default:
            console.error("error status code:", err.response.status);
            break;
        }
      });
  }
}
class BulletinBoardRow {
  constructor(id, date, info) {
    this.id = id;
    this.date = date;
    this.info = info;
  }
}
class Table {
  constructor(fieldsTitle, id, date, info) {
    this.fieldsTitle = fieldsTitle;
    this.rows = [];
  }
  setRows(rows) {
    this.rows = rows;
  }
}
// class Slide {
//   constructor(titleFields) {
//     this.fieldsTitle = fieldsTitle;
//   }
// }
// class Homework {
//   constructor(titleFields) {
//     this.fieldsTitle = fieldsTitle;
//   }
// }
var bfieldsTitle = ["Date", "Information"];
var sfieldsTitle = ["Chapter", "Type", "Title"];
var hfieldsTitle = ["#", "Type", "Title"];

function init() {
  console.log("init");
  getInitApiAndCreateView(api.getResourceUrl(INIT));
}
function getInitApiAndCreateView(url) {
  axios
    .get(url)
    .then((res) => {
      console.log("getInitApiAndCreateView success");
      isTeacher(res.data.auth);
      create_init_view(res.data.data);
    })
    .catch((err) => {
      console.error("getInitApiAndCreateView:", err);
    });
}

function isTeacher(valid) {
  if (valid) {
    // addEditField();
  }
}
// function addEditField() {
//   console.log("addEditField");

//   bulletinTitleFields.push("Edit");
//   slideTitleFields.push("Edit");
//   homework_fields.push("Edit");
// }

const ZERO_TIME = "0";
var contents = [];
var items = [];
var lastUpdatedOfCourse = [];
var optionSwitchIndex = 0;

// Courses      []Courses      `json:"courses"`
// Informations []Informations `json:"information"

function newItem(data, bulletin, slide, hw) {
  // nameZh: any, nameUs: any, bulletin: any, slide: any, homework: any)
  return new Item(data.name_zh, data.name_us, bulletin, slide, hw);
}
function newBulletin(data) {
  // todo: change bfieldsTitle
  let bulletinBoard = new Table(bfieldsTitle);
  let rows = [];
  for (let i = 0; i < data.length; i++) {
    rows.push(
      new BulletinBoardRow(data[i].id, data[i].created_date, data[i].info)
    );
  }
  bulletinBoard.setRows(rows);
  return bulletinBoard;
}

function create_init_view(data) {
  console.log("create_init_view");
  //* Information button
  let bulletin = newBulletin(data.information);
  items.push(newItem(data.information, bulletin, null, null));
  // todo: create button
  // contents.push(create_information_content(data.information));
  // lastUpdatedOfCourse.push(ZERO_TIME);
  // optionSwitchIndex += 1;
  //* course buttons
  for (let i = 0; i < data.courses.length; i++) {
    items.push(newItem(data.course[i], null, null, null));

    // items.push(createOptionButton(COURSE, optionSwitchIndex, data.courses[i]));
    // lastUpdatedOfCourse.push(ZERO_TIME);
    // optionSwitchIndex += 1;
  }
  showOptionButtons();
  // first page(information content)
  showContent(0);
}

// function createOptionButton(type, index, data) {
function createOptionButton(nameZh, nameUs, event) {
  console.log("createOptionButton");
  let btn = document.createElement("button");
  let span = document.createElement("span");
  let clickMethod;
  let text;
  btn.className = "option-item button button--anthe";

  switch (type) {
    case COURSE:
      clickMethod = function () {
        getCourseApi(index, data.id);
      };
      text = data.name_zh + br_tag + data.name_us;
      break;
    case INFO:
      clickMethod = function () {
        getInfoApi(index, data);
      };
      text = "公布欄" + br_tag + "Information";
      break;
    default:
      console.log("create option default???");
      break;
  }
  span.innerHTML = text;
  btn.onclick = clickMethod;
  btn.appendChild(span);
  return btn;
}
// todo:
function getInfoApi(index, id, last_updated) {
  console.log("not have getInfoApi");
}
// ++++++++++++

function getCourseApi(index, id) {
  let url = api.getCourseUrl(id, lastUpdatedOfCourse[index]);
  axios
    .get(url)
    .then((res) => {
      switch (res.status) {
        // the information is up to date
        case HTTP_STATUS_CODE.noContent:
          break;
        // need to update information
        case HTTP_STATUS_CODE.ok:
          lastUpdatedOfCourse[index] = res.data.data.last_updated;
          createContent(index, res.data.data);
          break;
        // todo:
        default:
          break;
      }
      showContent(index);
    })
    .catch((err) => {
      switch (err.response.status) {
        case HTTP_STATUS_CODE.badRequest:
          console.error("bad request");
          break;
        default:
          console.error("error status code:", err.response.status);
          break;
      }
    });
}

function create_information_content(data) {
  console.log("create_information_content");
  // var info_content = create_bulletin_board_card(data);
  let info_content = create_item(BULLETIN, data);
  return info_content;
}
// todo

function createContent(index, data) {
  console.log("createContent");

  let bulletin_content = create_item(BULLETIN, data.bulletin_board);
  let slide_content = create_item(SLIDE, data.slide);
  let homework_content = create_item(HOMEWORK, data.homework);
  // bulletin_board_card_elem = create_bulletin_board_card(
  //   data.bulletin_board
  // );
  // slide_card_elem = create_slide_card(data.slide);
  // homework_card_elem = create_homework_card(data.homework);
  // todo bulletin + slide + homework
  contents[index] = bulletin_content + slide_content + homework_content;
}

const BULLETIN = "bulletin";
const SLIDE = "slide";
const HOMEWORK = "hw";
const card_titles = {
  bulletin: "Bulletin Board",
  slide: "Slides",
  hw: "Homework",
};
const table_head_title = {
  bulletin: ["Date", "Information"],
  slide: ["Chapter", "Type", "Title"],
  hw: ["#", "Type", "Title"],
};

function create_item(type, data) {
  console.log("create_item");

  let body_contents = null;
  let _data = null;
  let thead = create_table_elem(table_head_title[type]);
  let tbody = "";
  for (let i = 0; i < data.length; i++) {
    _data = data[i];
    switch (type) {
      case BULLETIN:
        body_contents = [_data.created_date, _data.info];
        break;
      case SLIDE:
        body_contents = [
          "Ch" + _data.chapter,
          _data.file.type,
          _data.file.title,
        ];
        break;
      case HOMEWORK:
        body_contents = [
          "Hw" + _data.number,
          _data.file.type,
          _data.file.title,
        ];
        break;
      default:
        body_contents = [];
        break;
    }
    tbody += create_table_elem(body_contents);
  }
  return create_card(card_titles[type], thead, tbody);
}

function create_card(title, thead, tbody) {
  console.log("create_card");

  let card =
    `
      <div class="content-page">
        <div class="item-box">
          <div class="item-title h2">
            <div class="card">
              <div class="card-body">` +
    title +
    `</div>
            </div>
          </div>
          <div class = "item-content">
            <table class = "table table-dark table-striped">
              <thead>` +
    thead +
    `</thead> 
              <tbody>` +
    tbody +
    `</tbody>
            </table>
          </div>
        </div>
      </div>`;
  return card;
}
function create_table_elem(texts) {
  console.log("create_table_elem");

  let elems = "";
  for (let i = 0; i < texts.length; i++) {
    elems += "<td>" + texts[i] + "</td>";
  }
  return "<tr>" + elems + "</tr>";
}

function showOptionButtons() {
  console.log("showOptionButtons");
  for (let i = 0; i < items.length; i++) {
    elem = items[i].createOptionButton();
    option_box_elem.appendChild(elem);
  }
}

function showContent(index) {
  console.log("show content");

  content_switch_elem.innerHTML = contents[index];
}

const HTTP_STATUS_CODE = {
  ok: 200,
  // StatusCreated              = 201,
  // StatusAccepted             = 202,
  // StatusNonAuthoritativeInfo = 203,
  noContent: 204,
  // StatusResetContent         = 205,
  badRequest: 400,
  unauthorized: 401,
  notFound: 404,
};
