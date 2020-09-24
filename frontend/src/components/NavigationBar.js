import React from 'react';
import '../App.css';
import NotifDropdown from './NotifDropdown';

export default function NavigationBar() {
    return (
        <div className="navbar">
            <div className="navbar_title">               
                <h1>Celerates App</h1>                        
            </div>
            <div className="navbar_notification">
                <NotifDropdown />
            </div>
        </div>
    )
}