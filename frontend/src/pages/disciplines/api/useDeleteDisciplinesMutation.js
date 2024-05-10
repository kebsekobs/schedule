import {useMutation, useQueryClient} from "@tanstack/react-query";
import {service} from "./service";

export function useDeleteDisciplinesMutation() {
    const queryClient = useQueryClient();
    return useMutation({
        mutationKey: 'disciplines:delete',
        mutationFn: (classroomId) => service.deleteDisciplines(classroomId),
        onSuccess: (_, variables) => {
            const queryKey = ['disciplines:delete', variables];
            queryClient.invalidateQueries(queryKey);
        }
    });
}