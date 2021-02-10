import axios from 'axios';

const caller = axios.create({
  headers: {
    'x-api-key': process.env.NEXT_PUBLIC_API_KEY,
  },
  responseType: 'json'
});

export const api = {
  get: <TRes>(url: string) => caller.get(url).then(res => Promise.resolve(res.data as TRes)),
  post: <TP, TRes>(url: string, params?: TP) => caller.post(url, params).then(res => Promise.resolve(res.data as TRes)),
  put: <TP, TRes>(url: string, params?: TP) => caller.put(url, params).then(res => Promise.resolve(res.data as TRes)),
  delete: <TRes>(url: string) => caller.delete(url).then(res => Promise.resolve(res.data as TRes))
};

export const pathResolver = (endpoint: string) => `${process.env.NEXT_PUBLIC_API_ENDOPOINT}/${endpoint}`;
