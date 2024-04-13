import Main from "../pages/Main";
import Classrooms from "../pages/Classrooms";
import Teachers from "../pages/Teachers";
import Disciplines from "../pages/Disciplines";
import Error from "../pages/Error";
import Groups from "../pages/groups/Groups";

export const routes=[
    {path:'/',element:<Main />},
    {path:'/groups',element:<Groups />},
    {path:'/classrooms',element:<Classrooms />},
    {path:'/teachers',element:<Teachers />},
    {path:'/disciplines',element:<Disciplines />},
    {path: '*',element: <Error />}
]