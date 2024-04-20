import { useQuery } from "@tanstack/react-query";
import {service} from "../../groups/api/service";

export const useGetDisciplinesQuery = () => {
    return useQuery({
        queryKey: ['lessons'],
        queryFn: () => service.getLessons()
    });
};

