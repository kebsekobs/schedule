import {useMutation, useQueryClient} from "@tanstack/react-query";
import {service} from "./service";

export function useDeleteGroupMutation() {
    const queryClient = useQueryClient();
    return useMutation({
        mutationKey: 'group:delete',
        mutationFn: (groupId) => service.deleteGroup(groupId),
        onSuccess: (_, variables) => {
            const queryKey = ['group:delete', variables];
            queryClient.invalidateQueries(queryKey);
        }
    });
}
