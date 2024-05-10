import Main from "../pages/Main";
import Error from "../pages/Error";
import Groups from "../pages/groups/Groups";
import Disciplines from "../pages/disciplines/Disciplines";
import Classrooms from "../pages/classrooms/Classrooms";
import Teachers from "../pages/teachers/Teachers";

export const routes=[
    {path:'/',element:<Main />},
    {path:'/groups',element:<Groups />},
    {path:'/classrooms',element:<Classrooms />},
    {path:'/teachers',element:<Teachers />},
    {path:'/disciplines',element:<Disciplines />},
    {path: '*',element: <Error />}
]