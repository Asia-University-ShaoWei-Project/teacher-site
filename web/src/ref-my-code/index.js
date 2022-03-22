import React, { StrictMode, useState, useEffect } from "react";
import ReactDOM from "react-dom";
import CssBaseline from "@material-ui/core/CssBaseline";
// import Main from "./base/Main";
import { makeStyles } from "@material-ui/core/styles";
import { createStore } from "redux";
import { connect, Provider } from "react-redux";
const rootElement = document.getElementById("root");
const useBaseStyles = makeStyles((theme) => {
  const drawerWidth = 240;
  return {
    root: {
      display: "flex",
    },
    toolbarIcon: {
      display: "flex",
      alignItems: "center",
      justifyContent: "flex-end",
      padding: "0 8px",
      ...theme.mixins.toolbar,
    },

    title: {
      flexGrow: 1,
    },
    drawerPaper: {
      position: "relative",
      whiteSpace: "nowrap",
      width: drawerWidth,
      transition: theme.transitions.create("width", {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.enteringScreen,
      }),
    },
    drawerPaperClose: {
      overflowX: "hidden",
      transition: theme.transitions.create("width", {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
      width: theme.spacing(7),
      [theme.breakpoints.up("sm")]: {
        width: theme.spacing(9),
      },
    },
    content: {
      flexGrow: 1,
      height: "100vh",
      overflow: "auto",
    },
    container: {
      paddingTop: theme.spacing(2),
      paddingBottom: theme.spacing(4),
    },
  };
});

const initState = {
  name: "Jack",
};
const reducer = (state = initState, action) => {
  switch (action.type) {
    default:
      return state;
  }
};
const store = createStore(reducer);
const mapStateToProps = (state) => ({
  name: state.name,
});
class ConnectTitle extends React.Component {
  render() {
    return <h1>Hello!{this.props.name}</h1>;
  }
}
const Title = connect(mapStateToProps)(ConnectTitle);

function Index() {
  console.log("Index");
  const classes = useBaseStyles();
  const [drawerOpen, setDrawerOpen] = useState(true);

  // const handleDrawerOpen = () => {
  //   setDrawerOpen(!drawerOpen);
  // };

  return (
    <Provider store={store} className={classes.root}>
      <CssBaseline />
      <Title></Title>

      {/* <Main
        classes={classes}
        drawerOpen={drawerOpen}
        handleDrawerOpen={handleDrawerOpen}
        // breadcrumbItems={breadcrumbItems}
        // updateBreadcrumb={updateBreadcrumb}
      /> */}
    </Provider>
  );
}
ReactDOM.render(
  <StrictMode>
    <Index />
  </StrictMode>,
  rootElement
);
