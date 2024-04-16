import {useMutation, useQueryClient} from "@tanstack/react-query";
import {service} from "./service";

export function useDeleteTeacherMutation() {
    const queryClient = useQueryClient();
    return useMutation({
        mutationKey: 'teacher:delete',
        mutationFn: (teacherId) => service.deleteTeacher(teacherId),
        onSuccess: (_, variables) => {
            const queryKey = ['teacher:delete', variables];
            queryClient.invalidateQueries(queryKey);
        }
    });
}
