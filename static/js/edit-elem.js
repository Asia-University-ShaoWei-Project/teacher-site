function createAddButtonElem(pageType, tableType) {
  let pageTypeAttr = `data-bs-pageType="${pageType}"`;
  let methodAttr = `data-bs-method="${HTTP_METHOD.post}"`;
  let typeAttr = `data-bs-tableType="${tableType}"`;

  return `<button type="button" class="btn btn-outline-primary" data-bs-toggle="modal" 
  data-bs-target="#modal-edit" ${pageTypeAttr} ${methodAttr} ${typeAttr}>ADD</button>`;
}

function createEditButtonElem(pageType, tableType, rowIndex) {
  let pageTypeAttr = `data-bs-pageType="${pageType}"`;
  let methodAttr = `data-bs-method="${HTTP_METHOD.put}"`;
  let typeAttr = `data-bs-tableType="${tableType}"`;
  let indexAttr = `data-bs-index="${rowIndex}"`;

  return `<button type="button" class="btn" data-bs-toggle="modal"
  data-bs-target="#modal-edit" ${pageTypeAttr} ${methodAttr} ${indexAttr} ${typeAttr}><i class="fa fa-pencil-square-o" aria-hidden="true"></i></button>`;
}

function createDeleteButtonElem(pageType, tableType, rowIndex) {
  let pageTypeAttr = `data-bs-pageType="${pageType}"`;
  let typeAttr = `data-bs-tableType="${tableType}"`;
  let indexAttr = `data-bs-index="${rowIndex}"`;

  return `<button type="button" class="btn" data-bs-toggle="modal"
  data-bs-target="#modal-delete" ${pageTypeAttr} ${typeAttr} ${indexAttr}><i class="fa fa-trash" aria-hidden="true"></i></button>`;
}
// ? delete
var modalDeleteElem = document.getElementById("modal-delete");
var modalDeleteControl = new bootstrap.Modal(modalDeleteElem, {
  keyboard: false,
});
var modalDeleteBodyElem = document.getElementById("modal-delete-body");
var modalDeleteSubmitElem = document.getElementById("modal-delete-submit");
var submitDeleteFunc = () => {};
modalDeleteSubmitElem.addEventListener("click", submitDeleteFunc);

modalDeleteElem.addEventListener("show.bs.modal", function (event) {
  let btn = event.relatedTarget;
  let tableType = btn.getAttribute("data-bs-tableType");
  let index = btn.getAttribute("data-bs-index");
  let body;
  updateSubmitDeleteEvent(tableType, index);
  switch (tableType) {
    case attr.bulletin.tableType:
      body = items[optionSwitchIndex].bulletin.rows[index].content;
      break;
    case attr.slide.tableType:
      body = items[optionSwitchIndex].slide.rows[index].fileTitle;
      break;
    case attr.homework.tableType:
      body = items[optionSwitchIndex].homework.rows[index].fileTitle;
      break;
  }
  modalDeleteBodyElem.innerText = body;
});
function updateSubmitDeleteEvent(tableType, index) {
  modalDeleteSubmitElem.removeEventListener("click", submitDeleteFunc);
  submitDeleteFunc = () => {
    let resource = "/" + items[optionSwitchIndex].pageType;
    let itemID = items[optionSwitchIndex].id;
    let rowID = items[optionSwitchIndex][tableType].rows[index].id;
    let url = api.teacherPath + `${resource}/${itemID}/${tableType}/${rowID}`;

    console.log("delete url:", url);
    axios
      .delete(url, { headers: headers })
      .then((res) => {
        // code = 200
        if (res.status == HTTP_STATUS_CODE.ok) {
          // remove the row[index]
          items[optionSwitchIndex].lastModified = res.data.data.last_modified;
          items[optionSwitchIndex][tableType].rows.splice(index, 1);
          items[optionSwitchIndex].buildContent();
          showContent(items[optionSwitchIndex].content);
        }
      })
      // todo:test
      .catch((err) => {
        // 404
        alert("delete error");
        console.error(err);
      });
    modalDeleteControl.hide();
  };

  modalDeleteSubmitElem.addEventListener("click", submitDeleteFunc);
}
// ? create and update
var modalEditElem = document.getElementById("modal-edit");
var modalEditControl = new bootstrap.Modal(modalEditElem, { keyboard: false });
var modalTitle = modalEditElem.querySelector(".modal-title");
var modalFormElem = document.getElementById("modal-form");
var modalEditElemControl = new bootstrap.Modal(modalEditElem, {
  keyboard: false,
});
var inputBulletinContentElem;

