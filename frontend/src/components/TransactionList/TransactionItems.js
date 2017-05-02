import React from 'react';
import './Transaction.css'

export default (props) => {
    const {type, date, operation} = props.transaction;
    let prefix = '';
    let style = {color: 'black'};

    if (type === "Withdrawal") {
        prefix = "-";
    }

    return (
        <tr>
            <th >{type}</th>
            <th >{date}</th>
            <th style={style}>{prefix} {operation}</th>
        </tr>
    );
};
