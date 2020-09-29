import React, { Component } from 'react';
import Notifications from './Notifications';

class NotifDropdownItem extends Component {
    // constructor(props) {
    //     super(props);
    //     this.state = {
    //         id: 1,
    //         message: 'Hello there i am from Bekasi :)'
    //     }
    // }  

    render() {
        console.log(this.props.notifDropdownItem);
        const messages = this.props.notifDropdownItem.map(msg => <Notifications message={msg.data} /> );
        // console.log(messages)
        return (
            <div className="notif_dropdown_item">
                <div className="avatar" />
                <div className="content">
                    {messages}
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