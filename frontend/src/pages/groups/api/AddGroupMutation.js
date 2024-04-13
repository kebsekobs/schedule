import {useMutation, useQueryClient} from "@tanstack/react-query";
import {service} from "./service";

export function useAddGroupsMutation() {
    const queryClient = useQueryClient();
    return useMutation({
        mutationKey: ['group:add'],
        mutationFn: (requestParams) => service.addGroups(requestParams),
        onSuccess: (variables) => {
            const queryKey = ['group:add' ,variables]
            queryClient.invalidateQueries(queryKey);
        }
    });
}