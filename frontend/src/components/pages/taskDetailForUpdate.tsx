import React, { useState } from 'react';
import { Box, IconButton, TextField } from '@material-ui/core';
import { Save } from '@material-ui/icons';
import Task from '../../interfaces/task';
import DetailFrame from '../templates/taskDetail';
import { updateTask as ajaxUpdate } from '../../lib/http/task';

interface Props { 
  task: Task,
  onUpdate: (task: Task) => void
};

const UpdateDetail = (props: Props): JSX.Element => {
  const [task, setTask] = useState(props.task);
  const change = <T extends HTMLTextAreaElement | HTMLInputElement>(e: React.ChangeEvent<T>, setter: (t: Task, v: string) => void) => {
    let clone = {...task};
    setter(clone, e.target.value);
    setTask(clone);
  }

  const complete = () => {
    ajaxUpdate(task, (newTask: Task) => {
      props.onUpdate(newTask);
    });
  }

  const name = (
      <Box p={1} width={'100%'}>
        <TextField 
          label="名称" 
          id="name" 
          value={task.Name} 
          onChange={(e) => change(e, (t, v) => t.Name = v)} 
          variant={"outlined"} 
          required fullWidth autoFocus 
        />
      </Box>
    )
    , editButton = (
      <IconButton color='secondary' onClick={complete}>
        <Save  />
      </IconButton>
    )
    , body = (
      <TextField 
        label="説明" 
        id="description" 
        value={task.Description} 
        onChange={(e) => change(e, (t, v) => t.Description = v)} 
        variant={"outlined"} 
        required fullWidth autoFocus multiline rows={5} 
      />
    )
    , footer = (
      <TextField 
        label="編集者"
        id="author" 
        value={task.Author.Name} 
        onChange={(e) => change(e, (t, v) => t.Author.Name = v)} 
        variant={"outlined"} 
        required fullWidth autoFocus 
      />
    );
  
  return <DetailFrame name={name} editButton={editButton} body={body} footer={footer} />;
}

export default UpdateDetail;
