const br_tag = "<br>";
const elemOptionBox = document.getElementById("option-box");
const contentPageElem = document.getElementById("content-switch");
const COURSE = "course";
const INFO = "info";
const INIT = "init";
const tableType = {
  bulletin: "b",
  slide: "s",
  homework: "h",
};
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

var tableFieldTitle = {
  bulletin: ["Date", "Information"],
  slide: ["Chapter", "Type", "Title"],
  homework: ["#", "Type", "Title"],
};
var headers = {};

var isTeacher = false;
function init() {
  console.log("API -> /api/v1/auth/token");
  axios
    .post("/api/v1/auth/token", {
      params: {
        last_modified: "0",
      },
    })
    .then((res) => {
      console.log("init api success, is teacher");
      headers[headerKeys.auth] = res.headers.authorization;
      teacherMode();
      createInitElem();
    })
    .catch((err) => {
      console.error("init api error, not a teacher:", err);
      createInitElem();
    });
}
function teacherMode() {
  let editTxt = "Edit";
  isTeacher = true;
  tableFieldTitle.bulletin.push(editTxt);
  tableFieldTitle.slide.push(editTxt);
  tableFieldTitle.homework.push(editTxt);
}
function createInitElem() {
  let info_get_url = `/info/bulletin`;
  // let course_get_url = `/course`;
  getInfoApi(info_get_url);
  // getCourseApi(course_get_url);
}
// *
const rebuild = true;
function getInfoApi(url) {
  // [ bulletins<list>(id, date, content), id(info), last_modified(info) ]
  console.log("API -> getInfoApi -> ", api.url + url);
  axios
    .get(api.url + url, {
      params: {
        last_modified: "0",
      },
    })
    .then((res) => {
      console.log("getInfo api success");
      let data = res.data.data;
      // isTeacher(res.data.auth);
      // todo: create option for information button
      // todo: show the information view
      let infoItem = newItem(null, {
        name_zh: "公布欄",
        name_us: "Information",
      });
      let bulletin = newBulletin(data.bulletins);
      infoItem.id = data.id;
      infoItem.lastModified = data.last_modified;
      infoItem.bulletin = bulletin;
      infoItem.content = createContent(tableType.bulletin, bulletin);
      items.push(infoItem);
      showContent(items[0].buildContent(rebuild));
      // todo: temp
      showOptionButtons();
    })
    .catch((err) => {
      console.error("getInfoApi:", err);
    });
}
// todo
function getCourseApi(url) {
  console.log("API -> getCourse -> ", url);
  // [ bulletins<list>(id, date, content), id(info), last_modified(info) ]
  axios
    .get(url)
    .then((res) => {
      console.log("getCourseApi api success");
      console.log(res.data);
      let data = res.data.data;
      // isTeacher(res.data.auth);
      //* Course(#The b, s and h is undefined)
      for (let i = 0; i < data.courses.length; i++) {
        let item = newItem(api.getResourceUrl(COURSE), data.courses[i]);
        items.push();
        // items.push(createOptionButton(COURSE, optionSwitchIndex, data.courses[i]));
        // lastUpdatedOfCourse.push(zeroTime);
        // optionSwitchIndex += 1;
      }
      showOptionButtons();
    })
    .catch((err) => {
      console.error("getInitApiAndCreateView:", err);
    });
}

const zeroTime = "0";
var contents = [];
var items = [];
var lastUpdatedOfCourse = [];
// 0: information index
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
  let title = "Bulletin Board";
  let bulletinBoard = new Table(title, tableFieldTitle.bulletin);
  let rows = [];
  data.forEach((v) => {
    rows.push(new BulletinBoardRow(v.id, v.date, v.content));
  });
  bulletinBoard.rows = rows;
  return bulletinBoard;
}
// todo: slide, homework table
// function newSlide(data) {
//   let title = "Slide";
//   let slide = new Table(title, tableFieldTitle.slide);
//   let rows = [];
//   data.forEach((v) => {
//     rows.push(new SlideRow(v.id, v.chapter, v.content));
//   });
//   slide.setRows(rows);
//   return slide;
// }
// function newHomework(data) {
//   let title = "Slide";
//   let slide = new Table(title, tableFieldTitle.slide);
//   let rows = [];
//   data.forEach((v) => {
//     rows.push(new SlideRow(v.id, v.chapter, v.content));
//   });
//   slide.setRows(rows);
//   return slide;
// }
// *option

function showOptionButtons() {
  let elem;
  items.forEach((v) => {
    elem = createOptionButton(v);
    elemOptionBox.appendChild(elem);
  });
}

const optionClassAttr = `option-item button button--anthe`;
function createOptionButton(item) {
  console.log("createOptionButton");
  let btn = document.createElement("button");
  let span = document.createElement("span");
  let text = item.nameZh + br_tag + item.nameUs;
  btn.className = optionClassAttr;
  span.innerHTML = text;
  btn.onclick = () => item.updateData();
  btn.appendChild(span);
  return btn;
}

//* content

function showContent(content) {
  contentPageElem.innerHTML = content;
}
function createContent(type, data) {
  console.log("createContent");
  let content = "";
  if (data != null && data != undefined) {
    content += createTable(type, data);
  }
  return content;
}

// * create element
function createTable(type, table) {
  console.log("createTable");

  let thead = createTableHeadElem(table.fieldsTitle);
  let tbody = "";
  for (let rowIndex = 0; rowIndex < table.rowsLen; rowIndex++) {
    dataList = table.rows[rowIndex].dataList;
    tbody += createTableBodyElem(type, rowIndex, dataList);
  }
  return CreateCard(table.title, thead, tbody);
}

function createTableHeadElem(fields) {
  let elems = "";
  for (let i = 0; i < fields.length; i++) {
    elems += `<td>${fields[i]}</td>`;
  }
  return `<tr>${elems}</tr>`;
}
function createTableBodyElem(type, rowIndex, dataList) {
  let elems = "";
  for (let i = 0; i < dataList.length; i++) {
    elems += `<td>${dataList[i]}</td>`;
  }
  if (isTeacher) {
    let editBtn = createEditButton(type, rowIndex);
    elems += `<td>${editBtn}</td>`;
  }
  return `<tr>${elems}</tr>`;
}
function CreateCard(title, thead, tbody) {
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
