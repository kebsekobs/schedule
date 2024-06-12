import { useQuery } from "@tanstack/react-query";
import axios from "axios";

export const useGenerateMutation = () => {
    return useQuery({
        queryKey: ['generate'],
        queryFn: () => service.generate(),
    });
};
const url = `http://localhost:3001/generate`;
export const service = {
    generate: function (config = {}) {
           return axios.get(url, config).then(response => response.data);
    },
    
};