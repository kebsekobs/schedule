import {useMutation, useQueryClient} from "@tanstack/react-query";
import {service} from "./service";

export function useEditTeacherMutation() {
    const queryClient = useQueryClient();
    return useMutation({
        mutationKey: ['teacher:edit'],
        mutationFn: (requestParams) => service.editTeacher(requestParams),
        onSuccess: (variables) => {
            const queryKey = ['teacher:edit', variables]
            queryClient.invalidateQueries(queryKey);
        }
    });
}