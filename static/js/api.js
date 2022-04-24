class API {
  constructor(origin, version, teacherDomain) {
    this.origin = origin;
    this.version = version;
    this.teacherDomain = teacherDomain;
  }
  get url() {
    return this.origin + "/api/" + this.version + "/" + this.teacherDomain;
  }
}
const apiResourceUrl = {
  info: {
    // 200
    get: "/info/bulletin?last_modified=",
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
};
