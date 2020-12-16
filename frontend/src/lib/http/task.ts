import { pathResolver } from '../api';
export const getEndPoint: () => string = () => pathResolver("todo");
