import axios from "axios";

const group_url = `http://localhost:3001/groups`;
const disciplines_url = 'http://localhost:3001/disciplines'

export const service = {
       getGroups: function (config = {}) {
              return axios.get(group_url, config).then(response => response.data);
       },
       getDisciplines: function (config = {}) {
              return axios.get(disciplines_url, config).then(response => response.data);
       },
       addGroups: function (config = {}) {
              return axios.post(group_url, config).then(response => response.data);
       },
       deleteGroup: function (params, config = {}) {
              const url = `${group_url}/delete`;  // Создание полного URL с ID группы
              return axios.post(url, params, config)
                  .then(response => response.data);
       },
       editGroup: function (data,config = {}) {
              return axios.put(`${group_url}`,data, config).then(response => response.data);
       },
};
