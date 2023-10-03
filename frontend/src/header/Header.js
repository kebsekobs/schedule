import React from 'react';
import './header.css'
import {headerItems} from "./headerItems";
import { NavLink} from "react-router-dom";
const Header = () => {
    return (
        <div className={'header'}>
            <div className={'header-items'}>
                {headerItems.map(item =>
                    <NavLink  className={'header-item'} to={item.path} key={item.name}>
                        {item.name}
                    </NavLink>
                )}
            </div>
        </div>

    );
};

export default Header;