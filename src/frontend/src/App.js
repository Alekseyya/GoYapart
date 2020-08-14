import React, { Component } from "react";
import logo from './logo.svg';
import './App.css';
import { connect, sendMsg } from "./api";
import Header from "./components/Header/Header.jsx";

class App extends Component {
  constructor(props) {
    super(props);
    connect();
  }

  send() {
    console.log("HEEEEEEEEEEE");
    sendMsg("HEWADasdasdas");
  }

  render() {
    return (
      <div className="App">
        <Header />
        <button onClick={this.send}>Hit</button>
      </div>
    );
  }
}

export default App;
