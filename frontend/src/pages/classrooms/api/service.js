import axios from "axios";

const classrooms_url = `http://localhost:3001/classrooms`;

export const service = {
       getClassrooms: function (config = {}) {
              return axios.get(classrooms_url, config).then(response => response.data);
       },
       addClassroom: function (config = {}) {
              return axios.post(classrooms_url, config).then(response => response.data);
       },
       deleteClassroom: function (classroomId, config = {}) {
              const url = `${classrooms_url}/${classroomId}`;
              return axios.delete(url, config)
                  .then(response => response.data);
       },
       getClassroomById: function (classroomId, config = {}) {
              const url = `${classrooms_url}/${classroomId}`;
              return axios.get(url, config).then(response => response.data);
       },
       editClassroom: function (data, config = {}) {
              return axios.put(`${classrooms_url}`, data, config).then(response => response.data);
       },
};
