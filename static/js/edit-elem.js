// Delete
const modalDeleteElem = document.getElementById("modal-delete");
const modalDeleteControl = new bootstrap.Modal(modalDeleteElem, {
  keyboard: false,
});
const modalDeleteBodyElem = document.getElementById("modal-delete-body");
const modalDeleteSubmitElem = document.getElementById("modal-delete-submit");
var submitDeleteFunc = () => {};
modalDeleteSubmitElem.addEventListener("click", submitDeleteFunc);

modalDeleteElem.addEventListener("show.bs.modal", function (event) {
  let btn = event.relatedTarget;
  let resource = btn.getAttribute("data-bs-recoursetype");
  let tableType = btn.getAttribute("data-bs-tableType");
  let index = btn.getAttribute("data-bs-index");
  let body;
  addDeleteEventToSubmit(resource, tableType, index);
  switch (tableType) {
    case attr.bulletin.tableType:
      body = items[optCurrIndex].getBulletin().getRow(index).getContent();
      break;
    case attr.slide.tableType:
      body = items[optCurrIndex].getSlide().getRow(index).getFileTitle();
      break;
    case attr.homework.tableType:
      body = items[optCurrIndex].getHomework().getRow(index).getFileTitle();
      break;
  }
  modalDeleteBodyElem.innerText = body;
});
function addDeleteEventToSubmit(resource, tableType, index) {
  modalDeleteSubmitElem.removeEventListener("click", submitDeleteFunc);
  submitDeleteFunc = () => {
    let itemId = items[optCurrIndex].getId();
    let rowId = items[optCurrIndex][tableType].getRow(index).getId();
    let url =
      api.getTeacherPath() + `/${resource}/${itemId}/${tableType}/${rowId}`;
    axios
      .delete(url, { headers: headers })
      .then((res) => {
        // 200
        if (res.status == HttpStatusCode.OK) {
          // remove the row[index]
          items[optCurrIndex].lastModified = res.data.data.lastModified;
          items[optCurrIndex][tableType].rows.splice(index, 1);
          items[optCurrIndex].buildContent();
          showContent(items[optCurrIndex].content);
        }
      })
      .catch((err) => {
        // 400, 401, 404
        console.error(err);
        if (err.response) {
          errHandler(err.response.status);
        }
      });
    modalDeleteControl.hide();
  };
  modalDeleteSubmitElem.addEventListener("click", submitDeleteFunc);
}
// Create & Update
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

var inputHomeworkNumberElem;
var inputHomeworkTitleElem;
var inputHomeworkFileElem;

var modalEditSubmitElem = modalEditElem.querySelector("#modal-edit-submit");
var submitFunc = () => {};
modalEditSubmitElem.addEventListener("click", submitFunc);

const submitBtnTexts = {
  create: "ADD",
  update: "UPDATE",
};

modalEditElem.addEventListener("show.bs.modal", function (event) {
  let btn = event.relatedTarget;
  let recourseType = btn.getAttribute("data-bs-recourseType");
  let tableType = btn.getAttribute("data-bs-tableType");
  let method = btn.getAttribute("data-bs-method");
  let index = btn.getAttribute("data-bs-index");

  let item = items[optCurrIndex];
  let itemId = item.id;
  let rowId;
  let formElem;

  switch (tableType) {
    case attr.bulletin.tableType:
      formElem = createBulletinInputElem();
      break;
    case attr.slide.tableType:
      formElem = createSlideInputElem();
      break;
    case attr.homework.tableType:
      formElem = createHomeworkInputElem();
      break;
  }
  modalTitle.textContent = attr[tableType].tableTitle;
  modalFormElem.innerHTML = formElem;

  refreshInputElem(modalEditElem);

  switch (method) {
    case HttpMethod.POST:
      modalEditSubmitElem.textContent = submitBtnTexts.create;
      break;
    case HttpMethod.PUT:
      let row = item[tableType].getRow(index);
      rowId = row.id;
      modalEditSubmitElem.textContent = submitBtnTexts.update;
      switch (tableType) {
        case attr.bulletin.tableType:
          inputBulletinContentElem.value = row.getContent();
          break;
        case attr.slide.tableType:
          inputSlideChapterElem.value = row.getChapter();
          inputSlideTitleElem.value = row.getFileTitle();
          break;
        case attr.homework.tableType:
          inputHomeworkNumberElem.value = row.getNumber();
          inputHomeworkTitleElem.value = row.getFileTitle();
          break;
      }
      break;
  }

  updateSubmitEditEvent(recourseType, tableType, method, itemId, rowId, index);
});
function refreshInputElem(modalElem) {
  inputBulletinContentElem = modalElem.querySelector("#bulletin-content");

  inputSlideChapterElem = modalElem.querySelector("#slide-chapter");
  inputSlideTitleElem = modalElem.querySelector("#slide-title");
  inputSlideFileElem = modalElem.querySelector("#slide-file");

  inputHomeworkNumberElem = modalEditElem.querySelector("#homework-number");
  inputHomeworkTitleElem = modalEditElem.querySelector("#homework-title");
  inputHomeworkFileElem = modalEditElem.querySelector("#homework-file");

  modalEditSubmitElem = modalElem.querySelector("#modal-edit-submit");
}

