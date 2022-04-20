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
    apiUrl,
    id,
    nameZh,
    nameUs,
    bulletin,
    slide,
    homework,
    lastUpdated = "0"
  ) {
    this.apiUrl = apiUrl;
    this.id = id;
    this.nameZh = nameZh;
    this.nameUs = nameUs;
    this.bulletin = bulletin;
    this.slide = slide;
    this.homework = homework;
    this.lastUpdated = lastUpdated;
    this.content = "";
  }
  // todo: info or course
  createOptionButton() {
    return;
  }
  getContent(rebuild) {
    let _content;
    if (rebuild) {
      _content = createContent(this.bulletin, this.slide, this.homework);
    } else {
      _content = this.content;
    }
    return _content;
  }
  getData() {
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
class BulletinBoardRow {
  constructor(id, date, info) {
    this.id = id;
    this.date = date;
    this.info = info;
  }
  getDataList() {
    return [this.date, this.info];
  }
}
class SlideRow {
  constructor(id, chapter, fileTitle, fileType) {
    this.id = id;
    this.chapter = chapter;
    this.fileTitle = fileTitle;
    this.fileType = fileType;
  }
  getDataList() {
    return ["CH" + this.chapter, this.fileTitle, this.fileType];
  }
}
class HomeworkRow {
  constructor(id, number, fileTitle, fileType) {
    this.id = id;
    this.number = number;
    this.fileTitle = fileTitle;
    this.fileType = fileType;
  }
  getDataList() {
    return ["#" + this.number, this.fileTitle, this.fileType];
  }
}
class Table {
  constructor(title, fieldsTitle, rows) {
    this.title = title;
    this.fieldsTitle = fieldsTitle;
    this.rows = rows;
  }
  getTitle() {
    this.title;
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
  setRows(rows) {
    this.rows = rows;
  }
}
