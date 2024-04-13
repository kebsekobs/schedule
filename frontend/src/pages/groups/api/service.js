import axios from "axios";

const group_url = `http://localhost:3001/groups`;
const lessons_url = 'http://localhost:3001/lessons'

export const service = {
       getGroups: function (config = {}) {
              return axios.get(group_url, config).then(response => response.data);
       },
       getLessons: function (config = {}) {
              return axios.get(lessons_url, config).then(response => response.data);
       },
       addGroups: function (config = {}) {
              return axios.post(group_url, config).then(response => response.data);
       },
       deleteGroup: function (groupId, config = {}) {
              const url = `${group_url}/${groupId}`;  // Создание полного URL с ID группы
              return axios.delete(url, config)
                  .then(response => response.data);
       },
       getGroupById: function (groupId,config = {}) {
              const url = `${group_url}/${groupId}`;  // Создание полного URL с ID группы
              return axios.get(url, config).then(response => response.data);
       },
       editGroup: function (data,config = {}) {
              return axios.put(`${group_url}/${data.id}`,data, config).then(response => response.data);
       },
};