function updateSubmitEditEvent(
  recourseType,
  tableType,
  method,
  itemId,
  rowId,
  itemsIndex
) {
  // post or put
  url =
    api.getTeacherPath() +
    api.getResourceUrl(recourseType, tableType, method, itemId, rowId);

  modalEditSubmitElem.removeEventListener("click", submitFunc);
  submitFunc = () => {
    let params;
    switch (tableType) {
      case attr.bulletin.tableType:
        // json type
        params = { content: inputBulletinContentElem.value };
        break;
      case attr.slide.tableType:
        // form type
        params = new FormData();
        params.append("chapter", inputSlideChapterElem.value);
        params.append("fileTitle", inputSlideTitleElem.value);
        if (inputSlideFileElem.files[0] != undefined) {
          params.append("file", inputSlideFileElem.files[0]);
        }
        break;
      case attr.homework.tableType:
        // form type
        params = new FormData();
        params.append("number", inputHomeworkNumberElem.value);
        params.append("fileTitle", inputHomeworkTitleElem.value);
        if (inputHomeworkFileElem.files[0] != undefined) {
          params.append("file", inputHomeworkFileElem.files[0]);
        }
        break;
    }

    switch (method) {
      case HttpMethod.POST:
        apiCreateField(tableType, url, params);
        break;
      case HttpMethod.PUT:
        apiUpdateField(tableType, url, params, itemsIndex);
        break;
    }
  };
  modalEditSubmitElem.addEventListener("click", submitFunc);
}

