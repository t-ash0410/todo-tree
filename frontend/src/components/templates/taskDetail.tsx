import React from 'react';
import { Box, Container, Divider, Paper, Grid } from '@material-ui/core';

interface Props { 
  name: JSX.Element,
  editButton: JSX.Element,
  body: JSX.Element,
  footer: JSX.Element
};

const DetailFrame = (props: Props): JSX.Element => {
  return (
    <Paper elevation={3}>
      <Box p={2}>
        <Grid container justify='space-between' alignItems='center' wrap='nowrap'>
          {props.name}
          {props.editButton}
        </Grid>
      </Box>
      <Divider />
      <Container>
        <Box py={2}>
          {props.body}
        </Box>
        <Divider />
        <Box py={1} textAlign='right'>
          {props.footer}
        </Box>
      </Container>
    </Paper>
  )
}

export default DetailFrame;
