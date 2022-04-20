const br_tag = "<br>";
// var option_box_elem = document.getElementById("option-box");
// const content_switch_elem = ;
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

  let info_get_url = `/info/bulletin`;

  getInfoApi(info_get_url);
  // getCourseApi(api.getResourceUrl());

  // getInitApiAndCreateView(api.getResourceUrl(INIT));
}
function getInfoApi(url) {
  // [ bulletins<list>(id, date, content), id(info), last_modified(info) ]
  console.log("getInfoApi");
  axios
    .get(api.url + url, {
      params: {
        last_modified: "0",
      },
    })
    .then((res) => {
      console.log("getInfoApi api success");
      let data = res.data.data;
      console.log(data);
      console.log(data.id);
      console.log(data.last_modified);
      // isTeacher(res.data.auth);
      // todo: create option for information button
      // todo: show the information view
      let infoItem = newItem(null, {
        id: "0",
        name_zh: "公布欄",
        name_us: "Information",
      });
      let bulletin = newBulletin(data.bulletins);
      infoItem.setID(data.id);
      infoItem.setLastModified(data.last_modified);
      infoItem.setBulletin(bulletin);
      infoItem.setContent(createContent(bulletin));
      items.push(infoItem);
      showContent(items[0].content);
      // todo: temp
      showOptionButtons();
    })
    .catch((err) => {
      console.error("getInfoApi:", err);
    });
}
// todo
// function getCourseApi(url) {
//   // [ bulletins<list>(id, date, content), id(info), last_modified(info) ]
//   console.log("getCourseApi");
//   axios
//     .get(url)
//     .then((res) => {
//       console.log("getCourseApi api success");
//       console.log(res.data);
//       // isTeacher(res.data.auth);
//       create_init_view(res.data.data);
//     })
//     .catch((err) => {
//       console.error("getInitApiAndCreateView:", err);
//     });
// }
// todo:delete
// function getInitApiAndCreateView(url) {
//   axios
//     .get(url)
//     .then((res) => {
//       console.log("getInitApiAndCreateView success");
//       console.log(res.data);
//       // isTeacher(res.data.auth);
//       create_init_view(res.data.data);
//     })
//     .catch((err) => {
//       console.error("getInitApiAndCreateView:", err);
//     });
// }

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

function newItem(apiUrl, data, bulletin, slide, hw, lastUpdated) {
  // nameZh: any, nameUs: any, bulletin: any, slide: any, homework: any)
  return new Item(
    apiUrl,
    data.id,
    data.name_zh,
    data.name_us,
    bulletin,
    slide,
    hw,
    lastUpdated
  );
}

function newBulletin(data) {
  // todo: change bfieldsTitle
  let title = "Bulletin Board";
  let bulletinBoard = new Table(title, bfieldsTitle);
  let rows = [];
  for (let i = 0; i < data.length; i++) {
    rows.push(new BulletinBoardRow(data[i].id, data[i].date, data[i].content));
  }
  bulletinBoard.setRows(rows);
  return bulletinBoard;
}

function create_init_view(data) {
  console.log("create_init_view");
  //* Information()
  let bulletin = newBulletin(data.information);
  items.push(
    newItem(
      api.getResourceUrl(INFO),
      data.information,
      bulletin,
      null,
      null,
      data.info_last_updated
    )
  );
  // todo: create button
  // contents.push(create_information_content(data.information));
  // lastUpdatedOfCourse.push(ZERO_TIME);
  // optionSwitchIndex += 1;
  //* Course(#The b, s and h is undefined)
  for (let i = 0; i < data.courses.length; i++) {
    items.push(newItem(api.getResourceUrl(COURSE), data.courses[i]));
    // items.push(createOptionButton(COURSE, optionSwitchIndex, data.courses[i]));
    // lastUpdatedOfCourse.push(ZERO_TIME);
    // optionSwitchIndex += 1;
  }
  showOptionButtons();
  showContent(items[0].getContent(true));
}
// *option

function showOptionButtons() {
  let elem;
  for (let i = 0; i < items.length; i++) {
    elem = createOptionButton(items[i]);
    document.getElementById("option-box").appendChild(elem);
  }
}
function createOptionButton(item) {
  console.log("createOptionButton");
  let btn = document.createElement("button");
  let span = document.createElement("span");
  let text = item.nameZh + br_tag + item.nameUs;
  btn.className = "option-item button button--anthe";
  span.innerHTML = text;
  btn.onclick = () => item.getData();
  btn.appendChild(span);

  return btn;
}
//* content
function showContent(content) {
  document.getElementById("content-switch").innerHTML = content;
}
function createContent(bulletin, slide, homework) {
  console.log("createContent");
  let content = "";

  if (bulletin != null && bulletin != undefined) {
    content += create_item(bulletin);
  }
  if (slide != null && slide != undefined) {
    content += create_item(slide);
  }
  if (homework != null && homework != undefined) {
    content += create_item(homework);
  }
  return content;
}
function create_item(table) {
  console.log("create_item");

  let body_contents;
  let thead = create_table_elem(table.getFieldsTitle());
  let tbody = "";
  for (let i = 0; i < table.getRowsLen(); i++) {
    body_contents = table.getRow(i).getDataList();
    tbody += create_table_elem(body_contents);
  }
  return create_card(table.getTitle(), thead, tbody);
}

function create_table_elem(body_contents) {
  console.log("create_table_elem");
  let elems = "";
  for (let i = 0; i < body_contents.length; i++) {
    elems += "<td>" + body_contents[i] + "</td>";
  }
  return "<tr>" + elems + "</tr>";
}

function create_card(title, thead, tbody) {
  let card = `
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
    </div>`;
  return card;
}
