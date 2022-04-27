class API {
  constructor(origin, version, teacherDomain, resource) {
    this.origin = origin;
    this.version = version;
    this.teacherDomain = teacherDomain;
    this.recourseType = {
      INFO: "info",
      COURSE: "course",
      COURSE_CONTENT: "courseContent",
    };
    this.resource = resource;
  }
  getUrlPath() {
    return this.origin + "/api/" + this.version;
  }
  getTeacherPath() {
    return this.getUrlPath() + "/" + this.teacherDomain;
  }
  getVerifyAuthUrl() {
    return this.getUrlPath() + "/auth/token";
  }
  getInfoResourceType() {
    return this.recourseType.INFO;
  }
  getCourseResourceType() {
    return this.recourseType.COURSE;
  }
  getResourceUrl(recourseType, tableType, method, itemId, rowId) {
    let url = this.resource[recourseType][method];
    switch (recourseType) {
      case this.recourseType.INFO:
        url = url.replace(":infoId", itemId).replace(":rowId", rowId);
        break;
      case this.recourseType.COURSE:
        url = url
          .replace(":courseId", itemId)
          .replace(":tableType", tableType)
          .replace(":rowId", rowId);
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
      post: "/info/:infoId/bulletin",
      // 200, 400, 409
      put: "/info/:infoId/bulletin/:rowId",
      // 200, 404
      delete: "/info/:infoId/bulletin/:rowId",
    },
    course: {
      // 200
      get: "/course",
      // 201, 400
      post: "/course",
    },
    courseContent: {
      // 200, 404
      get: "/course/:courseId/?lastModified",
      // 201, 400
      post: "/course/:courseId/:tableType",
      // 200, 400, 409
      put: "/course/:courseId/:tableType/:rowId",
      // 200, 404
      delete: "/course/:courseId/:tableType/:rowId",
    },
  }
);
