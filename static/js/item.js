class Item {
  constructor(
    pageType,
    apiUrl = "",
    id,
    nameZh = "",
    nameUs = "",
    bulletin,
    slide,
    homework,
    lastModified = "0",
    content = ""
  ) {
    this.pageType = pageType;
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
  get pageType() {
    return this._pageType;
  }
  set pageType(pageType) {
    this._pageType = pageType;
  }
  get apiUrl() {
    return this._apiUrl;
  }
  set apiUrl(url) {
    this._apiUrl = url;
  }
  get id() {
    return this._id;
  }
  set id(id) {
    this._id = id;
  }
  get nameZh() {
    return this._nameZh;
  }
  set nameZh(name) {
    this._nameZh = name;
  }
  get nameUs() {
    return this._nameUs;
  }
  set nameUs(name) {
    this._nameUs = name;
  }
  get lastModified() {
    return this._lastModified;
  }
  set lastModified(lastModified) {
    this._lastModified = lastModified;
  }
  get bulletin() {
    return this._bulletin;
  }
  set bulletin(bulletin) {
    this._bulletin = bulletin;
  }
  get content() {
    return this._content;
  }
  set content(content) {
    this._content = content;
  }
  buildContent() {
    let content = "";
    content += createContent(
      this._pageType,
      attr.bulletin.tableType,
      this._bulletin
    );
    content += createContent(this._pageType, attr.slide.tableType, this._slide);
    content += createContent(
      this._pageType,
      attr.homework.tableType,
      this._homework
    );
    this._content = content;
  }
  // todo
  updateData() {
    // apiUrl = [ api.resources.info | api.resources.course]
    axios
      .get(this._apiUrl, {
        params: {
          last_modified: this._lastModified,
        },
      })
      .then((res) => {
        let rebuild = true;
        let data = res.data.data;
        let rows;
        switch (res.status) {
          // The information is up to date(Not need to updating the data)
          case HTTP_STATUS_CODE.noContent:
            rebuild = false;
            alert("The data is up to date!");
            console.warn("The data is up to date!");
            break;
          // need to update information
          case HTTP_STATUS_CODE.ok:
            console.log("update the content");

            this._lastModified = data.last_modified;
            if (data.bulletins != null && data.bulletins != undefined) {
              rows = newRows(attr.bulletin.tableType, data.bulletins);
              this._bulletin = newTable(attr.bulletin.tableType, rows);
            }
            if (data.slides != null && data.slides != undefined) {
              rows = newRows(attr.slide.tableType, data.slides);
              this._slide = newTable(attr.slide.tableType, data.slides);
            }
            if (data.homeworks != null && data.homeworks != undefined) {
              rows = newRows(attr.homework.tableType, data.homeworks);
              this._homework = newTable(attr.slide.tableType, data.homeworks);
            }
            this.buildContent();
            showContent(this._content);
            break;
        }
      })
      .catch((err) => {
        console.log(err);
        console.log(err.response);
        console.log(err.response.state);
        switch (err.response.statue) {
          case HTTP_STATUS_CODE.badRequest:
            console.error("bad request");
            break;
          default:
            console.error("error status code:", err.response.statue);
            break;
        }
      });
  }
}
class Table {
  constructor(title = "", fieldsTitle = [], rows = []) {
    this.title = title;
    this.fieldsTitle = fieldsTitle;
    this.rows = rows;
  }
  get title() {
    return this._title;
  }
  set title(title) {
    this._title = title;
  }
  get fieldTitles() {
    return this._fieldsTitle;
  }
  get rowsLen() {
    return this._rows.length;
  }
  get rows() {
    return this._rows;
  }
  set rows(rows) {
    this._rows = rows;
  }
}

class BulletinBoardRow {
  constructor(id, date, content) {
    this.id = id;
    this.date = date;
    this.content = content;
  }
  get id() {
    return this._id;
  }
  set id(id) {
    this._id = id;
  }
  get date() {
    return this._date;
  }
  set date(date) {
    this._date = date;
  }
  get content() {
    return this._content;
  }
  set content(content) {
    this._content = content;
  }
  get dataList() {
    return [this._date, this._content];
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
  get id() {
    return this._id;
  }
  set id(id) {
    this._id = id;
  }
  get chapter() {
    return this._chapter;
  }
  set chapter(chapter) {
    this._chapter = chapter;
  }
  get fileTitle() {
    return this._fileTitle;
  }
  set fileTitle(title) {
    this._fileTitle = title;
  }
  get fileType() {
    return this._fileType;
  }
  set fileType(type) {
    this._fileType = type;
  }
  get dataList() {
    return ["CH" + this._chapter, this._fileTitle, this._fileType];
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
  get number() {
    return this._number;
  }
  set number(number) {
    this._number = number;
  }
  get fileTitle() {
    return this._fileTitle;
  }
  set fileTitle(title) {
    this._fileTitle = title;
  }
  get fileType() {
    return this._fileType;
  }
  set fileType(type) {
    this._fileType = type;
  }
  get dataList() {
    return ["#" + this._number, this._fileTitle, this._fileType];
  }
}
function newHomeworkRow(id, number, fileTitle, fileType) {
  return new HomeworkRow(id, number, fileTitle, fileType);
}
