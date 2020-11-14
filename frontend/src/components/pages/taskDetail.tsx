import React from 'react';
import { Typography, IconButton } from '@material-ui/core';
import { Edit } from '@material-ui/icons';
import Task from '../../interfaces/task';
import DetailFrame from '../templates/taskDetail';

interface Props { 
  task: Task,
  toEditMode: () => void
};

const Detail = (props: Props): JSX.Element => {
  const name = <Typography variant='h3'>{props.task.Name}</Typography>
    , editButton = (
      <IconButton onClick={props.toEditMode}>
        <Edit  />
      </IconButton>
    )
    , body = <Typography variant='body1'>{props.task.Description}</Typography>
    , footer = <Typography variant='caption'>{props.task.Author.Name}</Typography>;

  return <DetailFrame name={name} editButton={editButton} body={body} footer={footer}/>;
}

export default Detail;
