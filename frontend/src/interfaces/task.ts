import User from "./user";

interface Task {
  Id: string;
  Name: string;
  Author: User;
  Description: string;
}

export default Task;
