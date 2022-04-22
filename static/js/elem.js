var bulletinFormModal = document.getElementById("form-modal-bulletin");
var slideFormModal = document.getElementById("form-modal-slide");
var homeworkFormModal = document.getElementById("form-modal-homework");

function createEditButton(type, rowIndex) {
  let typeAttr = `data-bs-type="${type}"`;
  let indexAttr = `data-bs-index="${rowIndex}"`;

  return `<button type="button" class="btn" data-bs-toggle="modal"
  data-bs-target="#form-modal" ${indexAttr} ${typeAttr}><i class="fa fa-pencil-square-o" aria-hidden="true"></i></button>`;
}

function createBulletinInput(params) {
  return `<div class="mb-3">
  <label for="bulletin-content" class="col-form-label">Content:</label>
  <input type="text" class="form-control" id="bulletin-content" />
  </div>`;
}
function createSlideInput(params) {
  return `<div class="mb-3">
  <label for="slide-chapter" class="col-form-label">Chapter:</label>
  <input type="text" class="form-control" id="slide-chapter" /></div>

  <div class="mb-3">
  <label for="slide-title" class="col-form-label">Title:</label>
  <input type="text" class="form-control" id="slide-title" />
  </div>`;
}
function createHomeworkInput(params) {
  return `<div class="mb-3">
  <label for="homework-number" class="col-form-label">Number:</label>
  <input type="text" class="form-control" id="homework-number" /></div>

  <div class="mb-3">
  <label for="homework-title" class="col-form-label">Title:</label>
  <input type="text" class="form-control" id="homework-title" />
  </div>`;
}

var modalElem = document.getElementById("form-modal");
var modalTitle = modalElem.querySelector(".modal-title");
var modalFormElem = document.getElementById("modal-form");
var modalFormSubmitBtn = modalElem.querySelector("#submit");

modalElem.addEventListener("show.bs.modal", function (event) {
  var btn = event.relatedTarget;
  var type = btn.getAttribute("data-bs-type");
  var index = btn.getAttribute("data-bs-index");
  switch (type) {
    case tableType.bulletin:
      modalFormElem.innerHTML = createBulletinInput();
      bulletinForm(modalElem, index);
      break;
    case tableType.slide:
      modalFormElem.innerHTML = createSlideInput();
      slideForm(btn, modalElem);
      break;
    // case homeworkType:
    //   homeworkForm(btn, modalElem);
    //   break;
    default:
      break;
  }
  // testSendForm(url, data);
});
const formTitle = {
  bulletin: "Bulletin",
  slide: "Slide",
  homework: "Homework",
};
function bulletinForm(modalElem, index) {
  modalTitle.textContent = formTitle.bulletin;

  // var inputDateElem = modalElem.querySelector("#bulletin-date");
  var inputContentElem = modalElem.querySelector("#bulletin-content");
  // optionSwitchIndex is global variable
  let row = items[optionSwitchIndex].bulletin.rows[index];
  // inputDateElem.value = row.date;
  inputContentElem.value = row.content;
  modalFormSubmitBtn.addEventListener("click", function name() {
    // var url = api.url + "/test/1";
    // todo: get update.slide url
    console.log("go to bulletin api");
    // todo: call api
    // todo: when response success, then change the object value
    // var data = {
    //   create_date: inputDateElem.value,
    //   info: inputContentElem.value,
    // };
  });
}
function slideForm(modalElem, index) {
  modalTitle.textContent = formTitle.bulletin;

  var chapterElem = modalElem.querySelector("#slide-chapter");
  var titleElem = modalElem.querySelector("#slide-file-title");
  // todo: upload form
  // var inputContentElem = modalElem.querySelector("#slide-file-title");
  // optionSwitchIndex is global variable
  let row = item[optionSwitchIndex].slide().row(index);
  chapterElem.value = row.chapter;
  titleElem.value = row.fileTitle;
  modalFormSubmitBtn.addEventListener("click", function name() {
    // var url = api.url + "/test/1";
    // todo: get update.slide url
    console.log("go to bulletin api");
    // todo: call api
    // var data = {
    //   create_date: inputDateElem.value,
    //   info: inputContentElem.value,
    // };
  });
}
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
