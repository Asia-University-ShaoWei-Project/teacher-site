// import { makeStyles } from "@material-ui/core/styles";
import Box from "@material-ui/core/Box";
import IconButton from "@material-ui/core/IconButton";

import GitHubIcon from "@material-ui/icons/GitHub";
import FacebookIcon from "@material-ui/icons/Facebook";
import AssignmentIcon from "@material-ui/icons/Assignment";

// const useStyles = makeStyles((theme) => ({
//   root: {
//     display: "flex",
//     margin: theme.spacing(2),
//     "& > *": {
//       margin: theme.spacing(1),
//     },
//   },
// }));

export default function Header(props) {
  return (
    <Box display="flex" flexDirection="row" justifyContent="flex-end" m={1}>
      <Box>
        <IconButton
          color="primary"
          aria-label="upload picture"
          component="span"
        >
          <GitHubIcon />
        </IconButton>
      </Box>
      <Box>
        <IconButton
          color="primary"
          aria-label="upload picture"
          component="span"
        >
          <FacebookIcon />
        </IconButton>
      </Box>
      {/* <Box p={1}>
      <IconButton
          color="primary"
          aria-label="upload picture"
          component="span"
        >
          <GitHubIcon />
        </IconButton>
      </Box> */}
    </Box>
  );
}
