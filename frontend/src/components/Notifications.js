import React, { Component } from 'react';
import { sendMsg } from '../api';

class Notifications extends Component {
    constructor(props) {
        super(props);
        let temp = sendMsg(JSON.parse(this.props.notifications));
        this.state = {
            notifications: temp
        }
    }

    render() {
        return (
            <div className='Notifications'>
                {this.state.notifications.message}
            </div>
        );
    };

}

export default Notifications;
