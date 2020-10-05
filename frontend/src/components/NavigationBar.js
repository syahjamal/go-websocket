import React from 'react';
import '../App.css';
import NotifDropdown from './NotifDropdown';

export default function NavigationBar() {
    return (
        <div className="navbar">
            <div className="navbar_title">               
                <a href="/" target="_blank"><h1>Celerates App</h1></a>                       
            </div>
            <div className="navbar_notification">
                <NotifDropdown />
            </div>
        </div>
    )
}