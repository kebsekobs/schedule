import {useMutation, useQueryClient} from "@tanstack/react-query";
import {service} from "./service";

export function useAddTeacherMutation() {
    const queryClient = useQueryClient();
    return useMutation({
        mutationKey: ['teacher:add'],
        mutationFn: (requestParams) => service.addTeachers(requestParams),
        onSuccess: (variables) => {
            const queryKey = ['teacher:add' ,variables]
            queryClient.invalidateQueries(queryKey);
        }
    });
}