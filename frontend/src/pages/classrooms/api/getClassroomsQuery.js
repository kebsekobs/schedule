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
        return a.groupId - b.groupId;
    });
}
