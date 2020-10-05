import React, { Component } from 'react';

class Notifications extends Component {
    constructor(props) {
        super(props);
<<<<<<< HEAD
        let temp = JSON.parse(this.props.notifications);
=======
        let temp = JSON.parse(this.props.notification);
>>>>>>> ca5d2deed5da57bd2cf5e270e0158dc360ee032c
        this.state = {
            notification: temp
        }
    }

    render() {
        return (
            <div className='Notifications'>
<<<<<<< HEAD
                {this.state.notifications}
=======
                {this.state.notification.message}
>>>>>>> ca5d2deed5da57bd2cf5e270e0158dc360ee032c
            </div>
        );
    };

}

export default Notifications;
