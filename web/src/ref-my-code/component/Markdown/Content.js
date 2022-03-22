import React from "react";
import ReactMarkdown from "react-markdown";
import axios from "axios";
import { CircularProgress } from "@material-ui/core";
import { Paper } from "@material-ui/core";
import CodeBlock from "./CodeBlock";
const markdown_url = [
  "https://raw.githubusercontent.com/asiashaowei/PublicNote/main/Program_Language/Python/hi.md",
  "https://raw.githubusercontent.com/asiashaowei/PublicNote/main/Frontend/React/base.md",
];

const renderers = {
  //This custom renderer changes how images are rendered
  //we use it to constrain the max width of an image to its container
  code: CodeBlock,
  image: ({ alt, src, title }) => (
    <Paper variant="outlined" square>
      <img alt={alt} src={src} title={title} style={{ maxWidth: 200 }} />
    </Paper>
  ),
  // heading: (props) => <Title>{props.children}</Title>,
};

class MarkdownContent extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      text: "",
    };
    this.getData = this.getData.bind(this);
  }
  componentDidMount() {
    this.getData();
  }
  getData() {
    axios
      .get(markdown_url[1])
      .then((response) => {
        console.log("success");
        // handle success
        this.setState({
          text: response.data,
        });
      })
      .catch(function (error) {
        // handle error
        console.log("error");
      })
      .then(function () {
        console.log("complete");
        // always executed
      });
  }
  render() {
    return (
      <div>
        {this.state.text === "" ? (
          <CircularProgress />
        ) : (
          <ReactMarkdown renderers={renderers}>{this.state.text}</ReactMarkdown>
        )}
      </div>
    );
  }
}

export default MarkdownContent;
