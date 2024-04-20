import {useMutation, useQueryClient} from "@tanstack/react-query";
import {service} from "./service";

export function useDeleteClassroomMutation() {
    const queryClient = useQueryClient();
    return useMutation({
        mutationKey: 'classroom:delete',
        mutationFn: (classroomId) => service.deleteClassroom(classroomId),
        onSuccess: (_, variables) => {
            const queryKey = ['classroom:delete', variables];
            queryClient.invalidateQueries(queryKey);
        }
    });
}