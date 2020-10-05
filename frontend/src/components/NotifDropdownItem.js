import React, { Component } from 'react';
import { connect } from '../api';
// import Notifications from './Notifications';

class NotifDropdownItem extends Component {
    constructor(props) {
        super(props);
        // let temp = JSON.parse(this.props.notification);
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
            console.log(JSON.parse(this.state.msgNotif[0].data));
        })
    }

    render() {
        const message = this.state.msgNotif.map(msg => msg.data);
        // console.log(message);

        return (
            <div className="notif_dropdown_item">
                <div className="avatar" />
                <div className="content">
                    {/* <p>Id: </p> */}
                    <p>Message: {message}</p>
                </div>
            </div>
        );
    };
}

export default NotifDropdownItem;


// export default function NotifDropdownItem(props) {
//     return (
//         <div className="notif_dropdown_item">
//             <div className="avatar" />
//             <div className="content">
//                 <p>Id: {props.id}</p>
//                 <p>Message: {props.msg}</p>
//             </div>
//         </div>
//     )
// }