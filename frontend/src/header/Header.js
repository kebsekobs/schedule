import React from "react";
import style from "./header.module.css";
import { headerItems } from "./headerItems";
import { NavLink } from "react-router-dom";
import { useGenerateMutation } from "./useGenerateMutation";

const Header = () => {
  const generateMutation = useGenerateMutation();
  return (
    <div className={style["header"]}>
      <div className={style["header-items"]}>
        {headerItems.map((item) => (
          <NavLink
            className={({ isActive }) =>
              isActive ? style["active"] : style["header-item"]
            }
            to={item.path}
            key={item.name}
          >
            {item.name}
          </NavLink>
        ))}
        <h2 className={style['header-item']} style={{marginTop: '7px'}} onClick={() => generateMutation()} >
          Сгенерировать
          </h2>
      </div>
    </div>
  );
};

export default Header;
