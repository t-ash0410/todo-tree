import React, { useState } from 'react';
import { GetServerSideProps } from 'next';
import Link from 'next/link';
import { Box, Grid } from '@material-ui/core';
import ChevronLeft from '@material-ui/icons/ChevronLeft';
import Task from '../../interfaces/task';
import resolve from '../../lib/apiPathResolver';
import TaskView from '../../components/pages/taskDetail';
import UpdateDetail from '../../components/pages/taskDetailForUpdate';

interface Props { 
  task?: Task,
  errors?: string
};

const Detail = (props: Props): JSX.Element => {
  const [task, setTask] = useState(props.task);
  const [edit, setEdit] = useState(false);
  const onUpdate = (task: Task) => {
    setTask(task);
    setEdit(false);
  }

  if (props.errors) {
    return (
      <p>
        <span style={{ color: 'red' }}>Error:</span> {props.errors}
      </p>
    )
  } else {
    return (
      <>
        {
          edit
          ? <UpdateDetail task={task} onUpdate={onUpdate} />
          : <TaskView task={task} toEditMode={() => setEdit(true)} />
        }
        <Box p={1}>
          <Link href='/todo/list'>
            <a>
              <Grid container direction='row' alignItems='center'>
                <ChevronLeft /> back
              </Grid>
            </a>
          </Link>
        </Box>
      </>
    )
  }
} 

export const getServerSideProps: GetServerSideProps<Props> = async (ctx) => {
  try {
    const id = ctx.params.id;
    const res = await fetch(resolve(ctx, 'todo'));
    const tasks: Task[] = await res.json();
    const task = tasks.find(d => d.Id === id);
    return { 
      props: { 
        task
      } 
    }
  } catch (err) {
    return { 
      props: { 
        errors: err.message 
      } 
    }
  }
};

export default Detail;
