import React, { Component } from "react";
import logo from './logo.svg';
import './App.css';
import { connect, sendMsg } from "./api";
import Header from "./components/Header/Header.jsx";
import ChatHistory from "./components/ChatHistory/ChatHistory.jsx"

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatHistory : []
    }
  }

  componentWillMount(){
    connect((msg) => {
      console.log("New message")
      this.setState(preState => ({
        chatHistory : [...this.state.chatHistory, msg]
      }))
      console.log(this.state)
    })
  }

  send() {
    console.log("HEEEEEEEEEEE");
    sendMsg("HEWADasdasdas");
  }

  render() {
    return (
      <div className="App">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <button onClick={this.send}>Hit</button>
      </div>
    );
  }
}

export default App;
