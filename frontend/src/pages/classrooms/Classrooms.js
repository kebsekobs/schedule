import React, { useState } from "react";
import { useClassroomsQuery } from "./api/getClassroomsQuery";
import CoursesTable from "./table/use-table";
import AddClassroomModal from "./modals/addModal.js";
import Button from "../../components/button";

const Classrooms = () => {
  const getClassroomsQuery = useClassroomsQuery();
  const [isAddModalOpen, setIsAddModalOpen] = useState(false);
  const toggleAddModal = () => {
    setIsAddModalOpen(!isAddModalOpen);
  };

  return (
    <>
      <div className={"page"}>
        {getClassroomsQuery.isLoading ? (
          "Загружаем"
        ) : (
          <>
            <Button onClick={toggleAddModal}>Добавить аудиторию</Button>
            <CoursesTable data={getClassroomsQuery.data} />
            <AddClassroomModal
              data={getClassroomsQuery.data}
              isOpen={isAddModalOpen}
              toggleModal={toggleAddModal}
            />
          </>
        )}
      </div>
    </>
  );
};

export default Classrooms;
