import React, { useState } from 'react';
import { useRouter } from 'next/router';
import Link from 'next/link';
import useSWR, { mutate } from "swr";
import { Box, Grid } from '@material-ui/core';
import ChevronLeft from '@material-ui/icons/ChevronLeft';
import Task from '../../interfaces/task';
import { api } from '../../lib/api';
import { getEndPoint } from '../../lib/http/task';
import TaskView from '../../components/pages/taskDetail';
import UpdateDetail from '../../components/pages/taskDetailForUpdate';

const Detail = (): JSX.Element => {
  const { endPoint, endPointWithParam } = getRouterInfo();
  const { task, isLoading, isError } = getTask(endPointWithParam);
  const [edit, setEdit] = useState(false);

  if (isLoading) return <div>loading...</div>
  if (isError) return <div>error!</div>
  return (
    <>
      {
        edit
          ? <UpdateDetail task={task} onUpdate={(task) => onUpdate(endPoint, endPointWithParam, task, setEdit)} />
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

export default Detail;

const getRouterInfo = () => {
  const router = useRouter();
  const endPoint = getEndPoint();
  return {
    endPoint: endPoint,
    endPointWithParam: `${endPoint}?id=${router.query.id}`
  };
}

const getTask = (endPointWithParam: string) => {
  const { data, error } = useSWR<Task, Error>(endPointWithParam, api.get, { shouldRetryOnError: false });
  return {
    task: data,
    isLoading: !error && !data,
    isError: error
  };
}

const onUpdate = (endPoint: string, endPointWithParam: string, task: Task, setEdit: (boolean) => void) => {
  api.put<Task, Task>(endPoint, task)
    .then((newTask) => {
      mutate(endPointWithParam, newTask, false);
      setEdit(false);
    });
}
