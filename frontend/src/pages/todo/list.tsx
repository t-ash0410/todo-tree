import React, { useState } from 'react';
import { GetServerSideProps } from 'next';
import { useRouter } from 'next/router';
import { TableRow, TableCell, IconButton } from '@material-ui/core';
import { Delete } from '@material-ui/icons';
import Task from '../../interfaces/task';
import resolve from '../../lib/apiPathResolver';
import Table from '../../components/templates/table';
import AddTaskButton from '../../components/organisms/addTaskButton';
import ModalButton from '../../components/organisms/modalButton';
import { deleteTask as ajaxDelete } from '../../lib/http/task';

interface Props { tasks: Task[] };

const ToDoList = (props: Props): JSX.Element => {
  const router = useRouter();
  const routing = (e: React.MouseEvent, id: string) => {
    if (e.defaultPrevented) return;
    router.push(`/todo/${id}`);
  }
  const [tasks, setTasks] = useState(props.tasks);
  const addTask = (newTask: Task) => {
    setTasks([...tasks, newTask]);
  };
  const deleteTask = (id: string) => {
    ajaxDelete(id, () => {
      const clone = tasks.concat([]);
      clone.splice(tasks.findIndex(t => t.Id === id), 1);
      setTasks(clone);
    });
  }
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
                callback={() => deleteTask(task.Id)}
              />
            </TableCell>
          </TableRow>
        )} 
      />
      <AddTaskButton addTask={addTask} />
    </>
  )
}

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  const res = await fetch(resolve(ctx, 'todo'));
  const data = await res.json();
  return { props: { tasks: data } };
};

export default ToDoList;
