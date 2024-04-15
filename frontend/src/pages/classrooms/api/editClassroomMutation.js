import {useMutation, useQueryClient} from "@tanstack/react-query";
import {service} from "./service";

export function useEditClassroomMutation() {
    const queryClient = useQueryClient();
    return useMutation({
        mutationKey: ['classroom:edit'],
        mutationFn: (requestParams) => service.editClassroom(requestParams),
        onSuccess: (variables) => {
            const queryKey = ['classroom:edit' ,variables]
            queryClient.invalidateQueries(queryKey);
        }
    });
}