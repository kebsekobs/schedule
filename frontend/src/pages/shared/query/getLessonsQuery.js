import { useQuery } from "@tanstack/react-query";
import {service} from "../../groups/api/service";

export const useLessonsQuery = () => {
    return useQuery({
        queryKey: ['lessons'],
        queryFn: () => service.getGroups()
    });
};

