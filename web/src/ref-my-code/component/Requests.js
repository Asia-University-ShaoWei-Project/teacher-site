import axios from "axios";

const url = "https://api.github.com/repos/asiashaowei/PublicNote/contents/";

export default function getGithubAPI(props) {
  console.log("into the get api");
  axios
    .get(url)
    // .get(props.url)
    .then((response) => {
      console.log("success");
      if (response.data) {
        console.log("has data");
        const data = [];
        response.data.map((val) =>
          data.push({
            title: val.name,
            isFolder: val.type === "dir" ? true : false,
            url: val.type === "dir" ? val.url : val.download_url,
          })
        );
        console.log(data);
        return data;
      } else {
        console.log("not array");
        return [];
      }
    })
    .catch(function (error) {
      // handle error
      console.log("error");
      return [];
    })
    .then(function () {
      console.log("complete");
      // always executed
    });
}
export function fake() {
  console.log("into fake");
  const json = [
    {
      name: "Github_init_push.png",
      path: "Program_Language/Python/Github_init_push.png",
      sha: "61d64775c7992ac1b32f73699a6c79b71c41b5a9",
      size: 670508,
      url:
        "https://api.github.com/repos/asiashaowei/PublicNote/contents/Program_Language/Python/Github_init_push.png?ref=main",
      html_url:
        "https://github.com/asiashaowei/PublicNote/blob/main/Program_Language/Python/Github_init_push.png",
      git_url:
        "https://api.github.com/repos/asiashaowei/PublicNote/git/blobs/61d64775c7992ac1b32f73699a6c79b71c41b5a9",
      download_url:
        "https://raw.githubusercontent.com/asiashaowei/PublicNote/main/Program_Language/Python/Github_init_push.png",
      type: "file",
      _links: {
        self:
          "https://api.github.com/repos/asiashaowei/PublicNote/contents/Program_Language/Python/Github_init_push.png?ref=main",
        git:
          "https://api.github.com/repos/asiashaowei/PublicNote/git/blobs/61d64775c7992ac1b32f73699a6c79b71c41b5a9",
        html:
          "https://github.com/asiashaowei/PublicNote/blob/main/Program_Language/Python/Github_init_push.png",
      },
    },
    {
      name: "Hook.html",
      path: "Program_Language/Python/Hook.html",
      sha: "02d47eb2b79767af625007621112735db8b7ffc4",
      size: 25168,
      url:
        "https://api.github.com/repos/asiashaowei/PublicNote/contents/Program_Language/Python/Hook.html?ref=main",
      html_url:
        "https://github.com/asiashaowei/PublicNote/blob/main/Program_Language/Python/Hook.html",
      git_url:
        "https://api.github.com/repos/asiashaowei/PublicNote/git/blobs/02d47eb2b79767af625007621112735db8b7ffc4",
      download_url:
        "https://raw.githubusercontent.com/asiashaowei/PublicNote/main/Program_Language/Python/Hook.html",
      type: "file",
      _links: {
        self:
          "https://api.github.com/repos/asiashaowei/PublicNote/contents/Program_Language/Python/Hook.html?ref=main",
        git:
          "https://api.github.com/repos/asiashaowei/PublicNote/git/blobs/02d47eb2b79767af625007621112735db8b7ffc4",
        html:
          "https://github.com/asiashaowei/PublicNote/blob/main/Program_Language/Python/Hook.html",
      },
    },
    {
      name: "hi.html",
      path: "Program_Language/Python/hi.html",
      sha: "c52a578633c7d0be99f48a549e9f342d5a7fcce2",
      size: 13784,
      url:
        "https://api.github.com/repos/asiashaowei/PublicNote/contents/Program_Language/Python/hi.html?ref=main",
      html_url:
        "https://github.com/asiashaowei/PublicNote/blob/main/Program_Language/Python/hi.html",
      git_url:
        "https://api.github.com/repos/asiashaowei/PublicNote/git/blobs/c52a578633c7d0be99f48a549e9f342d5a7fcce2",
      download_url:
        "https://raw.githubusercontent.com/asiashaowei/PublicNote/main/Program_Language/Python/hi.html",
      type: "file",
      _links: {
        self:
          "https://api.github.com/repos/asiashaowei/PublicNote/contents/Program_Language/Python/hi.html?ref=main",
        git:
          "https://api.github.com/repos/asiashaowei/PublicNote/git/blobs/c52a578633c7d0be99f48a549e9f342d5a7fcce2",
        html:
          "https://github.com/asiashaowei/PublicNote/blob/main/Program_Language/Python/hi.html",
      },
    },
    {
      name: "hi.md",
      path: "Program_Language/Python/hi.md",
      sha: "b4a7e62657eb034ea935a9b2e91f44a6201d182d",
      size: 299,
      url:
        "https://api.github.com/repos/asiashaowei/PublicNote/contents/Program_Language/Python/hi.md?ref=main",
      html_url:
        "https://github.com/asiashaowei/PublicNote/blob/main/Program_Language/Python/hi.md",
      git_url:
        "https://api.github.com/repos/asiashaowei/PublicNote/git/blobs/b4a7e62657eb034ea935a9b2e91f44a6201d182d",
      download_url:
        "https://raw.githubusercontent.com/asiashaowei/PublicNote/main/Program_Language/Python/hi.md",
      type: "file",
      _links: {
        self:
          "https://api.github.com/repos/asiashaowei/PublicNote/contents/Program_Language/Python/hi.md?ref=main",
        git:
          "https://api.github.com/repos/asiashaowei/PublicNote/git/blobs/b4a7e62657eb034ea935a9b2e91f44a6201d182d",
        html:
          "https://github.com/asiashaowei/PublicNote/blob/main/Program_Language/Python/hi.md",
      },
    },
    {
      name: "other",
      path: "Program_Language/Python/other",
      sha: "b7b1747aec2b74a4798fdcb9223b2013740f1b0d",
      size: 0,
      url:
        "https://api.github.com/repos/asiashaowei/PublicNote/contents/Program_Language/Python/other?ref=main",
      html_url:
        "https://github.com/asiashaowei/PublicNote/tree/main/Program_Language/Python/other",
      git_url:
        "https://api.github.com/repos/asiashaowei/PublicNote/git/trees/b7b1747aec2b74a4798fdcb9223b2013740f1b0d",
      download_url: null,
      type: "dir",
      _links: {
        self:
          "https://api.github.com/repos/asiashaowei/PublicNote/contents/Program_Language/Python/other?ref=main",
        git:
          "https://api.github.com/repos/asiashaowei/PublicNote/git/trees/b7b1747aec2b74a4798fdcb9223b2013740f1b0d",
        html:
          "https://github.com/asiashaowei/PublicNote/tree/main/Program_Language/Python/other",
      },
    },
  ];
  console.log(typeof json === Array.isArray);
  if (json) {
    const data = [];
    json.map((val) =>
      data.push({
        title: val.name,
        isFolder: val.type === "dir" ? true : false,
        url: val.type === "dir" ? val.url : val.download_url,
      })
    );
    console.log(data);
    return data;
  } else {
    console.log("not array");
    return [];
  }
}
// async function getUser() {
//   try {
//     const response = await axios.get("/user?ID=12345");
//     console.log(response);
//   } catch (error) {
//     console.error(error);
//   }
// }

// axios
//   .get("/user", {
//     params: {
//       ID: 12345,
//     },
//   })
//   .then(function (response) {
//     console.log(response);
//   })
//   .catch(function (error) {
//     console.log(error);
//   })
//   .then(function () {
//     // always executed
//   });

// axios
//   .post("/user", {
//     firstName: "Fred",
//     lastName: "Flintstone",
//   })
//   .then(function (response) {
//     console.log(response);
//   })
//   .catch(function (error) {
//     console.log(error);
//   });

// function getUserAccount() {
//   return axios.get("/user/12345");
// }

// function getUserPermissions() {
//   return axios.get("/user/12345/permissions");
// }

// Promise.all([getUserAccount(), getUserPermissions()]).then(function (results) {
//   const acct = results[0];
//   const perm = results[1];
// });
