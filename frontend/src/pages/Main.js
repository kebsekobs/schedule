import React from "react";
import "../App.css";
import SelectComp from "../components/select";
import Input from "../components/input";
import Button from "../components/button";

const testList = [
  {
    id: "123(а)",
    capacity: 164,
  },
  {
    id: "432",
    capacity: 164,
  },
  {
    id: "753",
    capacity: 32,
  },
  {
    id: "554",
    capacity: 15,
  },
];

const Main = () => {
  return (
    <div className={"page"}>
      Главная
      <Input />
      <SelectComp
        list={testList}
        placeholder="Выберите аудиторию"
        name="Аудитории"
      />
    </div>
  );
};

export default Main;
