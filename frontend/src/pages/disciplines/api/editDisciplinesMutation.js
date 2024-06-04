import {useMutation, useQueryClient} from "@tanstack/react-query";
import {service} from "./service";

export function useEditDisciplinesMutation() {
    const queryClient = useQueryClient();
    return useMutation({
        mutationKey: ['disciplines:edit'],
        mutationFn: (requestParams) => service.editDisciplines(requestParams),
        onSuccess: (variables) => {
            const queryKey = ['disciplines:edit' ,variables]
            queryClient.invalidateQueries(queryKey);
        }
    });
}