class Item {
  constructor(
    /** @type {number} */
    itemIndex,
    /** @type {number} */
    recourseType,
    /** @type {string} */
    apiUrl = "",
    /** @type {number} */
    id,
    /** @type {string} */
    nameZh = "",
    /** @type {string} */
    nameUs = "",
    /** @type {Table} */
    bulletin,
    /** @type {Table} */
    slide,
    /** @type {Table} */
    homework,
    /** @type {string} */
    lastModified = "0",
    /** @type {string} */
    content = ""
  ) {
    this.itemIndex = itemIndex;
    this.recourseType = recourseType;
    this.apiUrl = apiUrl;
    this.id = id;
    this.nameZh = nameZh;
    this.nameUs = nameUs;
    this.bulletin = bulletin;
    this.slide = slide;
    this.homework = homework;
    this.lastModified = lastModified;
    this.content = content;
  }
  getItemIndex() {
    return this.itemIndex;
  }
  setItemIndex(index) {
    this.itemIndex = index;
  }
  getRecourseType() {
    return this.recourseType;
  }
  setRecourseType(recourseType) {
    this.recourseType = recourseType;
  }
  getApiUrl() {
    return this.apiUrl;
  }
  setApiUrl(url) {
    this.apiUrl = url;
  }
  getId() {
    return this.id;
  }
  setId(id) {
    this.id = id;
  }
  getNameZh() {
    return this.nameZh;
  }
  setNameZh(name) {
    this.nameZh = name;
  }
  getNameUs() {
    return this.nameUs;
  }
  setNameUs(name) {
    this.nameUs = name;
  }
  getLastModified() {
    return this.lastModified;
  }
  setLastModified(lastModified) {
    this.lastModified = lastModified;
  }
  getBulletin() {
    return this.bulletin;
  }
  setBulletin(bulletin) {
    this.bulletin = bulletin;
  }
  getSlide() {
    return this.slide;
  }
  setSlide(slide) {
    this.slide = slide;
  }
  getHomework() {
    return this.homework;
  }
  setHomework(homework) {
    this.homework = homework;
  }
  getContent() {
    return this.content;
  }
  setContent(content) {
    this.content = content;
  }
  buildContent() {
    let content = "";
    content += createContent(
      this.recourseType,
      attr.bulletin.tableType,
      this.bulletin
    );
    content += createContent(
      this.recourseType,
      attr.slide.tableType,
      this.slide
    );
    content += createContent(
      this.recourseType,
      attr.homework.tableType,
      this.homework
    );
    this.content = content;
  }
  updateData() {
    optCurrIndex = this.itemIndex;
    axios
      .get(this.apiUrl, {
        params: {
          lastModified: this.lastModified,
        },
      })
      .then((res) => {
        let resData = res.data.data;
        let rows;
        switch (res.status) {
          // The information is up to date(Not need to updating the data)
          case HttpStatusCode.NO_CONTENT:
            console.warn("The data is up to date!");
            break;
          // need to update information
          case HttpStatusCode.OK:
            this.lastModified = resData.lastModified;
            if (resData.bulletins != null && resData.bulletins != undefined) {
              rows = newRows(attr.bulletin.tableType, resData.bulletins);
              this.bulletin = newTable(attr.bulletin.tableType, rows);
            }
            if (resData.slides != null && resData.slides != undefined) {
              rows = newRows(attr.slide.tableType, resData.slides);
              this.slide = newTable(attr.slide.tableType, rows);
            }
            if (resData.homeworks != null && resData.homeworks != undefined) {
              rows = newRows(attr.homework.tableType, resData.homeworks);
              this.homework = newTable(attr.homework.tableType, rows);
            }
            this.buildContent();
            break;
        }
        showContent(this.content);
      })
      .catch((err) => {
        console.log(err);
        switch (err.response.statue) {
          case HttpStatusCode.BAD_REQUEST:
            console.error("bad request");
            break;
          default:
            alert("Unknown error, please reload the page again.");
            break;
        }
      });
  }
}
function newItem(
  itemIndex,
  recourseType,
  apiUrl,
  data,
  bulletin,
  slide,
  homework,
  lastModified
) {
  return new Item(
    itemIndex,
    recourseType,
    apiUrl,
    data.id,
    data.nameZh,
    data.nameUs,
    bulletin,
    slide,
    homework,
    lastModified
  );
}
class Table {
  constructor(title = "", fieldsTitle = [], rows = []) {
    this.title = title;
    this.fieldsTitle = fieldsTitle;
    this.rows = rows;
  }
  getTitle() {
    return this.title;
  }
  setTitle(title) {
    this.title = title;
  }
  getFieldsTitle() {
    return this.fieldsTitle;
  }
  getRowsLen() {
    return this.rows.length;
  }
  getRow(index) {
    return this.rows[index];
  }
  getRows() {
    return this.rows;
  }
  setRows(rows) {
    this.rows = rows;
  }
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
        rows.push(newBulletinBoardRow(v.id, v.date, v.content));
      });
      break;
    case attr.slide.tableType:
      data.forEach((v) => {
        rows.push(newSlideRow(v.id, v.chapter, v.fileTitle, v.filename));
      });
      break;
    case attr.homework.tableType:
      data.forEach((v) => {
        rows.push(newHomeworkRow(v.id, v.number, v.fileTitle, v.filename));
      });
      break;
  }
  return rows;
}
class BulletinBoardRow {
  constructor(id, date, content) {
    this.id = id;
    this.date = date;
    this.content = content;
  }
  getId() {
    return this.id;
  }
  setId(id) {
    this.id = id;
  }
  getDate() {
    return this.date;
  }
  setDate(date) {
    this.date = date;
  }
  getContent() {
    return this.content;
  }
  setContent(content) {
    this.content = content;
  }
  getDataList() {
    return [this.date, this.content];
  }
}
function newBulletinBoardRow(id, date, content) {
  return new BulletinBoardRow(id, date, content);
}
class SlideRow {
  constructor(id, chapter, fileTitle, filename) {
    this.id = id;
    this.chapter = chapter;
    this.fileTitle = fileTitle;
    this.filename = filename;
  }
  getId() {
    return this.id;
  }
  setId(id) {
    this.id = id;
  }
  getChapter() {
    return this.chapter;
  }
  setChapter(chapter) {
    this.chapter = chapter;
  }
  getFileTitle() {
    return this.fileTitle;
  }
  setFileTitle(title) {
    this.fileTitle = title;
  }
  getFilename() {
    return this.filename;
  }
  setFilename(name) {
    this.filename = name;
  }
  getDataList() {
    return [
      "CH" + this.chapter,
      this.fileTitle,
      createFileBtn("slide", this.filename),
    ];
  }
}
function newSlideRow(id, chapter, fileTitle, filename) {
  return new SlideRow(id, chapter, fileTitle, filename);
}
class HomeworkRow {
  constructor(id, number, fileTitle, filename) {
    this.id = id;
    this.number = number;
    this.fileTitle = fileTitle;
    this.filename = filename;
  }
  getId() {
    return this.id;
  }
  setId(id) {
    this.id = id;
  }
  getNumber() {
    return this.number;
  }
  setNumber(number) {
    this.number = number;
  }
  getFileTitle() {
    return this.fileTitle;
  }
  setFileTitle(title) {
    this.fileTitle = title;
  }
  getFilename() {
    return this.filename;
  }
  setFilename(name) {
    this.filename = name;
  }
  getDataList() {
    return [
      "#" + this.number,
      this.fileTitle,
      createFileBtn("homework", this.filename),
    ];
  }
}
function newHomeworkRow(id, number, fileTitle, filename) {
  return new HomeworkRow(id, number, fileTitle, filename);
}
function createFileBtn(resourceName, filename) {
  let url =
    "/static/doc/" +
    api.getTeacherDomain() +
    "/" +
    resourceName +
    "/" +
    filename;
  return `<a href="${url}" target="_blank"><i class="fa fa-file-text" aria-hidden="true" ></i></a>`;
}
