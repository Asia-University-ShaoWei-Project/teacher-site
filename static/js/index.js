const brTag = "<br>";
const optionBoxElem = document.getElementById("option-box");
const pageContentElem = document.getElementById("page-content");
const loadingElem = document.getElementById("loading");
var /** !Object<!Item> */ items = [];
// 0: information index
var optionSwitchIndex = 0;
const PageTypes = {
  INFO: "info",
  COURSE: "course",
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
    tableFieldTitles: ["Chapter", "Title", "Type"],
  },
  homework: {
    tableType: "homework",
    tableTitle: "Homework",
    tableFieldTitles: ["#", "Title", "Type"],
  },
};
var isTeacher = false;
function init() {
  axios
    .post(api.getVerifyAuthUrl(), {
      params: {
        lastModified: "0",
      },
    })
    .then((res) => {
      if (res.status == HttpStatusCode.OK) {
        if (res.data.isTeacher) {
          console.log("init api success, is teacher");
          headers[HeaderKeys.AUTH] = res.headers.authorization;
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
const Texts = {
  TABLE_FIELD: { EDIT: "Edit", DEL: "Delete" },
};
function teacherMode() {
  let editTxt = Texts.TABLE_FIELD.EDIT;
  let deleteTxt = Texts.TABLE_FIELD.DEL;
  isTeacher = true;
  attr.bulletin.tableFieldTitles.push(editTxt, deleteTxt);
  attr.slide.tableFieldTitles.push(editTxt, deleteTxt);
  attr.homework.tableFieldTitles.push(editTxt, deleteTxt);
}

var infoWorkDone = false;
var courseWorkDone = false;
function createInitElem() {
  var count = 0;
  // 5 second
  let timeoutTime = 5;
  let delay = 200;
  let timeoutCount = timeoutTime * (1000 /* 1 second */ / delay);

  initInfoApi();
  initCourseApi();
  let work = setInterval(() => {
    if (count >= timeoutCount) {
      console.log("timeout");
      alert("connect error");
      clearInterval(work);
    }
    if (infoWorkDone && courseWorkDone) {
      showOptionButtons();
      clearInterval(work);
    }
    console.log("add time");
    count += 1;
  }, delay);
}

function initInfoApi(workDone) {
  console.log("API -> getInfoApi");
  let url =
    api.getTeacherPath() +
    api.getResourceUrl(api.getInfoResourceType(), null, HttpMethod.GET);
  axios
    .get(url, {
      params: {
        lastModified: "0",
      },
    })
    .then((res) => {
      console.log("getInfo api success");
      let resData = res.data.data;
      let bulletinRows = newRows(attr.bulletin.tableType, resData.bulletins);
      let bulletinTable = newTable(attr.bulletin.tableType, bulletinRows);
      let infoItem = new Item(
        api.getInfoResourceType(),
        url,
        resData.id,
        "公布欄",
        "Information",
        bulletinTable,
        null,
        null,
        resData.lastModified,
        createContent(
          api.getInfoResourceType(),
          attr.bulletin.tableType,
          bulletinTable
        )
      );
      items.push(infoItem);
      infoWorkDone = true;
      loadingView(false);
      showContent(infoItem.getContent());
    })
    .catch((err) => {
      infoWorkDone = true;
      console.error("getInfoApi:", err);
    });
}
// todo

function initCourseApi(workDone) {
  let url =
    api.getTeacherPath() +
    api.getResourceUrl(api.getCourseResourceType(), null, HttpMethod.GET);
  axios
    .get(url)
    .then((res) => {
      console.log("getInfo api success");
      if (res.status == HttpStatusCode.OK) {
        let resData = res.data.data;
        let apiUrl;
        resData.courses.forEach((v) => {
          apiUrl = url + "/" + v.id;
          items.push(
            new Item(
              api.getCourseResourceType(),
              apiUrl,
              v.id,
              v.nameZh,
              v.nameUs
            )
          );
        });
      }
      courseWorkDone = true;
    })
    .catch((err) => {
      console.error("getInfoApi:", err);
      courseWorkDone = true;
    });
}

// *option

function showOptionButtons() {
  let elem = "";
  items.forEach((v) => {
    elem = createOptionButton(v);
    optionBoxElem.appendChild(elem);
  });
}

// const optionClassAttr = `option-item option-button button--anthe`;
const optionClassAttr = `btn btn-outline-dark option-item`;
function createOptionButton(item) {
  console.log("createOptionButton");
  let btn = document.createElement("button");
  let span = document.createElement("span");
  let text = item.getNameZh() + brTag + item.getNameUs();
  btn.className = optionClassAttr;
  span.innerHTML = text;
  btn.onclick = () => item.updateData();
  btn.appendChild(span);
  return btn;
}

//* content

function showContent(content = "") {
  pageContentElem.innerHTML = content;
}
function createContent(recourseType, tableType, data) {
  console.log("createContent");
  let content = "";
  if (data != null && data != undefined) {
    content += createTable(recourseType, tableType, data);
  }
  return content;
}

// * create element
function createTable(recourseType, tableType, table) {
  console.log("createTable");

  let thead = createTableHeadElem(table.getFieldsTitle());
  let tbody = "";
  let addBtnElem = "";
  let editBtnElem = "";
  let deleteBtnElem = "";

  for (let rowIndex = 0; rowIndex < table.getRowsLen(); rowIndex++) {
    dataList = table.getRow(rowIndex).getDataList();
    if (isTeacher) {
      editBtnElem = createEditButtonElem(recourseType, tableType, rowIndex);
      deleteBtnElem = createDeleteButtonElem(recourseType, tableType, rowIndex);
    }
    tbody += createTableBodyElem(dataList, editBtnElem, deleteBtnElem);
  }
  if (isTeacher) {
    addBtnElem = createAddButtonElem(recourseType, tableType);
  }

  return createCard(table.getTitle(), thead, tbody, addBtnElem);
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

function loadingView(turn = false) {
  if (turn) {
    loadingElem.style.display = "block";
    return;
  }
  loadingElem.style.display = "none";
}
