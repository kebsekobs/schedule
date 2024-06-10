import React, { useState } from "react";
import { useTeachersQuery } from "./api/getTeachersQuery";
import CoursesTable from "./table/use-table";
import AddTeacherModal from "./modals/AddModal.js";
import Button from "../../components/button";

const Teachers = () => {
  const getTeachersQuery = useTeachersQuery();
  const [isAddModalOpen, setIsAddModalOpen] = useState(false);

  const toggleAddModal = () => {
    setIsAddModalOpen(!isAddModalOpen);
  };

  return (
    <>
      <div className={"page"}>
        {getTeachersQuery.isLoading ? (
          "Загружаем"
        ) : (
            <>
                <CoursesTable data={getTeachersQuery.data} />
                <Button onClick={toggleAddModal}>Добавить преподавателя</Button>
                <AddTeacherModal
                  isOpen={isAddModalOpen}
                  toggleModal={toggleAddModal}
                />
            </>
        )}
      </div>
    </>
  );
};

export default Teachers;