const br_tag = "<br>";
const elemOptionBox = document.getElementById("option-box");
const contentPageElem = document.getElementById("content-switch");
const loadingElem = document.getElementById("loading");
const pageTypes = {
  info: "info",
  course: "course",
};
const attr = {
  bulletin: {
    tableType: "bulletin",
    tableTitle: "Bulletin Board",
    tableFieldTitles: ["Date", "Information"],
  },
  slide: {
    tableType: "slide",
    tableTitle: "Slide",
    tableFieldTitles: ["Chapter", "Type", "Title"],
  },
  homework: {
    tableType: "homework",
    tableTitle: "Homework",
    tableFieldTitles: ["#", "Type", "Title"],
  },
};
//? temporary
var apiData;

var options = [];
var currPageType;

var isTeacher = false;
function init() {
  axios
    .post(api.getVerifyAuthUrl(), {
      params: {
        last_modified: "0",
      },
    })
    .then((res) => {
      if (res.status == HTTP_STATUS_CODE.ok) {
        if (res.data.isTeacher) {
          console.log("init api success, is teacher");
          headers[headerKeys.auth] = res.headers.authorization;
          teacherMode();
        }
        createInitElem();
      }
    })
    .catch((err) => {
      console.error("init api error, not a teacher:", err);
      createInitElem();
    });
}
function teacherMode() {
  let editTxt = "Edit";
  let deleteTxt = "Delete";
  isTeacher = true;
  attr.bulletin.tableFieldTitles.push(editTxt, deleteTxt);
  attr.slide.tableFieldTitles.push(editTxt, deleteTxt);
  attr.homework.tableFieldTitles.push(editTxt, deleteTxt);
}

function createInitElem() {
  // let course_get_url = `/course`;
  getInfoApi();
  // getCourseApi(course_get_url);
}
// *

const turn = {
  on: true,
  off: false,
};
function loadingView(switcher = false) {
  if (switcher) {
    loadingElem.style.display = "block";
    return;
  }
  loadingElem.style.display = "none";
}
function getInfoApi() {
  // [ bulletins<list>(id, date, content), id(info), last_modified(info) ]
  console.log("API -> getInfoApi");
  let url =
    api.teacherPath + api.getResourceUrl(pageTypes.info, null, HTTP_METHOD.get);
  axios
    .get(url, {
      params: {
        last_modified: "0",
      },
    })
    .then((res) => {
      console.log("getInfo api success");
      apiData = res.data.data;

      let bulletinRows = newRows(attr.bulletin.tableType, apiData.bulletins);
      let bulletinTable = newTable(attr.bulletin.tableType, bulletinRows);
      let infoItem = new Item(
        pageTypes.info,
        url,
        apiData.id,
        "公布欄",
        "Information",
        bulletinTable,
        null,
        null,
        apiData.last_modified,
        createContent(pageTypes.info, attr.bulletin.tableType, bulletinTable)
      );
      items.push(infoItem);
      loadingView(turn.off);
      showContent(infoItem.content);
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
      apiData = res.data.data;
      // isTeacher(res.data.auth);
      //* Course(#The b, s and h is undefined)
      for (let i = 0; i < apiData.courses.length; i++) {
        // let item = newItem(pageType.course, apiData.courses[i]);
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

function newItem(
  pageType,
  apiUrl,
  data,
  bulletin,
  slide,
  homework,
  lastModified
) {
  // nameZh: any, nameUs: any, bulletin: any, slide: any, homework: any)
  return new Item(
    pageType,
    apiUrl,
    data.id,
    data.name_zh,
    data.name_us,
    bulletin,
    slide,
    homework,
    lastModified
  );
}
function newTable(tableType, rows) {
  let title = attr[tableType].tableTitle;
  let tableFieldTitles = attr[tableType].tableFieldTitles;
  return new Table(title, tableFieldTitles, rows);
}
function newRows(tableType, data) {
  let rows = [];
  switch (tableType) {
    case attr.bulletin.tableType:
      data.forEach((v) => {
        rows.push(new BulletinBoardRow(v.id, v.date, v.content));
      });
      break;
    case attr.slide.tableType:
      data.forEach((v) => {
        rows.push(new SlideRow(v.id, v.chapter, v.file.title, v.file.type));
      });
      break;
    case attr.homework.tableType:
      data.forEach((v) => {
        rows.push(new HomeworkRow(v.id, v.number, v.file.title, v.file.type));
      });
      break;
  }
  return rows;
}
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
function createContent(pageType, tableType, data) {
  console.log("createContent");
  let content = "";
  if (data != null && data != undefined) {
    content += createTable(pageType, tableType, data);
  }
  return content;
}

// * create element
function createTable(pageType, tableType, table) {
  console.log("createTable");

  let thead = createTableHeadElem(table.fieldsTitle);
  let tbody = "";
  let addBtnElem = "";
  let editBtnElem = "";
  let deleteBtnElem = "";

  for (let rowIndex = 0; rowIndex < table.rowsLen; rowIndex++) {
    dataList = table.rows[rowIndex].dataList;
    if (isTeacher) {
      editBtnElem = createEditButtonElem(pageType, tableType, rowIndex);
      deleteBtnElem = createDeleteButtonElem(pageType, tableType, rowIndex);
    }
    tbody += createTableBodyElem(dataList, editBtnElem, deleteBtnElem);
  }
  if (isTeacher) {
    addBtnElem = createAddButtonElem(pageType, tableType);
  }

  return createCard(table.title, thead, tbody, addBtnElem);
}

function createTableHeadElem(fields) {
  let elems = "";
  for (let i = 0; i < fields.length; i++) {
    elems += `<td>${fields[i]}</td>`;
  }
  return `<tr>${elems}</tr>`;
}
function createTableBodyElem(dataList, editBtnElem, deleteBtnElem) {
  let elems = "";
  for (let i = 0; i < dataList.length; i++) {
    elems += `<td>${dataList[i]}</td>`;
  }
  if (isTeacher) {
    elems += `<td>${editBtnElem}</td><td>${deleteBtnElem}</td>`;
  }
  return `<tr>${elems}</tr>`;
}
function createCard(title, thead, tbody, addBtnElem) {
  let card = `
    <div class="content-page">
      <div class="item-box">
        <div class="item-title h2">
          <div class="card">
            <div class="card-body">${title}</div>
          </div>
        </div>
        <div class = "item-content">
          <table class = "table table-striped table-striped">
            <thead>${thead}</thead> 
            <tbody>${tbody}</tbody>
          </table>
          ${addBtnElem}
        </div>
      </div>
    </div>`;
  return card;
}
