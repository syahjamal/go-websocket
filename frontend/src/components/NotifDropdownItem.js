import React, { Component } from 'react';
import { connect } from '../api';
// import Notifications from './Notifications';

class NotifDropdownItem extends Component {
    constructor(props) {
        super(props);
        this.state = {         
            msgNotif: []
        }
    }
    
    componentDidMount() {
        connect((msg) => {
            console.log("New Message")
            this.setState(prevState => ({
                msgNotif: [...prevState.msgNotif, msg]
            }))
            // console.log(JSON.parse(this.state.msgNotif[0].data));
        })
    }

    render() {
        const message = this.state.msgNotif.map(msg => msg.data);
        return (
            <div className="notif_dropdown_item">
                <div className="content">
                    {message}
                </div>
            </div>
        );
    };
}

export default NotifDropdownItem;