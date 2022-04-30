const brTag = "<br>";
const optionBoxElem = document.getElementById("option-box");
const contentElem = document.getElementById("content");
const loadingElem = document.getElementById("loading");
const authBtnElem = document.getElementById("auth-btn");

var /** !Object<!Item> */ items = [];

// 0: information index
var optCurrIndex = 0;
var isTeacher = false;

const attr = {
  bulletin: {
    tableType: "bulletin",
    tableTitle: "Bulletin Board",
    tableFieldTitles: [
      `<i class="fa fa-calendar" aria-hidden="true"></i> Date`,
      "Information",
    ],
  },
  slide: {
    tableType: "slide",
    tableTitle: "Slide",
    tableFieldTitles: ["chapter", "title", "file"],
  },
  homework: {
    tableType: "homework",
    tableTitle: "Homework",
    tableFieldTitles: ["#", "title", "file"],
  },
};
function init() {
  axios
    .post(api.getVerifyAuthUrl())
    .then((res) => {
      if (res.status == HttpStatusCode.OK) {
        if (res.data.isTeacher) {
          headers[HeaderKeys.AUTH] = res.headers.authorization;
          teacherMode();
        }
        createInitElem();
        setupAuthBtn();
      }
    })
    .catch((err) => {
      console.error("init api error, not a teacher:", err);
      createInitElem();
      setupAuthBtn();
    });
}
const Texts = {
  TABLE_FIELD: { EDIT: "Edit", DEL: "Delete" },
};
function setupAuthBtn() {
  if (isTeacher) {
    authBtnElem.innerText = "Logout";
    let url =
      api.getUrlPath() +
      api.getResourceUrl(api.getAuthResourceType(), null, HttpMethod.POST);
    authBtnElem.onclick = () => {
      axios
        .post(url, {}, axiosConfig)
        .then((res) => {
          if (res.status == HttpStatusCode.NO_CONTENT) {
            location.reload();
          }
        })
        .catch((err) => {
          location.reload();
        });
    };
    return;
  }
  authBtnElem.innerText = "Login";
  authBtnElem.onclick = () => {
    location.replace("/login");
  };
}
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

  let work = setInterval(() => {
    if (count >= timeoutCount) {
      alert("connect error");
      loadingView(false);
      clearInterval(work);
    }
    if (infoWorkDone && courseWorkDone) {
      loadingView(false);
      showContent(items[0].getContent());
      showOptionButtons(items);
      clearInterval(work);
    }
    count += 1;
  }, delay);
  initInfoApi();
  initCourseApi();
}

function initInfoApi() {
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
      let infoIndex = 0;
      let resData = res.data.data;
      let bulletinRows = newRows(attr.bulletin.tableType, resData.bulletins);
      let bulletinTable = newTable(attr.bulletin.tableType, bulletinRows);
      let infoItem = new Item(
        infoIndex,
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
    })
    .catch((err) => {
      infoWorkDone = true;
      console.error("getInfoApi:", err);
    });
}

function initCourseApi() {
  let url =
    api.getTeacherPath() +
    api.getResourceUrl(api.getCourseResourceType(), null, HttpMethod.GET);
  axios
    .get(url)
    .then((res) => {
      if (res.status == HttpStatusCode.OK) {
        let resData = res.data.data;
        let apiUrl;
        var count = 0;
        // 5 second
        let timeoutTime = 5;
        let delay = 100;
        let timeoutCount = timeoutTime * (1000 /* 1 second */ / delay);
        let work = setInterval(() => {
          console.log("waiting for info api...");
          if (count >= timeoutCount) {
            alert("connect error");
            loadingView(false);
            clearInterval(work);
          }
          if (infoWorkDone) {
            resData.courses.forEach((v, itemIndex) => {
              apiUrl = url + "/" + v.id;
              items.push(
                new Item(
                  // +1: 0 is information index, so the course index begin at 1
                  itemIndex + 1,
                  api.getCourseResourceType(),
                  apiUrl,
                  v.id,
                  v.nameZh,
                  v.nameUs
                )
              );
            });
            courseWorkDone = true;
            clearInterval(work);
          }
          count += 1;
        }, delay);
      }
    })
    .catch((err) => {
      console.error("getInfoApi:", err);
      courseWorkDone = true;
    });
}

//* content
function reloadContent() {
  contentElem.innerHTML = "";
  loadingView(true);
}
function showContent(content = "") {
  loadingView(false);
  contentElem.innerHTML = content;
}
function createContent(recourseType, tableType, data) {
  let content = "";
  if (data != null && data != undefined) {
    content += createTable(recourseType, tableType, data);
  }
  return content;
}

// * create element
function createTable(recourseType, tableType, table) {
  let thead = createTableHeadElem(table.getFieldsTitle());
  let tbody = "";
  let addBtnElem, editBtnElem, deleteBtnElem;

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
    elems += `<th scope="col">${fields[i]}</th>`;
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
  if (addBtnElem == undefined) {
    addBtnElem = "";
  }
  let card = `
    <div class="content-page">
      <div class="item-box">
        <div class="item-title h2">
          <div class="card">
            <div class="card-body">${title}</div>
          </div>
        </div>
        <div class="table-responsive">
          <table class="table table-striped table-sm">
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
