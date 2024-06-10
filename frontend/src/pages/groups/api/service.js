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
       deleteGroup: function (groupId, config = {}) {
              const url = `${group_url}`;  // Создание полного URL с ID группы
              return axios.delete(url, config)
                  .then(response => response.data);
       },
       getGroupById: function (groupId,config = {}) {
              const url = `${group_url}/${groupId}`;  // Создание полного URL с ID группы
              return axios.get(url, config).then(response => response.data);
       },
       editGroup: function (data,config = {}) {
              return axios.put(`${group_url}`,data, config).then(response => response.data);
       },
};
