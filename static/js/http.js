const HTTP_STATUS_CODE = {
  ok: 200,
  created: 201,
  // StatusAccepted             = 202,
  // StatusNonAuthoritativeInfo = 203,
  noContent: 204,
  // StatusResetContent         = 205,
  found: 302,
  badRequest: 400,
  unauthorized: 401,
  notFound: 404,
  statusInternalServerError: 500,
};

const HTTP_METHOD = {
  post: "post",
  get: "get",
  put: "put",
  delete: "delete",
};

const headerKeys = {
  auth: "Authorization",
};

var headers = {};

var axiosConfig = { headers: headers };
