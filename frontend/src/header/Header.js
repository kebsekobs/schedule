import React, { useState } from "react";
import style from "./header.module.css";
import { headerItems } from "./headerItems";
import { NavLink } from "react-router-dom";
import { useGenerateMutation } from "./useGenerateMutation";
import Button from "../components/button";

const Header = () => {
  const generateMutation = useGenerateMutation();
 function generateMutationHandle () {
  generateMutation.mutateAsync()
 }
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
        <Button className={style['header-item']} style={{marginTop: '7px'}} onClick ={() => generateMutationHandle()} >
          Сгенерировать
          </Button>
      </div>
    </div>
  );
};

export default Header;
