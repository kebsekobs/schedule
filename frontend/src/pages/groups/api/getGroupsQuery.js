import { useQuery } from "@tanstack/react-query";
import { service } from "./service";

export const useGroupsQuery = () => {
    return useQuery({
        queryKey: ['groups'],
        queryFn: () => service.getGroups(),
        select: data => sortData(data)
    });
};

function sortData(data) {
    return data.sort((a, b) => {
        if (!a.magistracy && b.magistracy) return -1;
        if (a.magistracy && !b.magistracy) return 1;

        return a.groupId - b.groupId;
    });
}
