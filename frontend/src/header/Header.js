import React from 'react';
import style from './header.module.css'
import {headerItems} from "./headerItems";
import { NavLink} from "react-router-dom";
const Header = () => {
    return (
        <div className={style['header']}>
            <div className={style['header-items']}>
                {headerItems.map(item =>
                    <NavLink  className={style['header-item']} to={item.path} key={item.name}>
                        {item.name}
                    </NavLink>
                )}
            </div>
        </div>

    );
};

export default Header;