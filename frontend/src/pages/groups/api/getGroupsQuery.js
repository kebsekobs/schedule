import { useQuery } from "@tanstack/react-query";
import { service } from "./service";

export const useGroupsQuery = () => {
    return useQuery({
        queryKey: ['groups'],
        queryFn: () => service.getGroups(),
        select: data => sortToExpand(sortData(data))
    });
};

function sortData(data) {
    return data.sort((a, b) => {
        if (!a.magistracy && b.magistracy) return -1;
        if (a.magistracy && !b.magistracy) return 1;

        return a.groupId - b.groupId;
    });
}
function sortToExpand (data) {
// курсы мага
const firstYearMag = data.filter(
    (group) =>
      Math.floor(Number(group.groupId) / 100) === 1 && group.magistracy === true
  );
  const secondYearMag = data.filter(
    (group) =>
      Math.floor(Number(group.groupId) / 100) === 2 && group.magistracy === true
  );

  // курсы бакалавры
  const firstYearBach = data.filter(
    (group) =>
      Math.floor(Number(group.groupId) / 100) === 1 && group.magistracy === false
  );
  const secondYearBach = data.filter(
    (group) =>
      Math.floor(Number(group.groupId) / 100) === 2 && group.magistracy === false
  );
  const thirdYearBach = data.filter(
    (group) =>
      Math.floor(Number(group.groupId) / 100) === 3 && group.magistracy === false
  );
  const fourthYearBach = data.filter(
    (group) =>
      Math.floor(Number(group.groupId) / 100) === 4 && group.magistracy === false
  );
  const fifthYearBach = data.filter(
    (group) =>
      Math.floor(Number(group.groupId) / 100) === 5 && group.magistracy === false
  );
  const undefinedBanch = data.filter((group) => !group.groupId)
  return [
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
    {
        year: 'Неопрделенные',
        subRows: undefinedBanch
    }
  ]
}
 

  