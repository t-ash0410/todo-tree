interface User {
  Id: string;
  Name: string;
}

export const getInitialObject: () => User = () => {
  return {
    Id: "",
    Name: ""
  };
}

export default User;
