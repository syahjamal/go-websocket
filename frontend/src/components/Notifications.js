import React, { Component } from 'react';

class Notifications extends Component {
    constructor(props) {
        super(props);
        let temp = JSON.parse(this.props.notification);
        this.state = {
            notification: temp
        }
    }

    render() {
        return (
            <div className='Notifications'>
                {this.state.notification.message}
            </div>
        );
    };

}

export default Notifications;
