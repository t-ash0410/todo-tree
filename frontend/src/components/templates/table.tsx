import React from 'react';
import { Table, TableBody, TableContainer, TableHead, TableRow, Paper } from '@material-ui/core';

interface Props<T> { 
  data: T[], 
  headerCols: JSX.Element[],
  createBodyRow: (obj: T) => JSX.Element
};

const List = <T extends any>(props: Props<T>) => {
  return (
    <TableContainer component={Paper}>
      <Table>
        <TableHead>
          <TableRow>
            { props.headerCols }
          </TableRow>
        </TableHead>
        <TableBody>
          { props.data.map(task => props.createBodyRow(task)) }
        </TableBody>
      </Table>
    </TableContainer>
  )
}

export default List;
