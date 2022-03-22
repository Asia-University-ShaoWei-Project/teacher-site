import "./styles.css";
import axios from "axios";
import React from "react";
import { Button } from "@material-ui/core";
// import Dashboard from "./pages/Index";
function YY() {
  return <Button color="primary">Hello World</Button>;
}
const url = "https://jsonplaceholder.typicode.com/users";
class People extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      id: props.id,
      name: props.name,
      username: props.username,
      email: props.email
    };
  }
  render() {
    const state = this.state;
    console.log("___");

    console.log(state.id);
    return (
      <div>
        <p>id: {state.id}</p>
        <p>name: {state.name}</p>
        <p>username: {state.username}</p>
        <p>email: {state.email}</p>
      </div>
    );
  }
}
class Te extends React.Component {
  constructor(props) {
    super(props);
    this.t1_func = this.t1_func.bind(this);
    this.getPeopleJson = this.getPeopleJson.bind(this);
    this.state = {
      people: [],
      check: true
    };
  }
  t1_func() {
    console.log("check value ");
    this.setState((state) => ({ check: !state.check }));
  }
  getPeopleJson() {
    axios
      .get(url)
      .then((response) => {
        // handle success
        console.log("success request");
        this.setState((state) => ({
          people: response.data.map((item) => (
            <People
              key={item.id}
              id={item.id}
              name={item.name}
              username={item.username}
              email={item.email}
            />
          ))
        }));
      })
      .catch(function (error) {
        // handle error
        console.log(error);
      })
      .then(function () {
        console.log("complete");
        // always executed
      });
  }
  render() {
    return (
      <div>
        state = {this.state.check ? "true" : "false"}
        <p></p>
        <button onClick={this.t1_func}>click t1</button>
        <p></p>
        <h3> Peoples </h3>
        count: <p>{this.state.people.length}</p>
        {this.state.people}
        <p></p>
        <button onClick={() => this.getPeopleJson()}>click</button>
      </div>
    );
  }
}
export default function App() {
  return (
    // <test />

    <div className="App">
      {/* <Dashboard /> */}
      <YY />
      <Te />
      {/* <h1>Hello1 CodeSandbox</h1> */}
      {/* <h2>Start editing to see some magic happen!</h2> */}
    </div>
  );
}
