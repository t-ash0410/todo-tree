import axios from 'axios';
import Task from '../../interfaces/task';

const endPoint = "/api/todo";

export const createTask = (task: Task, callback: (task: Task) => void) => {
  axios.post(endPoint, { params: { task: task } }).then((results) => {
    callback(results.data);
  });
};

export const updateTask = (task: Task, callback: (task: Task) => void) => {
  axios.put(endPoint, { params: { task: task } }).then((results) => {
    callback(results.data);
  });
};

export const deleteTask = (id: string, callback: () => void) => {
  axios.delete(endPoint, { params: { id: id } }).then((results) => {
    callback();
  });
}
