import axios from 'axios';

export const api = {
  get: <TP, TRes>(url: string, params?: TP) => axios.get(url, { params: params }).then(res => Promise.resolve(res.data as TRes)),
  post: <TP, TRes>(url: string, params?: TP) => axios.post(url, { params: params }).then(res => Promise.resolve(res.data as TRes)),
  put: <TP, TRes>(url: string, params?: TP) => axios.put(url, { params: params }).then(res => Promise.resolve(res.data as TRes)),
  delete: <TP, TRes>(url: string, params?: TP) => axios.delete(url, { params: params }).then(res => Promise.resolve(res.data as TRes))
};

export const pathResolver = (endpoint: string) => `http://localhost:8080/${endpoint}`;
