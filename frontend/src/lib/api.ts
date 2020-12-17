import axios from 'axios';

export const api = {
  get: <TRes>(url: string) => axios.get(url).then(res => Promise.resolve(res.data as TRes)),
  post: <TP, TRes>(url: string, params?: TP) => axios.post(url, params).then(res => Promise.resolve(res.data as TRes)),
  put: <TP, TRes>(url: string, params?: TP) => axios.put(url, params).then(res => Promise.resolve(res.data as TRes)),
  delete: <TRes>(url: string) => axios.delete(url).then(res => Promise.resolve(res.data as TRes))
};

export const pathResolver = (endpoint: string) => `http://localhost:8080/${endpoint}`;
