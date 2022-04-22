class API {
  constructor(origin, version, teacherDomain, resources) {
    this.origin = origin;
    this.version = version;
    this.teacherDomain = teacherDomain;
    this.resources = resources;
    this.url = this.origin + "/api/" + this.version + "/" + this.teacherDomain;
  }
  getResourceUrl(type) {
    return this.url + "/" + this.resources[type];
  }
}

class Item {
  constructor(
    apiUrl = "",
    id,
    nameZh = "",
    nameUs = "",
    bulletin,
    slide,
    homework,
    lastModified = "0"
  ) {
    this.apiUrl = apiUrl;
    this.id = id;
    this.nameZh = nameZh;
    this.nameUs = nameUs;
    this.bulletin = bulletin;
    this.slide = slide;
    this.homework = homework;
    this.lastModified = lastModified;
    this.content = "";
  }

  set id(id) {
    this._id = id;
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
  buildContent(rebuild) {
    if (rebuild) {
      let content = "";
      content += createContent(tableType.bulletin, this.bulletin);
      content += createContent(tableType.slide, this.slide);
      content += createContent(tableType.homework, this.homework);
      return content;
    }
    return this.content;
  }
  set content(content) {
    this._content = content;
  }
  updateData() {
    // apiUrl = [ api.resources.info | api.resources.course]
    let url = this.apiUrl + "/" + this.id + "/" + this.lastUpdateTime;
    axios
      .get(url)
      .then((res) => {
        let rebuild = true;
        switch (res.status) {
          // the information is up to date
          case HTTP_STATUS_CODE.noContent:
            rebuild = false;
            alert("the data is up to date!");
            break;
          // need to update information
          case HTTP_STATUS_CODE.ok:
            this.lastUpdated = res.data.data.last_updated;
            this.bulletin = newBulletin(res.data.data.bulletin_board);
            // todo
            // this.slide = newSlide()
            // this.homework = newHomework()
            break;
          // todo:
          default:
            break;
        }
        this.showContent(rebuild);
      })
      .catch((err) => {
        switch (err.response.status) {
          case HTTP_STATUS_CODE.badRequest:
            console.error("bad request");
            break;
          default:
            console.error("error status code:", err.response.status);
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
class SlideRow {
  constructor(id, chapter, fileTitle, fileType) {
    this.id = id;
    this.chapter = chapter;
    this.fileTitle = fileTitle;
    this.fileType = fileType;
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
  get dataList() {
    return ["CH" + this._chapter, this._fileTitle, this._fileType];
  }
}
// class HomeworkRow {
//   constructor(id, number, fileTitle, fileType) {
//     this.id = id;
//     this.number = number;
//     this.fileTitle = fileTitle;
//     this.fileType = fileType;
//   }
//   get dataList() {
//     return ["#" + this._number, this._fileTitle, this._fileType];
//   }
// }
