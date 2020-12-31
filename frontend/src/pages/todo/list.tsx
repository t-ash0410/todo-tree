import React from 'react';
import { useRouter } from 'next/router';
import useSWR, { mutate } from "swr";
import { TableRow, TableCell, IconButton } from '@material-ui/core';
import { Delete } from '@material-ui/icons';
import Task from '../../interfaces/task';
import Table from '../../components/templates/table';
import AddTaskButton from '../../components/organisms/addTaskButton';
import ModalButton from '../../components/organisms/modalButton';
import { getEndPoint } from '../../lib/http/task';
import { api } from '../../lib/api';

const ToDoList = () => {
  const { tasks, isLoading, isError } = getTasks();
  const routing = getRouting();
  
  if (isLoading) return <div>loading...</div>
  if (isError) return <div>error!</div>
  return (
    <>
      <Table data={tasks} 
        headerCols={[
          <TableCell component="th" size='small' >ID</TableCell>,
          <TableCell component="th">名称</TableCell>,
          <TableCell component="th"></TableCell>
        ]} 
        createBodyRow={(task: Task) => (
          <TableRow key={task.Id} hover onClick={(e) => routing(e, task.Id)}>
            <TableCell component="th" scope="row">{task.Id}</TableCell>
            <TableCell>{task.Name}</TableCell>
            <TableCell align="right">
              <ModalButton 
                createButton={(onclick) => (
                  <IconButton aria-label="削除" color='secondary' onClick={(e) => onclick(e)}>
                    <Delete fontSize="small" />
                  </IconButton>
                )} 
                title='タスク削除'
                body={<>タスクを削除します。よろしいですか？</>}
                callback={() => deleteTask(tasks, task.Id)}
              />
            </TableCell>
          </TableRow>
        )} 
      />
      <AddTaskButton addTask={(task) => addTask(tasks, task)} />
    </>
  )
}

export default ToDoList;

const defaultEndPoint = getEndPoint()
      , getListEndPoint = `${defaultEndPoint}/list`;

const getTasks = () => {
  const { data, error } = useSWR<Task[], Error>(getListEndPoint, api.get);
  return {
    tasks: data,
    isLoading: !error && !data,
    isError: error
  };
}

const getRouting = () => {
  const router = useRouter();
  const routing = (e: React.MouseEvent, id: number) => {
    if (e.defaultPrevented) return;
    router.push(`/todo/${id}`);
  }
  return routing;
};

const addTask = (tasks: Task[], newTask: Task) => {
  return api.post<Task, Task>(defaultEndPoint, newTask)
    .then(res => {
      mutate(getListEndPoint, [...tasks, res], false);
      return Promise.resolve(res);
    });
};

const deleteTask = (tasks: Task[], id: number) => {
  api.delete(`${defaultEndPoint}?id=${id}`);
  const clone = tasks.concat([]);
  clone.splice(tasks.findIndex(t => t.Id === id), 1);
  mutate(getListEndPoint, clone, false);
};
