import Link from "@material-ui/core/Link";
import Typography from "@material-ui/core/Typography";
// import Box from "@material-ui/core/Box";

export function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {"Copyright © "}
      <Link color="inherit" href="https://material-ui.com/">
        Your Website
      </Link>{" "}
      {new Date().getFullYear()}
      {"."}
    </Typography>
  );
}

// {
/* <Box pt={4}>
  <Copyright />
</Box>; */
// }
