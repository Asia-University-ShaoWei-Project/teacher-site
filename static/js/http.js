const HttpStatusCode = {
  OK: 200,
  CREATED: 201,
  // Accepted             :202,
  // NonAuthoritativeInfo : 203,
  NO_CONTENT: 204,
  // ResetContent         : 205,
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

var axiosConfig = { headers: headers };
function clearAuthHeader() {
  headers[HeaderKeys.AUTH] = "";
}
