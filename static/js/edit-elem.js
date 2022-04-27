// ? delete
const modalDeleteElem = document.getElementById("modal-delete");
const modalDeleteControl = new bootstrap.Modal(modalDeleteElem, {
  keyboard: false,
});
const modalDeleteBodyElem = document.getElementById("modal-delete-body");
const modalDeleteSubmitElem = document.getElementById("modal-delete-submit");
var submitDeleteFunc = () => {};
modalDeleteSubmitElem.addEventListener("click", submitDeleteFunc);

modalDeleteElem.addEventListener("show.bs.modal", function (event) {
  console.log("open window");
  let btn = event.relatedTarget;
  let tableType = btn.getAttribute("data-bs-tableType");
  let index = btn.getAttribute("data-bs-index");
  let body;
  updateSubmitDeleteEvent(tableType, index);
  switch (tableType) {
    case attr.bulletin.tableType:
      console.log("is bulletin type");
      body = items[optionSwitchIndex].getBulletin().getRow(index).getContent();
      break;
    case attr.slide.tableType:
      console.log("is slide type");

      body = items[optionSwitchIndex].getSlide().getRow(index).getFileTitle();
      break;
    case attr.homework.tableType:
      body = items[optionSwitchIndex]
        .getHomework()
        .getRow(index)
        .getFileTitle();
      break;
  }
  modalDeleteBodyElem.innerText = body;
});
function updateSubmitDeleteEvent(tableType, index) {
  console.log("update my event");
  modalDeleteSubmitElem.removeEventListener("click", submitDeleteFunc);
  submitDeleteFunc = () => {
    console.log("exec api call");
    let resource = "/" + items[optionSwitchIndex].getRecourseType();
    let itemId = items[optionSwitchIndex].getId();
    let rowId = items[optionSwitchIndex][tableType].getRow(index).getId();
    let url =
      api.getTeacherPath() + `${resource}/${itemId}/${tableType}/${rowId}`;

    axios
      .delete(url, { headers: headers })
      .then((res) => {
        // 200
        if (res.status == HttpStatusCode.OK) {
          // remove the row[index]
          items[optionSwitchIndex].lastModified = res.data.data.lastModified;
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
        if (err.response) {
          switch (err.response.status) {
            case HttpStatusCode.NO_FOUND:
              alert("Not found");
              break;
          }
        }
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
  let recourseType = btn.getAttribute("data-bs-recourseType");
  let tableType = btn.getAttribute("data-bs-tableType");
  let method = btn.getAttribute("data-bs-method");
  let rowIndex = btn.getAttribute("data-bs-index");

  let item = items[optionSwitchIndex];
  let itemId = item.id;
  let rowId;
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
    case HttpMethod.POST:
      modalEditSubmitElem.textContent = submitBtnTexts.create;
      // url = api.getResourceUrl(pageType, method, itemId);
      break;
    case HttpMethod.PUT:
      let row = item[tableType].getRow(rowIndex);
      rowId = row.id;
      modalEditSubmitElem.textContent = submitBtnTexts.update;
      // url = api.getResourceUrl(pageType, method, itemId, rowId, tableType);

      switch (tableType) {
        case attr.bulletin.tableType:
          inputBulletinContentElem.value = row.getContent();
          break;
        case attr.slide.tableType:
          inputSlideChapterElem.value = row.getChapter();
          inputSlideTitleElem.value = row.getFileTitle();
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
  // url = api.getResourceUrl(pageType, method, itemId, rowId, tableType);

  updateSubmitEditEvent(recourseType, tableType, method, itemId, rowId);
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

function updateSubmitEditEvent(recourseType, tableType, method, itemId, rowId) {
  // post or put
  url =
    api.getTeacherPath() +
    api.getResourceUrl(recourseType, tableType, method, itemId, rowId);

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
      case HttpMethod.POST:
        console.log("use: createFieldApi()");
        createFieldApi(tableType, url, params);
        break;
      case HttpMethod.PUT:
        console.log("use: updateFieldApi()");
        updateFieldApi(tableType, url, params, rowId);
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
      if (res.status == HttpStatusCode.CREATED) {
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
        items[optionSwitchIndex].lastModified = resData.lastModified;
        items[optionSwitchIndex][tableType].getRows().splice(0, 0, rows[0]);
        items[optionSwitchIndex].buildContent();
        showContent(items[optionSwitchIndex].getContent());
        modalEditElemControl.hide();
      }
    })
    .catch((err) => {
      if (err.response) {
        switch (err.response.status) {
          case HttpStatusCode.UNAUTHORIZED:
            alert("驗證過期，請重新登入");
            break;
          case HttpStatusCode.BAD_REQUEST:
            alert("參數錯誤");
            break;
        }
      }
    });
}
function updateFieldApi(tableType, url, params, rowId) {
  axios
    .put(url, params, axiosConfig)
    .then((res) => {
      // code = 200
      if (res.status == HttpStatusCode.OK) {
        let resData = res.data.data;
        console.log("resData:", resData);
        items[optionSwitchIndex].setLastModified(resData.lastModified);
        // update item content
        params.id = resData.id;
        switch (tableType) {
          case attr.bulletin.tableType:
            items[optionSwitchIndex][tableType]
              .getRow(rowId)
              .setContent(params.content);
            break;
          case attr.slide.tableType:
            items[optionSwitchIndex][tableType]
              .getRow(rowId)
              .setChapter(params.chapter);
            items[optionSwitchIndex][tableType]
              .getRow(rowId)
              .setFileTitle(params.fileTitle);
          // todo: file url
          // items[optionSwitchIndex][tableType].rows[rowId].chapter = params.content;
          // todo: homework
          // case attr.homework.tableType:
          //   items[optionSwitchIndex][tableType].rows[rowId].number =
          //     params.number;
          //   items[optionSwitchIndex][tableType].rows[rowId].fileTitle =
          //     params.fileTitle;
          //   break;
        }
        items[optionSwitchIndex].buildContent();
        showContent(items[optionSwitchIndex].getContent());
        modalEditElemControl.hide();
      }
    })
    .catch((err) => {
      // 400, 409
      // todo:test
      console.error(err);
      if (err.response.status == HttpStatusCode.BAD_REQUEST) {
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
