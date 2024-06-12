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

  // курсы мага
  const firstYearMag = getGroupsQuery.data?.filter(
    (group) =>
      Math.floor(group.groupId / 100) === 1 && group.magistracy === true
  );
  const secondYearMag = getGroupsQuery.data?.filter(
    (group) =>
      Math.floor(group.groupId / 100) === 2 && group.magistracy === true
  );

  // курсы бакалавры
  const firstYearBach = getGroupsQuery.data?.filter(
    (group) =>
      Math.floor(group.groupId / 100) === 1 && group.magistracy === false
  );
  const secondYearBach = getGroupsQuery.data?.filter(
    (group) =>
      Math.floor(group.groupId / 100) === 2 && group.magistracy === false
  );
  const thirdYearBach = getGroupsQuery.data?.filter(
    (group) =>
      Math.floor(group.groupId / 100) === 3 && group.magistracy === false
  );
  const fourthYearBach = getGroupsQuery.data?.filter(
    (group) =>
      Math.floor(group.groupId / 100) === 4 && group.magistracy === false
  );
  const fifthYearBach = getGroupsQuery.data?.filter(
    (group) =>
      Math.floor(group.groupId / 100) === 5 && group.magistracy === false
  );

  const data = [
    {
      year: '1 курс маг.',
      subRows: firstYearMag,
    },
    {
      year: '2 курс маг.',
      subRows: secondYearMag,
    },
    {
      year: '1 курс бак.',
      subRows: firstYearBach,
    },
    {
      year: '2 курс бак.',
      subRows: secondYearBach,
    },
    {
      year: '3 курс бак.',
      subRows: thirdYearBach,
    },
    {
      year: '4 курс бак.',
      subRows: fourthYearBach,
    },
    {
      year: '5 курс бак.',
      subRows: fifthYearBach,
    },
  ]

  return (
    <>
      <div className={"page"}>
        {getGroupsQuery.isLoading ? (
          "Загружаем"
        ) : (
          <>
            <Button onClick={toggleAddModal}>Добавить группу</Button>
            <CoursesTable data={data} />
            <AddGroupModal
              data={getGroupsQuery.data}
              isOpen={isAddModalOpen}
              toggleModal={toggleAddModal}
            />
          </>
        )}
      </div>
    </>
  );
};

export default Groups;
