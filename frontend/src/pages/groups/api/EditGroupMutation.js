import {useMutation, useQueryClient} from "@tanstack/react-query";
import {service} from "./service";

export function useEditGroupMutation() {
    const queryClient = useQueryClient();
    return useMutation({
        mutationKey: ['group:edit'],
        mutationFn: (requestParams) => service.editGroup(requestParams),
        onSuccess: (variables) => {
            const queryKey = ['group:edit' ,variables]
            queryClient.invalidateQueries(queryKey);
        }
    });
}