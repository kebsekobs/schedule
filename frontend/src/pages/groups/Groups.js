import React, {useState} from "react";
import {useGroupsQuery} from "./api/getGroupsQuery";
import CoursesTable from "./table/use-table";
import AddGroupModal from "./modals/AddModal";

const Groups = () => {
   const getGroupsQuery = useGroupsQuery()
    const [isAddModalOpen, setIsAddModalOpen] = useState(false);

    const toggleAddModal = () => {
        setIsAddModalOpen(!isAddModalOpen);
    };
  return (
    <div className={'page'}>
        {getGroupsQuery.isLoading ? 'Загружаем' : <CoursesTable data={getGroupsQuery.data} />}
        <button
            style={{cursor: 'pointer'}}
            title={'Добавить'}
            name={'Добавить'}
            onClick={toggleAddModal}>
            Добавить
        </button>
        <AddGroupModal isOpen={isAddModalOpen} toggleModal={toggleAddModal} />
    </div>
  );
};

export default Groups;
