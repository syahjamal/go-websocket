import React, { Component } from 'react';

class Notifications extends Component {
    constructor(props) {
        super(props);
        let temp = JSON.parse(this.props.notifications);
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
