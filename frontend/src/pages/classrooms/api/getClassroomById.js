import {useQuery} from "@tanstack/react-query";
import {service} from "./service";

export const useClassroomByIdQuery = (id) => {
    return useQuery({
        queryKey: ['classrooms', id],
        queryFn: () => service.getClassroomById(id),
    });
};