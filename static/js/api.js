class API {
  constructor(origin, version, teacherDomain, resource) {
    this.origin = origin;
    this.version = version;
    this.teacherDomain = teacherDomain;
    this.resource = resource;
  }
  get urlPath() {
    return this.origin + "/api/" + this.version;
  }
  get teacherPath() {
    return this.urlPath + "/" + this.teacherDomain;
  }
  get resource() {
    return this._resource;
  }
  set resource(resource) {
    this._resource = resource;
  }
  getVerifyAuthUrl() {
    return this.urlPath + "/auth/token";
  }
  getResourceUrl(pageType, tableType, method, itemID, rowID) {
    let url = this._resource[pageType][method];
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
}

var api = new API(
  // origin(e.g. http://domain)
  window.location.origin,
  // version
  "v1",
  // teacher domain
  window.location.pathname.replace("/", ""),
  // resource
  {
    info: {
      // 200
      get: "/info/bulletin",
      // 201, 400
      post: "/info/:info_id/bulletin",
      // 200, 400, 409
      put: "/info/:info_id/bulletin/:row_id",
      // 200, 404
      delete: "/info/:info_id/bulletin/:row_id",
    },
    course: {
      // 200
      get: "/course",
      // 201, 400
      post: "/course",
    },
    courseContent: {
      // 200, 404
      get: "/course/:course_id/?last_modified",
      // 201, 400
      post: "/course/:course_id/:type",
      // 200, 400, 409
      put: "/course/:course_id/:table_type/:row_id",
      // 200, 404
      delete: "/course/:course_id/:table_type/:row_id",
    },
  }
);
