import Main from "../pages/Main";
import Classrooms from "../pages/Classrooms";
import Teachers from "../pages/Teachers";
import Error from "../pages/Error";
import Groups from "../pages/groups/Groups";
import Disciplines from "../pages/disciplines/Disciplines";

export const routes=[
    {path:'/',element:<Main />},
    {path:'/groups',element:<Groups />},
    {path:'/classrooms',element:<Classrooms />},
    {path:'/teachers',element:<Teachers />},
    {path:'/disciplines',element:<Disciplines />},
    {path: '*',element: <Error />}
]