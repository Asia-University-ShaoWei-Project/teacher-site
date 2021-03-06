const HttpStatusCode = {
  OK: 200,
  CREATED: 201,
  NO_CONTENT: 204,
  FOUND: 302,
  BAD_REQUEST: 400,
  UNAUTHORIZED: 401,
  NO_FOUND: 404,
  INTERNAL_SERVER_ERROR: 500,
};

const HttpMethod = {
  POST: "post",
  GET: "get",
  PUT: "put",
  DELETE: "delete",
};

const HeaderKeys = {
  AUTH: "Authorization",
};

var headers = {};

var axiosConfig = {
  headers: headers,
  withCredentials: true,
};
function clearAuthHeader() {
  headers[HeaderKeys.AUTH] = "";
}
function errHandler(statusCode) {
  switch (statusCode) {
    case HttpStatusCode.BAD_REQUEST:
      alert("input error");
      break;
    case HttpStatusCode.UNAUTHORIZED:
      alert("驗證過期，請重新登入");
      location.replace("/login");
      break;
    case HttpStatusCode.NO_FOUND:
      alert("Not found");
      break;
  }
}
