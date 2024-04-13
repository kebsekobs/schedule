import {useQuery} from "@tanstack/react-query";
import {service} from "./service";

export const useGroupByIdQuery = (id) => {
    return useQuery({
        queryKey: ['groups', id],
        queryFn: () => service.getGroupById(id),
    });
};