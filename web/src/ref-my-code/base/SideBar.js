import React from "react";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import ListSubheader from "@material-ui/core/ListSubheader";

import AssignmentIcon from "@material-ui/icons/Assignment";
import { router_config } from "./Router";
import { Link } from "react-router-dom";
// import { routerPath } from "./Router";

export default function ListItems(props) {
  return router_config.map((val, index) => {
    return (
      <Link
        key={index}
        to={val.path}
        style={{ textDecoration: "none", color: "black" }}
        onClick={() =>
          props.setBreadcrumbItems([
            {
              label: val.title,
              icon: val.icon,
            },
          ])
        }
      >
        <ListItem button>
          <ListItemIcon>{val.icon}</ListItemIcon>
          <ListItemText primary={val.title} />
        </ListItem>
      </Link>
    );
  });
}

export const secondaryListItems = (
  <div>
    <ListSubheader inset>Saved reports</ListSubheader>
    <ListItem button>
      <ListItemIcon>
        <AssignmentIcon />
      </ListItemIcon>
      <ListItemText primary="Current month" />
    </ListItem>
    <ListItem button>
      <ListItemIcon>
        <AssignmentIcon />
      </ListItemIcon>
      <ListItemText primary="Last quarter" />
    </ListItem>
    <ListItem button>
      <ListItemIcon>
        <AssignmentIcon />
      </ListItemIcon>
      <ListItemText primary="Year-end sale" />
    </ListItem>
  </div>
);
