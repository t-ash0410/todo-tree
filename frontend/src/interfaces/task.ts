import User, { getInitialObject as getUserInitialObject } from "./user";

interface Task {
  Id: string;
  Name: string;
  Author: User;
  Description: string;
}

export const getInitialObject: () => Task = () => {
  return {
    Id: "",
    Name: "",
    Author: getUserInitialObject(),
    Description: ""
  }
}

export default Task;
