class Item {
  constructor(
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
  // todo
  updateData() {
    // apiUrl = [ api.resources.info | api.resources.course]
    axios
      .get(this.apiUrl, {
        params: {
          lastModified: this.lastModified,
        },
      })
      .then((res) => {
        let rebuild = true;
        let data = res.data.data;
        let rows;
        switch (res.status) {
          // The information is up to date(Not need to updating the data)
          case HttpStatusCode.NO_CONTENT:
            rebuild = false;
            alert("The data is up to date!");
            console.warn("The data is up to date!");
            break;
          // need to update information
          case HttpStatusCode.OK:
            console.log("update the content");

            this.lastModified = data.lastModified;
            if (data.bulletins != null && data.bulletins != undefined) {
              rows = newRows(attr.bulletin.tableType, data.bulletins);
              this.bulletin = newTable(attr.bulletin.tableType, rows);
            }
            if (data.slides != null && data.slides != undefined) {
              rows = newRows(attr.slide.tableType, data.slides);
              this.slide = newTable(attr.slide.tableType, data.slides);
            }
            if (data.homeworks != null && data.homeworks != undefined) {
              rows = newRows(attr.homework.tableType, data.homeworks);
              this.homework = newTable(attr.slide.tableType, data.homeworks);
            }
            this.buildContent();
            showContent(this.content);
            break;
        }
      })
      .catch((err) => {
        console.log(err);
        console.log(err.response);
        console.log(err.response.state);
        switch (err.response.statue) {
          case HttpStatusCode.BAD_REQUEST:
            console.error("bad request");
            break;
          default:
            console.error("error status code:", err.response.statue);
            break;
        }
      });
  }
}
function newItem(
  recourseType,
  apiUrl,
  data,
  bulletin,
  slide,
  homework,
  lastModified
) {
  return new Item(
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
        rows.push(newSlideRow(v.id, v.chapter, v.file.title, v.file.type));
      });
      break;
    case attr.homework.tableType:
      data.forEach((v) => {
        rows.push(newHomeworkRow(v.id, v.number, v.file.title, v.file.type));
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
  constructor(id, chapter, fileTitle, fileType) {
    this.id = id;
    this.chapter = chapter;
    this.fileTitle = fileTitle;
    this.fileType = fileType;
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
  getFileType() {
    return this.fileType;
  }
  setFileType(type) {
    this.fileType = type;
  }
  setFataList() {
    return ["CH" + this.chapter, this.fileTitle, this.fileType];
  }
}
function newSlideRow(id, chapter, fileTitle, fileType) {
  return new SlideRow(id, chapter, fileTitle, fileType);
}
class HomeworkRow {
  constructor(id, number, fileTitle, fileType) {
    this.id = id;
    this.number = number;
    this.fileTitle = fileTitle;
    this.fileType = fileType;
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
  getFileType() {
    return this.fileType;
  }
  setFileType(type) {
    this.fileType = type;
  }
  setFataList() {
    return ["#" + this.number, this.fileTitle, this.fileType];
  }
}
function newHomeworkRow(id, number, fileTitle, fileType) {
  return new HomeworkRow(id, number, fileTitle, fileType);
}
