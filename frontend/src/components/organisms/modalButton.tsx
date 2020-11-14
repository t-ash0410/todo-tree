import React, { useState } from 'react';
import { 
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle
} from '@material-ui/core';

interface Props { 
  createButton: (handleOpen: (e: React.MouseEvent) => void) => JSX.Element,
  title: string,
  body: JSX.Element,
  callback: () => void
};

const ModalButton = (props: Props): JSX.Element => {
  const [open, setOpen] = useState(false);
  const handleOpen = (e: React.MouseEvent) => {
    e.preventDefault();
    e.stopPropagation();
    setOpen(true);
  };
  const handleClose = (e: React.MouseEvent) => {
    e.preventDefault();
    e.stopPropagation();
    setOpen(false);
  };
  const execute = (e: React.MouseEvent) => {
    handleClose(e);
    props.callback();
  }

  return (
    <>
      {props.createButton(handleOpen)}
      <Dialog open={open} onClose={handleClose} aria-labelledby="form-dialog-title">
        <DialogTitle id="form-dialog-title">
          {props.title}
        </DialogTitle>
        <DialogContent>
          <DialogContentText>
            {props.body}
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={execute} color='secondary'>
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

export default ModalButton;
