import React, { Component } from "react";
import "./Message.scss"

class Message extends Component {
    constructor(props) {
        super(props)
        let temp = JSON.parse(this.props.message)
        console.log(temp)
        this.state = {
            message: temp,
        }
    }

    render() {
        return (
            <div>
                <p className="From">{this.state.message.From}</p>
                <p className="Message">{this.state.message.body}</p>
            </div>
        )
    }
}

export default Message