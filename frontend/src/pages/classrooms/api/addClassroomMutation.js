import {useMutation, useQueryClient} from "@tanstack/react-query";
import {service} from "./service";

export function useAddClassroomMutation() {
    const queryClient = useQueryClient();
    return useMutation({
        mutationKey: ['classroom:add'],
        mutationFn: (requestParams) => service.addClassroom(requestParams),
        onSuccess: (variables) => {
            const queryKey = ['classroom:add' ,variables]
            queryClient.invalidateQueries(queryKey);
        }
    });
}