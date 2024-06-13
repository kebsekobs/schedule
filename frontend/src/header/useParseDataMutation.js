import { useMutation } from "@tanstack/react-query";
import axios from "axios";

export const useParseMutation = () => {
    return useMutation({
        mutationKey: 'parse',
        mutationFn: () => service.parse()
        });
};
const url = `http://localhost:3001/parsedata`;
export const service = {
    parse: function (config = {}) {
        return axios.get(url, config).then(response => response.data);
    },

};