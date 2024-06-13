import React from "react";
import style from "./header.module.css";
import { headerItems } from "./headerItems";
import { NavLink } from "react-router-dom";
import { useGenerateMutation } from "./useGenerateMutation";
import {useParseMutation} from "./useParseDataMutation";

const Header = () => {
  const generateMutation = useGenerateMutation();
  const parseDataMutation = useParseMutation();
 function generateMutationHandle () {
  generateMutation.mutateAsync()
 }
    function parseDataMutationHandle () {
        parseDataMutation.mutateAsync()
    }
  return (
    <div className={style["header"]}>
      <div className={style["header-items"]}>
          <h2 className={style['header-item']} style={{marginTop: '7px'}} onClick ={() => parseDataMutationHandle()} >
              Обновление БД
          </h2>
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
        <h2 className={style['header-item']} style={{marginTop: '7px'}} onClick ={() => generateMutationHandle()} >
          Сгенерировать
          </h2>
      </div>
    </div>
  );
};

export default Header;
