import React from "react";
import { emphasize, withStyles } from "@material-ui/core/styles";
import Breadcrumbs from "@material-ui/core/Breadcrumbs";
import Chip from "@material-ui/core/Chip";

export const StyledBreadcrumb = withStyles((theme) => ({
  root: {
    backgroundColor: theme.palette.white,
    height: theme.spacing(3),
    color: theme.palette.grey[800],
    fontWeight: theme.typography.fontWeightRegular,
    "&:hover, &:focus": {
      backgroundColor: theme.palette.grey[300],
    },
    "&:active": {
      boxShadow: theme.shadows[1],
      backgroundColor: emphasize(theme.palette.grey[300], 0.12),
    },
  },
}))(Chip); // TypeScript only: need a type cast here because https://github.com/Microsoft/TypeScript/issues/26591

export default function Breadcrumb(props) {
  console.log("breadcrumb");
  return (
    <Breadcrumbs aria-label="breadcrumb">
      {props.breadcrumbItems &&
        props.breadcrumbItems.map((val, index) => (
          <StyledBreadcrumb
            key={index}
            component="a"
            href="#"
            label={val.label}
            icon={val.icon}
            // onClick={handleClick}
          />
        ))}
    </Breadcrumbs>
  );
}
