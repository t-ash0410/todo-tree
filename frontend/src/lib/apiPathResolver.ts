import { GetServerSidePropsContext } from 'next';

const resolve = (ctx: GetServerSidePropsContext, endpoint: string): string => {
  return `http://${ctx.req.headers.host}/api/${endpoint}`;
};

export default resolve;
