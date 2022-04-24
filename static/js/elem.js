function createAddButtonElem(pageType, tableType) {
  let pageTypeAttr = `data-bs-pageType="${pageType}"`;
  let methodAttr = `data-bs-method="${HTTP_METHOD.post}"`;
  let typeAttr = `data-bs-tableType="${tableType}"`;

  return `<button type="button" class="btn btn-primary" data-bs-toggle="modal" 
  data-bs-target="#modal-edit" ${pageTypeAttr} ${methodAttr} ${typeAttr}>Add</button>`;
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

modalDeleteElem.addEventListener("show.bs.modal", function (event) {
  let btn = event.relatedTarget;
  let tableType = btn.getAttribute("data-bs-tableType");
  let index = btn.getAttribute("data-bs-index");
  let body;
  addEventToDeleteSubmit(tableType, index);
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
  // testSendForm(url, data);
});
function addEventToDeleteSubmit(tableType, index) {
  modalDeleteSubmitElem.addEventListener("click", function () {
    let resource = "/" + items[optionSwitchIndex].pageType;
    let itemID = items[optionSwitchIndex].id;
    let rowID = items[optionSwitchIndex][tableType].rows[index].id;
    let url = api.url + `${resource}/${itemID}/${tableType}/${rowID}`;

    console.log("delete url:", url);
    axios
      .delete(url, { headers: headers })
      .then((res) => {
        // code = 200
        if (res.status == HTTP_STATUS_CODE.ok) {
          // remove the row[index]
          items[optionSwitchIndex][tableType].rows.splice(index, 1);
          items[optionSwitchIndex].buildContent();
          showContent(items[optionSwitchIndex].content);
        }
      })
      .catch((err) => {
        // 404
        alert("delete error");
        console.error(err);
      });
    modalDeleteControl.hide();
  });
}
// ? create and update
var modalEditElem = document.getElementById("modal-edit");
var modalEditControl = new bootstrap.Modal(modalEditElem, { keyboard: false });
var modalTitle = modalEditElem.querySelector(".modal-title");
var modalFormElem = document.getElementById("modal-form");

let inputBulletinContentElem;

let inputSlideChapterElem;
let inputSlideTitleElem;
let inputSlideFileElem;

var modalEditSubmitElem;
// let inputHomeworkNumberElem = modalEditElem.querySelector("#homework-chapter");
// let inputHomeworkTitleElem = modalEditElem.querySelector("#homework-title");
// let inputHomeworkFileElem = modalEditElem.querySelector("#homework-file");

// add, edit
modalEditElem.addEventListener("show.bs.modal", function (event) {
  let btn = event.relatedTarget;
  let tableType = btn.getAttribute("data-bs-tableType");
  let method = btn.getAttribute("data-bs-method");
  let rowIndex = btn.getAttribute("data-bs-index");
  let pageType = btn.getAttribute("data-bs-pageType");
  let url, formElem;
  let item = items[optionSwitchIndex];
  let itemID = item.id;
  switch (tableType) {
    case attr.bulletin.tableType:
      formElem = createBulletinInputElem();
      break;
    case attr.slide.tableType:
      formElem = createSlideInputElem();
      break;
    // case attr.homework.tableType:
    //   break;
  }
  modalTitle.textContent = attr[tableType].tableTitle;
  modalFormElem.innerHTML = formElem;

  inputBulletinContentElem = modalEditElem.querySelector("#bulletin-content");

  inputSlideChapterElem = modalEditElem.querySelector("#slide-chapter");
  inputSlideTitleElem = modalEditElem.querySelector("#slide-title");
  inputSlideFileElem = modalEditElem.querySelector("#slide-file");

  modalEditSubmitElem = modalEditElem.querySelector("#modal-edit-submit");
  switch (method) {
    case HTTP_METHOD.post:
      modalEditSubmitElem.textContent = "ADD";

      url = newApiUrl(pageType, method, tableType, itemID);
      break;
    case HTTP_METHOD.put:
      let row = item[tableType].rows[rowIndex];
      let rowID = row.id;
      url = newApiUrl(pageType, method, tableType, itemID, rowID);
      modalEditSubmitElem.textContent = "UPDATE";

      switch (tableType) {
        case attr.bulletin.tableType:
          inputBulletinContentElem.value = row.content;
          break;
        case attr.slide.tableType:
          inputSlideChapterElem.value = row.chapter;
          inputSlideTitleElem.value = row.fileTitle;
          break;
        // case attr.homework.tableType:
        //   inputHomeworkNumberElem.value = row.number;
        //   inputHomeworkTitleElem.value = row.fileTitle;
        //   break;
      }
      break;
  }
  addEventToEditSubmit(tableType, method, url);
});
function newApiUrl(pageType, method, tableType, itemID, rowID) {
  let url = apiResourceUrl[pageType][method];
  switch (pageType) {
    case pageTypes.info:
      url = url.replace(":info_id", itemID).replace(":row_id", rowID);
      break;
    case pageTypes.course:
      url = url
        .replace(":course_id", itemID)
        .replace(":tableType", tableType)
        .replace(":row_id", rowID);
      break;
  }
  return url;
}

function addEventToEditSubmit(tableType, method, url) {
  // post or put
  modalEditSubmitElem.addEventListener("click", function () {
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
      // case attr.homework.tableType:
      //   params = {
      //     number: hNumberElem.value,
      //     title: hTitleElem.value,
      //     file: hFileElem.value,
      //   };
      //   break;
    }
    console.log("Method:", method);
    console.log("URL:", url);
    console.log("Params", params);
    switch (method) {
      case HTTP_METHOD.post:
        console.log("use: createFieldApi()");
        // createFieldApi(url, params);
        break;
      case HTTP_METHOD.put:
        console.log("use: updateFieldApi()");
        // updateFieldApi(url, params);
        break;
    }
  });
}

function createFieldApi(url, params) {
  axios
    .post(url, params, axiosConfig)
    .then((res) => {
      // code = 201
      if (res.status == HTTP_STATUS_CODE.created) {
        alert("add success");
        // todo:call rebuild
      }
    })
    .catch((err) => {
      // 400
      if (err.response.status == HTTP_STATUS_CODE.badRequest) {
        console.log("Bad request");
      }
      alert("add error");
      console.error(err);
    });
}
function updateFieldApi(url, params) {
  axios
    .put(url, params, axiosConfig)
    .then((res) => {
      // code = 200
      if (res.status == HTTP_STATUS_CODE.ok) {
        // todo:call rebuild
        alert("update work");
      }
    })
    .catch((err) => {
      // 400, 409
      if (err.response.status == HTTP_STATUS_CODE.badRequest) {
        console.log("Bad request");
      }
      alert("update error");
      console.error(err);
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
