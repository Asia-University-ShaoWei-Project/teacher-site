import React, { useState, useEffect } from "react";

// import Paper from "@material-ui/core/Paper";
import Container from "@material-ui/core/Container";

import getGithubAPI, { fake } from "./Requests";
import axios from "axios";

import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import FolderIcon from "@material-ui/icons/Folder";
import DescriptionIcon from "@material-ui/icons/Description";

// onClick={() =>
//   props.setBreadcrumbItems([
//     {
//       label: val.title,
//       icon: val.icon,
//     },
//   ])
// }
const getAPI = (url, setItems) => {
  const data = [];
  console.log(url);
  axios
    .get(url)
    // .get(props.url)
    .then((response) => {
      console.log("axios success");
      if (response.data) {
        console.log("response data exist");
        response.data.map((val) =>
          data.push({
            title: val.name,
            isFolder: val.type === "dir" ? true : false,
            url: val.type === "dir" ? val.url : val.download_url,
          })
        );
      } else {
        console.log("no data");
      }
    })
    .catch(function (error) {
      console.log("get api catch");
    })
    .then(function () {
      console.log("complete");
      console.log(data);
      setItems(data);
    });
};

// Item Component
function Item(props) {
  console.log("Create Item");
  const updateBreadcrumb = (attr) => {
    if (props.breadcrumbItems) {
      console.log("更新導航列");

      props.setBreadcrumbItems([
        ...props.breadcrumbItems,
        {
          label: attr.title,
          icon: attr.isFolder ? <FolderIcon /> : <DescriptionIcon />,
        },
      ]);
    } else {
      console.log("無法更新導航列");
      console.log(props.breadcrumbItems);
      console.log("---------");
    }
  };
  return (
    <ListItem
      button
      onClick={() => {
        updateBreadcrumb({ title: props.title, isFolder: props.isFolder });
        props.setURL(props.url);
      }}
    >
      <ListItemIcon>
        {/* if this item is Folder */}
        {props.isFolder ? <FolderIcon /> : <DescriptionIcon />}
      </ListItemIcon>
      <ListItemText
        primary={props.title}
        // secondary={secondary ? "Secondary text" : null}
      />
    </ListItem>
  );
}

// Item List Component
export default function ItemList(props) {
  console.log("---------create ItemList------------");

  const [url, setURL] = useState(props.url);
  const [items, setItems] = useState([]);
  useEffect(() => {
    console.log("item list effect");
    getAPI(url, setItems);
  }, [url]);

  return (
    <Container maxWidth="sm">
      <button
        onClick={() => {
          console.log(props.breadcrumbItems);
        }}
      >
        click
      </button>
      <List>
        {items ? (
          items.map((val, index) => (
            <Item
              key={index}
              title={val.title}
              isFolder={val.isFolder}
              url={val.url}
              setURL={setURL}
              breadcrumbItems={props.breadcrumbItems}
              setBreadcrumbItems={props.setBreadcrumbItems}
            />
          ))
        ) : (
          <ListItem>
            <ListItemText
              primary="None"
              // secondary={secondary ? "Secondary text" : null}
            />
          </ListItem>
        )}
      </List>
    </Container>
  );
}
