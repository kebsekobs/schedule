import Main from "../pages/Main";
import Groups from "../pages/Groups";
import Classrooms from "../pages/Classrooms";
import Teachers from "../pages/Teachers";
import Disciplines from "../pages/Disciplines";

export const routes=[
    {path:'/',element:<Main />},
    {path:'/groups',element:<Groups />},
    {path:'/classrooms',element:<Classrooms />},
    {path:'/teachers',element:<Teachers />},
    {path:'/disciplines',element:<Disciplines />},
]