import axios from "axios";

const disciplines_url = `http://localhost:3001/disciplines`;

export const service = {
       addDisciplines: function (config = {}) {
              return axios.post(disciplines_url, config).then(response => response.data);
       },
       deleteDisciplines: function (params, config = {}) {
              const url = `${disciplines_url}/delete`;
              return axios.post(url, params, config)
                  .then(response => response.data);
       },
       getDisciplinesById: function (classroomId, config = {}) {
              const url = `${disciplines_url}/${classroomId}`;
              return axios.get(url, config).then(response => response.data);
       },
       editDisciplines: function (data, config = {}) {
              return axios.put(`${disciplines_url}`, data, config).then(response => response.data);
       },
};