var inputSlideChapterElem;
var inputSlideTitleElem;
var inputSlideFileElem;

var modalEditSubmitElem = modalEditElem.querySelector("#modal-edit-submit");
var submitFunc = () => {};
modalEditSubmitElem.addEventListener("click", submitFunc);

// add, edit
const submitBtnTexts = {
  create: "ADD",
  update: "UPDATE",
};

modalEditElem.addEventListener("show.bs.modal", function (event) {
  let btn = event.relatedTarget;
  let pageType = btn.getAttribute("data-bs-pageType");
  let tableType = btn.getAttribute("data-bs-tableType");
  let method = btn.getAttribute("data-bs-method");
  let rowIndex = btn.getAttribute("data-bs-index");

  let item = items[optionSwitchIndex];
  let itemID = item.id;
  let rowID;
  let url, formElem;

  switch (tableType) {
    case attr.bulletin.tableType:
      formElem = createBulletinInputElem();
      break;
    case attr.slide.tableType:
      formElem = createSlideInputElem();
      break;
    // todo: homework
    // case attr.homework.tableType:
    //   break;
  }
  modalTitle.textContent = attr[tableType].tableTitle;
  modalFormElem.innerHTML = formElem;

  refreshInputElem(modalEditElem);

  switch (method) {
    case HTTP_METHOD.post:
      modalEditSubmitElem.textContent = submitBtnTexts.create;
      // url = api.getResourceUrl(pageType, method, itemID);
      break;
    case HTTP_METHOD.put:
      let row = item[tableType].rows[rowIndex];
      rowID = row.id;
      modalEditSubmitElem.textContent = submitBtnTexts.update;
      // url = api.getResourceUrl(pageType, method, itemID, rowID, tableType);

      switch (tableType) {
        case attr.bulletin.tableType:
          inputBulletinContentElem.value = row.content;
          break;
        case attr.slide.tableType:
          inputSlideChapterElem.value = row.chapter;
          inputSlideTitleElem.value = row.fileTitle;
          break;
        // todo: homework
        // case attr.homework.tableType:
        //   inputHomeworkNumberElem.value = row.number;
        //   inputHomeworkTitleElem.value = row.fileTitle;
        //   break;
      }
      break;
  }
  // todo:test
  // url = api.getResourceUrl(pageType, method, itemID, rowID, tableType);

  updateSubmitEditEvent(pageType, tableType, method, itemID, rowID);
});
function refreshInputElem(modalElem) {
  inputBulletinContentElem = modalElem.querySelector("#bulletin-content");

  inputSlideChapterElem = modalElem.querySelector("#slide-chapter");
  inputSlideTitleElem = modalElem.querySelector("#slide-title");
  inputSlideFileElem = modalElem.querySelector("#slide-file");

  modalEditSubmitElem = modalElem.querySelector("#modal-edit-submit");

  // todo: homework
  // let inputHomeworkNumberElem = modalEditElem.querySelector("#homework-chapter");
  // let inputHomeworkTitleElem = modalEditElem.querySelector("#homework-title");
  // let inputHomeworkFileElem = modalEditElem.querySelector("#homework-file");
}

function updateSubmitEditEvent(pageType, tableType, method, itemID, rowID) {
  // post or put
  url =
    api.teacherPath +
    api.getResourceUrl(pageType, tableType, method, itemID, rowID);

  modalEditSubmitElem.removeEventListener("click", submitFunc);
  submitFunc = () => {
    let params;
    switch (tableType) {
      case attr.bulletin.tableType:
        params = { content: inputBulletinContentElem.value };
        break;
      case attr.slide.tableType:
        params = {
          chapter: inputSlideChapterElem.value,
          title: inputSlideTitleElem.value,
          file: inputSlideFileElem.value,
        };
        break;
      // todo: homework
      // case attr.homework.tableType:
      //   params = {
      //     number: hNumberElem.value,
      //     title: hTitleElem.value,
      //     file: hFileElem.value,
      //   };
      //   break;
    }
    console.log("URL:", url, "|| Method:", method, "|| Params", params);
    switch (method) {
      case HTTP_METHOD.post:
        console.log("use: createFieldApi()");
        createFieldApi(tableType, url, params);
        break;
      case HTTP_METHOD.put:
        console.log("use: updateFieldApi()");
        updateFieldApi(tableType, url, params, rowID);
        break;
    }
  };
  modalEditSubmitElem.addEventListener("click", submitFunc);
}

