import React from 'react';
import '../App.css';

export default function GetCv() {
    function handleClick(e) {
        e.preventDefault();
        console.log('Clicked');
    }

    return (      
        <a href="#" onClick={handleClick}>
            <button>Get CV</button>
        </a>
    )
}