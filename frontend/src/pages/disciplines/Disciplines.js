import React from 'react';
import DisciplinesTable from "./table/use-table";
import {useGetDisciplinesQuery} from "../shared/query/getLessonsQuery";

const Disciplines = () => {
    const getDisciplines = useGetDisciplinesQuery();
    return (
        <div className={'page'}>
            {getDisciplines.isLoading ? 'Загружаем' : <DisciplinesTable data={getDisciplines.data} />}
        </div>
    );
};

export default Disciplines;