import {useMutation, useQueryClient} from "@tanstack/react-query";
import {service} from "./service";

export function useAddDisciplinesMutation() {
    const queryClient = useQueryClient();
    return useMutation({
        mutationKey: ['disciplines:add'],
        mutationFn: (requestParams) => service.addDisciplines(requestParams),
        onSuccess: (variables) => {
            const queryKey = ['disciplines:add' ,variables]
            queryClient.invalidateQueries(queryKey);
        }
    });
}