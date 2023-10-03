import React from 'react';
import './header.css'
const Header = () => {
    return (
        <div className={'header'}>
            <div className={'header-items'}>
                <div className={'header-item'}>
                    Группы
                </div>
                <div className={'header-item'}>
                    Аудитории
                </div>
                <div className={'header-item'}>
                    Преподаватели
                </div>
                <div className={'header-item'}>
                    Предметы
                </div>
                <div className={'header-item'}>
                    Сгенерировать
                </div>
            </div>
        </div>

    );
};

export default Header;