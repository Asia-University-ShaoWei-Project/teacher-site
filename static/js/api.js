class API {
  constructor(origin, version, teacherDomain, recourseType, resource) {
    this.origin = origin;
    this.version = version;
    this.teacherDomain = teacherDomain;
    this.recourseType = recourseType;
    this.resource = resource;
  }
  getUrlPath() {
    return this.origin + "/api/" + this.version;
  }
  getTeacherDomain() {
    return this.teacherDomain;
  }
  getTeacherPath() {
    return this.getUrlPath() + "/" + this.teacherDomain;
  }
  getVerifyAuthUrl() {
    return this.getUrlPath() + "/auth/token";
  }
  getAuthResourceType() {
    return this.recourseType.AUTH;
  }
  getInfoResourceType() {
    return this.recourseType.INFO;
  }
  getCourseResourceType() {
    return this.recourseType.COURSE;
  }
  getResourceUrl(recourseType, tableType, method, itemId, rowId) {
    let url;
    switch (recourseType) {
      case this.recourseType.INFO:
        url = this.resource.info[method];
        url = url.replace(":infoId", itemId).replace(":rowId", rowId);
        break;
      case this.recourseType.COURSE:
        if (method == "get") {
          url = this.resource.course[method];
        } else {
          url = this.resource.courseContent[method];
        }
        url = url
          .replace(":courseId", itemId)
          .replace(":tableType", tableType)
          .replace(":rowId", rowId);
        break;
      case this.recourseType.AUTH:
        url = this.resource.auth[method];
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
  // resource type
  {
    AUTH: "auth",
    INFO: "info",
    COURSE: "course",
    COURSE_CONTENT: "courseContent",
  },
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
    auth: {
      post: "/auth/logout",
    },
  }
);
