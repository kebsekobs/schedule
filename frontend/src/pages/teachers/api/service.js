import axios from "axios";

const teachers_url = `http://localhost:3001/teachers`;

export const service = {
       getTeachers: function (config = {}) {
              return axios.get(teachers_url, config).then(response => response.data);
       },
       addTeachers: function (config = {}) {
              return axios.post(teachers_url, config).then(response => response.data);
       },
       deleteTeacher: function (teacherId, config = {}) {
              const url = `${teachers_url}`;
              return axios.delete(url, config)
                  .then(response => response.data);
       },
       getTeacherById: function (teacherId, config = {}) {
              const url = `${teachers_url}`;
              return axios.get(url, config).then(response => response.data);
       },
       editTeacher: function (data, config = {}) {
              return axios.put(`${teachers_url}`, data, config).then(response => response.data);
       },
};