function apiCreateField(tableType, url, params) {
  let _params = {};
  axios
    .post(url, params, axiosConfig)
    .then((res) => {
      // code = 201
      if (res.status == HttpStatusCode.CREATED) {
        let resData = res.data.data;

        switch (tableType) {
          case attr.bulletin.tableType:
            _params = params;
            _params.id = resData.id;
            _params.date = resData.date;
            break;
          case attr.slide.tableType:
            _params.id = resData.id;
            _params.chapter = params.get("chapter");
            _params.fileTitle = params.get("fileTitle");
            _params.filename = resData.filename;
            break;
          case attr.homework.tableType:
            _params.id = resData.id;
            _params.number = params.get("number");
            _params.fileTitle = params.get("fileTitle");
            _params.filename = resData.filename;
            break;
        }
        let rows = newRows(tableType, [_params]);
        items[optCurrIndex].lastModified = resData.lastModified;

        switch (tableType) {
          case attr.bulletin.tableType:
            items[optCurrIndex][tableType].getRows().splice(0, 0, rows[0]);
            break;
          case attr.slide.tableType:
            items[optCurrIndex][tableType].getRows().push(rows[0]);
            break;
          case attr.homework.tableType:
            items[optCurrIndex][tableType].getRows().push(rows[0]);
            break;
        }
        items[optCurrIndex].buildContent();
        showContent(items[optCurrIndex].getContent());
        modalEditElemControl.hide();
      }
    })
    .catch((err) => {
      console.error(
        `tableType:${tableType},url:${url}, file:${params.get("file")}`
      );
      if (err.response) {
        errHandler(err.response.status);
      }
    });
}
function apiUpdateField(tableType, url, params, itemsIndex) {
  axios
    .put(url, params, axiosConfig)
    .then((res) => {
      // code = 200
      if (res.status == HttpStatusCode.OK) {
        let resData = res.data.data;
        items[optCurrIndex].setLastModified(resData.lastModified);
        // update item content
        switch (tableType) {
          case attr.bulletin.tableType:
            items[optCurrIndex][tableType]
              .getRow(itemsIndex)
              .setContent(params.content);
            break;
          case attr.slide.tableType:
            items[optCurrIndex][tableType]
              .getRow(itemsIndex)
              .setChapter(params.get("chapter"));
            items[optCurrIndex][tableType]
              .getRow(itemsIndex)
              .setFileTitle(params.get("fileTitle"));
            if (params.get("file") != undefined) {
              items[optCurrIndex][tableType]
                .getRow(itemsIndex)
                .setFilename(resData.filename);
            }
            break;
          case attr.homework.tableType:
            items[optCurrIndex][tableType]
              .getRow(itemsIndex)
              .setNumber(params.get("number"));
            items[optCurrIndex][tableType]
              .getRow(itemsIndex)
              .setFileTitle(params.get("fileTitle"));
            if (params.get("file") != undefined) {
              items[optCurrIndex][tableType]
                .getRow(itemsIndex)
                .setFilename(resData.filename);
            }
            break;
        }
        items[optCurrIndex].buildContent();
        showContent(items[optCurrIndex].getContent());
        modalEditElemControl.hide();
      }
    })
    .catch((err) => {
      // 400, 401, 409
      if (err.response) {
        console.error(
          `tableType:${tableType}, url:${url}, params:${params}, itemsIndex: ${itemsIndex}`
        );
        if (err.response) {
          errHandler(err.response.status);
        }
      }
    });
}

function createBulletinInputElem() {
  return `<div class="mb-3">
  <label for="bulletin-content" class="col-form-label">Content:</label>
  <textarea class="form-control" id="bulletin-content"" rows="3"></textarea>
  </div>`;
}
function createSlideInputElem() {
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
function createHomeworkInputElem() {
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
function createAddButtonElem(recourseType, tableType) {
  let recourseTypeAttr = `data-bs-recourseType="${recourseType}"`;
  let methodAttr = `data-bs-method="${HttpMethod.POST}"`;
  let typeAttr = `data-bs-tableType="${tableType}"`;

  return `<button type="button" class="btn btn-outline-primary" data-bs-toggle="modal" 
  data-bs-target="#modal-edit" ${recourseTypeAttr} ${methodAttr} ${typeAttr}>ADD</button>`;
}

function createEditButtonElem(recourseType, tableType, rowIndex) {
  let recourseTypeAttr = `data-bs-recourseType="${recourseType}"`;
  let methodAttr = `data-bs-method="${HttpMethod.PUT}"`;
  let typeAttr = `data-bs-tableType="${tableType}"`;
  let indexAttr = `data-bs-index="${rowIndex}"`;

  return `<button type="button" class="btn" data-bs-toggle="modal"
  data-bs-target="#modal-edit" ${recourseTypeAttr} ${methodAttr} ${indexAttr} ${typeAttr}><i class="fa fa-pencil-square-o" aria-hidden="true"></i></button>`;
}

function createDeleteButtonElem(recourseType, tableType, rowIndex) {
  let recourseTypeAttr = `data-bs-recourseType="${recourseType}"`;
  let typeAttr = `data-bs-tableType="${tableType}"`;
  let indexAttr = `data-bs-index="${rowIndex}"`;

  return `<button type="button" class="btn" data-bs-toggle="modal"
  data-bs-target="#modal-delete" ${recourseTypeAttr} ${typeAttr} ${indexAttr}><i class="fa fa-trash" aria-hidden="true"></i></button>`;
}
function createFileBtnElem(resourceName, filename) {
  let elem = "";
  if (filename != undefined && filename != "") {
    let url =
      "/static/doc/" +
      api.getTeacherDomain() +
      "/" +
      resourceName +
      "/" +
      filename;
    elem = `<a href="${url}" target="_blank"><i class="fa fa-file-text" aria-hidden="true" ></i></a>`;
  }

  return elem;
}
