import React, { useState } from 'react';
import { 
  Box, 
  Button, 
  TextField, 
  IconButton, 
  Dialog, 
  DialogActions, 
  DialogContent,
  DialogContentText,
  DialogTitle,
} from '@material-ui/core';
import AddIcon from '@material-ui/icons/Add';
import Task from '../../interfaces/task';

interface Props { addTask: (newTask: Task) => Promise<Task> };

const AddTaskButton = (props: Props): JSX.Element => {
  const [open, setOpen] = useState(false);
  const handleClickOpen = () => {
    setOpen(true);
  };
  const handleClose = () => {
    setOpen(false);
  };

  const defaultValue = { Id: 0, Name: "", Description: "", Author: { Id: "", Name: "" } };
  const [task, setTask] = useState(defaultValue);
  const changeName = (name: string) => {
    setTask({...task, Name: name });
  }
  const changeDesc = (desc: string) => {
    setTask({...task, Description: desc });
  }
  const changeAuthorName = (authorName: string) => {
    setTask({...task, Author: {...task.Author, Name: authorName} });
  }

  const submit = () => {
    props.addTask(task)
      .then(res => {
        setTask(defaultValue);
        handleClose();
      });
  }

  return (
    <>
      <Box textAlign='right' p={1}>
        <IconButton aria-label="追加" color="primary" onClick={() => handleClickOpen()}>
          <AddIcon fontSize="large" />
        </IconButton>
      </Box>
      <Dialog open={open} onClose={handleClose} aria-labelledby="form-dialog-title">
        <DialogTitle id="form-dialog-title">新規タスク追加</DialogTitle>
        <DialogContent>
          <DialogContentText>
            新規タスクを追加します。各フィールドに情報を入力してください。
          </DialogContentText>
          <TextField margin="dense" id="name" label="名称" value={task.Name} onChange={(e) => changeName(e.target.value)} fullWidth autoFocus />
          <TextField margin="dense" id="description" label="説明" value={task.Description} onChange={(e) => changeDesc(e.target.value)} fullWidth />
          <TextField margin="dense" id="author" label="編集者" value={task.Author.Name} onChange={(e) => changeAuthorName(e.target.value)} fullWidth />
        </DialogContent>
        <DialogActions>
          <Button onClick={submit} color="primary">
            OK
          </Button>
          <Button onClick={handleClose}>
            閉じる
          </Button>
        </DialogActions>
      </Dialog>
    </>
  )
}

export default AddTaskButton;
