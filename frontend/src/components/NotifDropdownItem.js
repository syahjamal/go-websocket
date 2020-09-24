import React from 'react'

export default function NotifDropdownItem(props) {
    return (
        <div className="notif_dropdown_item">
            <div className="avatar" />
            <div className="content">
                <p>Id: {props.id}</p>
                <p>Message: {props.msg}</p>
            </div>
        </div>
    )
}
