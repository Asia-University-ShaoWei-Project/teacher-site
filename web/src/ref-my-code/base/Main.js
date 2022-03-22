import React, { useState } from "react";
import clsx from "clsx";

import List from "@material-ui/core/List";
import Container from "@material-ui/core/Container";
import Drawer from "@material-ui/core/Drawer";
import Divider from "@material-ui/core/Divider";
import IconButton from "@material-ui/core/IconButton";
import Paper from "@material-ui/core/Paper";
import Box from "@material-ui/core/Box";

import ChevronLeftIcon from "@material-ui/icons/ChevronLeft";
import ListItems from "./SideBar";
import { BrowserRouter, Switch, Route } from "react-router-dom";

import Breadcrumb from "./Breadcrumbs";
import { router_config } from "./Router";
import Header from "./Header";

export default function Main(props) {
  const classes = props.classes;

  const [breadcrumbItems, setBreadcrumbItems] = useState([
    {
      label: router_config[0].title,
      icon: router_config[0].icon,
    },
  ]);

  // const updateBreadcrumb = (item) => {
  //   console.log("update bread");
  //   setBreadcrumbItems([
  //     {
  //       label: item.title,
  //       icon: item.icon,
  //     },
  //   ]);
  // };
  return (
    <BrowserRouter>
      <Drawer
        variant="permanent"
        classes={{
          paper: clsx(
            classes.drawerPaper,
            !props.drawerOpen && classes.drawerPaperClose
          ),
        }}
        open={props.drawerOpen}
      >
        <div className={classes.toolbarIcon}>
          <IconButton onClick={props.handleDrawerOpen}>
            <ChevronLeftIcon />
          </IconButton>
        </div>
        <Divider />
        <List>
          <ListItems setBreadcrumbItems={setBreadcrumbItems} />
        </List>
        {/* <Divider />
        <List>{secondaryListItems}</List> */}
      </Drawer>
      <main className={classes.content}>
        <Header />
        <Container maxWidth="lg" className={classes.container}>
          <Paper>
            <Box p={1}>
              <Breadcrumb breadcrumbItems={breadcrumbItems} />
            </Box>
          </Paper>

          <Paper>
            <Box p={1} m={3}>
              <Switch>
                {router_config.map((val, index) => {
                  return (
                    <Route key={index} path={val.path}>
                      {val.page(breadcrumbItems, setBreadcrumbItems)}
                    </Route>
                  );
                })}
                <Route path="*">
                  <h1>404</h1>
                </Route>
              </Switch>
            </Box>
          </Paper>
        </Container>
      </main>
    </BrowserRouter>
  );
}
