import { useQuery } from "@tanstack/react-query";
import { service } from "./service";

export const useTeachersQuery = () => {
  return useQuery({
    queryKey: ["teacher"],
    queryFn: () => service.getTeachers(),
    select: (data) => sortData(data),
  });
};

function sortData(data) {
  return data.sort((a, b) => {
    if (a.name < b.name) {
      return -1;
    }
    if (a.name > b.name) {
      return 1;
    }
    return 0;
  });
}
