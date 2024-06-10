import React, { useState } from "react";
import { useGroupsQuery } from "./api/getGroupsQuery";
import CoursesTable from "./table/use-table";
import AddGroupModal from "./modals/AddModal";
import Button from "../../components/button";

const Groups = () => {
  const getGroupsQuery = useGroupsQuery();
  const [isAddModalOpen, setIsAddModalOpen] = useState(false);

  const toggleAddModal = () => {
    setIsAddModalOpen(!isAddModalOpen);
  };
  return (
    <>
      <div className={"page"}>
        {getGroupsQuery.isLoading ? (
          "Загружаем"
        ) : (
            <>
              <CoursesTable data={getGroupsQuery.data} />
              <Button onClick={toggleAddModal}>Добавить группу</Button>
              <AddGroupModal data={getGroupsQuery.data} isOpen={isAddModalOpen} toggleModal={toggleAddModal} />
            </>
        )}
      </div>
    </>
  );
};

export default Groups;
