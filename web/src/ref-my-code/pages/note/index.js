import MarkdownContent from "../../component/Markdown/Content";
export default function NotePage() {
  return (
    <div>
      this is note page
      <MarkdownContent />
    </div>
  );
}

// import React from "react";

// const note_url = "https://api.github.com/repos/asiashaowei/PublicNote/contents";
// // https://api.github.com/repos/asiashaowei/PublicNote/contents/Program_Language
// class NotePage extends React.Component {
//   constructor(props) {
//     super(props);
//     this.state = {
//       key1: true,
//       date: new Date(),
//     };
//     // 為了讓 `this` 能在 callback 中被使用，這裡的綁定是必要的：
//     // this.myFunction = this.defFunction.bind(this);
//   }

//   // componentDidMount()會在 component 被 render 到 DOM 之後才會執行
//   componentDidMount() {
//     // 雖然 this.props 是由 React 本身設定的，而且 this.state 具有特殊的意義，如果需要儲存一些不相關於資料流的內容（像是 timer ID）， 可以自由的手動加入。
//     this.timerID = setInterval(() => this.tick(), 1000);
//   }
//   // 如果 component 從 DOM 被移除了，React 會呼叫 componentWillUnmount() 生命週期方法.
//   componentWillUnmount() {
//     clearInterval(this.timerID);
//   }

//   tick() {
//     this.setState({
//       date: new Date(),
//     });
//   }
//   // myFunction(){}

//   render() {
//     return (
//       <div>
//         <MyClassName warn={this.state.key1} />
//         <button onClick={this.defFunction}>
//           {this.state.key1 ? "Hide" : "Show"}
//         </button>
//       </div>
//     );
//   }
// }

// // name
// // type : dir  or file
