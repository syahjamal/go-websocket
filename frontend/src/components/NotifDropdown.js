import React from 'react';
import '../App.css';
import Button from '@material-ui/core/Button';
import NotifDropdownItem from './NotifDropdownItem';

export default function NotifDropdown() {
    const handleClick = () => {
        document.getElementById("myNotifDropdown").classList.toggle("show");
    }

    window.onclick = (event) => {
        if (!event.target.matches('.notif_dropbtn')) {
            var dropdowns = document.getElementsByClassName("notif_dropdown_content");
            var i;
            for (i = 0; i < dropdowns.length; i++) {
                var openDropdown = dropdowns[i];
                if (openDropdown.classList.contains('show')) {
                    openDropdown.classList.remove('show');
                }
            }
        }
    }
    
    return (
        <div className="notif_dropdown">
            <div className="notif_dropdown_icon">
                <i onClick={handleClick} className="notif_dropbtn fas fa-bell"></i>
                <span className="badge"></span>
            </div>                       
            
            <div id="myNotifDropdown" className="notif_dropdown_content">
                <div className="notif_dropdown_header">
                    <h2>Notifications</h2>
                </div>
                <div className="notif_dropdown_body">
                   <NotifDropdownItem id="1" msg="Hello there, greet from Bekasi :)"/>        
                   <NotifDropdownItem id="2" msg="Hello there, greet from Bandung :)"/>        
                   <NotifDropdownItem id="3" msg="Hello there, greet from Bogor :)"/>        
                </div>
                <div className="notif_dropdown_footer">
                    <Button size="small" variant="outlined" color="primary">See All</Button>
                </div>            
            </div>
        </div>
    )
}
