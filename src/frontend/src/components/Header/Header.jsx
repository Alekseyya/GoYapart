import React, { Component } from "react";
import "./header.css";

class Header extends Component {
    constructor(props) {
        super(props)
    }
    render() {
        return (
            <div className="header">
                <h2>Realtime Chat App</h2>
            </div>
        )
    }
}

export default Header;