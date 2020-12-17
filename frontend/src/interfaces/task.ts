import User, { getInitialObject as getUserInitialObject } from "./user";

interface Task {
  Id: number;
  Name: string;
  Author: User;
  Description: string;
}

export const getInitialObject: () => Task = () => {
  return {
    Id: 0,
    Name: "",
    Author: getUserInitialObject(),
    Description: ""
  }
}

export default Task;
