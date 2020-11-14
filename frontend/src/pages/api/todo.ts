import { NextApiRequest, NextApiResponse } from 'next';
import Task from '../../interfaces/task';

let data: Task[] = [
  {
    Id: "1",
    Name: "task 1",
    Author: {
      Id: "1",
      Name: "Isono"
    },
    Description: "this is first task"
  },
  {
    Id: "2",
    Name: "二個目のタスク",
    Author: {
      Id: "1",
      Name: "Isono"
    },
    Description: "優先順位が2位のタスクですよ"
  }
];

const read = () => {
  return data;
}

const create = (req: NextApiRequest): Task => {
  const task = paramToObj(req);
  data.push(task);
  return task;
}

const update = (req: NextApiRequest): Task => {
  const task = paramToObj(req);
  data[data.findIndex(t => t.Id === task.Id)] = task;
  return task;
}

const del = (req: NextApiRequest) => {
  data.splice(data.findIndex(t => t.Id === req.body.id), 1);
}

const paramToObj = (req: NextApiRequest): Task => {
  let task: Task = req.body.params.task;
  if (task.Id.length === 0){
    task.Id = `id-${Math.random()}`;
  }
  if (task.Author.Id.length === 0){
    task.Author.Id = `id-${Math.random()}`;
  }
  return task;
}

export default (req: NextApiRequest, res: NextApiResponse<Task | Task[] | {}>) => {
  switch(req.method){
    case "GET": {
      res.status(200).json(read());
      break;
    }
    case "POST": {
      const task = create(req);
      res.status(200).json(task);
      break;
    }
    case "PUT": {
      const task = update(req);
      res.status(200).json(task);
      break;
    }
    case "DELETE": {
      del(req);
      res.status(200).json({});
      break;
    }
  }
}
