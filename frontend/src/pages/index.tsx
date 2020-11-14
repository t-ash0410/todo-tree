import { NextPage } from 'next';
import Link from 'next/link';

const Index: NextPage = () => {
  return (
    <>
      <h1>Welcome!</h1>
      <ul>
        <li>
          <Link href="/todo/list">
            <a>todo</a>
          </Link>
        </li>
      </ul>
    </>
  )
}

export default Index;
