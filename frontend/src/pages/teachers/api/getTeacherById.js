import {useQuery} from "@tanstack/react-query";
import {service} from "./service";

export const useTeacherByIdQuery = (id) => {
    return useQuery({
        queryKey: ['teachers', id],
        queryFn: () => service.getTeacherById(id),
    });
};