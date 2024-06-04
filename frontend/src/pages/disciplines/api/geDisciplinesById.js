import {useQuery} from "@tanstack/react-query";
import {service} from "./service";

export const useDisciplinesByIdQuery = (id) => {
    return useQuery({
        queryKey: ['disciplines:id', id],
        queryFn: () => service.getDisciplinesById(id),
    });
};