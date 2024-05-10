import React, {useState} from 'react';
import DisciplinesTable from "./table/use-table";
import {useGetDisciplinesQuery} from "../shared/query/getLessonsQuery";
import AddDisciplinesModal from "./modals/addModal";
import Button from "../../components/button";

const Disciplines = () => {
    const getDisciplines = useGetDisciplinesQuery();
    const [isAddModalOpen, setIsAddModalOpen] = useState(false);
    const toggleAddModal = () => {
        setIsAddModalOpen(!isAddModalOpen);
    };
    return (
        <div className={'page'}>
            {getDisciplines.isLoading ? 'Загружаем' : <DisciplinesTable data={getDisciplines.data} />}
            <Button onClick={toggleAddModal}>Добавить аудиторию</Button>
            <AddDisciplinesModal toggleModal={toggleAddModal} isOpen={isAddModalOpen}/>
        </div>
    );
};

export default Disciplines;