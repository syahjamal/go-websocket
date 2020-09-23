import React from 'react';
import '../App.css';
import SimpleModal from './SimpleModal';

export default function NavigationBar() {
    return (
        <div className="navbar">
            <div className="navbar_title">
                <h1>Celerates App</h1>
            </div>
            <div className="navbar-notification">
                <SimpleModal />
            </div>
        </div>
    )
}