function createFieldApi(tableType, url, params) {
  axios
    .post(url, params, axiosConfig)
    .then((res) => {
      // code = 201
      if (res.status == HTTP_STATUS_CODE.created) {
        let resData = res.data.data;
        params.id = resData.id;
        switch (tableType) {
          case attr.bulletin.tableType:
            params.date = resData.date;
            break;
          case attr.slide.tableType:
            break;
          // todo: homework
          // case attr.homework.tableType:
          //   break;
        }
        let rows = newRows(tableType, [params]);
        items[optionSwitchIndex].lastModified = resData.last_modified;
        items[optionSwitchIndex][tableType].rows.splice(0, 0, rows[0]);
        items[optionSwitchIndex].buildContent();
        showContent(items[optionSwitchIndex].content);
        modalEditElemControl.hide();
      }
    })
    .catch((err) => {
      if (err.response) {
        switch (err.response.status) {
          case HTTP_STATUS_CODE.unauthorized:
            alert("驗證過期，請重新登入");
            break;
          case HTTP_STATUS_CODE.badRequest:
            alert("參數錯誤");
            break;
        }
      }
    });
}
function updateFieldApi(tableType, url, params, rowID) {
  axios
    .put(url, params, axiosConfig)
    .then((res) => {
      // code = 200
      if (res.status == HTTP_STATUS_CODE.ok) {
        let resData = res.data.data;
        console.log("resData:", resData);
        items[optionSwitchIndex].lastModified = resData.last_modified;
        // update item content
        params.id = resData.id;
        switch (tableType) {
          case attr.bulletin.tableType:
            items[optionSwitchIndex][tableType].rows[rowID].content =
              params.content;
            break;
          case attr.slide.tableType:
            items[optionSwitchIndex][tableType].rows[rowID].chapter =
              params.content;
            items[optionSwitchIndex][tableType].rows[rowID].fileTitle =
              params.fileTitle;
          // todo: file url
          // items[optionSwitchIndex][tableType].rows[rowID].chapter = params.content;
          // todo: homework
          // case attr.homework.tableType:
          //   items[optionSwitchIndex][tableType].rows[rowID].number =
          //     params.number;
          //   items[optionSwitchIndex][tableType].rows[rowID].fileTitle =
          //     params.fileTitle;
          //   break;
        }
        items[optionSwitchIndex].buildContent();
        showContent(items[optionSwitchIndex].content);
        modalEditElemControl.hide();
      }
    })
    .catch((err) => {
      // 400, 409
      // todo:test
      console.error(err);
      if (err.response.status == HTTP_STATUS_CODE.badRequest) {
        alert("update error");
        console.log("Bad request");
      }
    });
}

function createBulletinInputElem(params) {
  return `<div class="mb-3">
  <label for="bulletin-content" class="col-form-label">Content:</label>
  <textarea class="form-control" id="bulletin-content"" rows="3"></textarea>
  </div>`;
}
function createSlideInputElem(params) {
  return `<div class="mb-3">
  <label for="slide-chapter" class="col-form-label">Chapter:</label>
  <input type="text" class="form-control" id="slide-chapter" /></div>

  <div class="mb-3">
  <label for="slide-title" class="col-form-label">Title:</label>
  <textarea class="form-control" id="slide-title"" rows="3"></textarea></div>

  <div class="input-group mb-3">
  <input type="file" class="form-control" id="slide-file">
  <label class="input-group-text" for="slide-file">Upload</label></div>
  
  `;
}
function createHomeworkInputElem(params) {
  return `<div class="mb-3">
  <label for="homework-number" class="col-form-label">Number:</label>
  <input type="text" class="form-control" id="homework-number" /></div>

  <div class="mb-3">
  <label for="homework-title" class="col-form-label">Title:</label>
  <textarea class="form-control" id="homework-title" rows="3"></textarea></div>
  
  <div class="input-group mb-3">
  <input type="file" class="form-control" id="homework-file"> 
  <label class="input-group-text" for="homework-file">Upload</label></div>
  `;
}
//? tmp
function bug() {
  console.log(items[0].bulletin.rows);
}
