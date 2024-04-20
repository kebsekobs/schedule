import { useQuery } from "@tanstack/react-query";
import { service } from "./service";

export const useClassroomsQuery = () => {
    return useQuery({
        queryKey: ['classrooms'],
        queryFn: () => service.getClassrooms(),
        select: data => sortData(data)
    });
};

function sortData(data) {
    return data.sort((a, b) => {
        if (a.classroomId < b.classroomId) {
            return -1;
          }
          if (a.classroomId > b.classroomId) {
            return 1;
          }
          return 0;
    });
}
