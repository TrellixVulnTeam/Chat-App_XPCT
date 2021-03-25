import React, { Component } from "react"
import "./App.css"
import { connect, sendMessage } from "./api"
import Header from "./components/Header"
import ChatHistory from "./components/ChatHistroy"
import ChatInput from "./components/ChatInput"

class App extends Component {
  constructor(props) {
    super(props)
    this.state = {
      chatHistory: []
    }
    // this.componentDidMount = this.componentDidMount.bind(this)
  }

  componentDidMount() {
    connect((msg) => {
      console.log("New Message")
      this.setState((prevState) => ({
        chatHistory: [...this.state.chatHistory, msg],
      }))
      console.log(this.state.chatHistory)
    })
  }

  send(event) {
    if (event.keyCode === 13) {
      sendMessage(event.target.value)
      event.target.value = ""
    }
  }

  render() {
    return (
      <div className="App">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <ChatInput send={this.send} />
      </div>
    )
  }
}

export default App