// dynamic routes are used to create pages with dynamic paths.
// any parameter name (id in here) can be used as a query parameter, when file name is wrapped in square brackets (eg> [id].js).
// for example, /movies/1234 will be rendered by this page.

import { useRouter } from "next/router";

export default function Movie() {
    const router = useRouter();
    console.log(router); // {pathname: "/movies/[id]", query: {…}, asPath: "/movies/1234", …}
    return (
        <div>
            <h4>{router.query.title || "Loading ... "}</h4>
        </div>
    );
}
