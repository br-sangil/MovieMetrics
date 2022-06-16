import React from "react";
import {Outlet} from "react-router-dom";
import Navbar from "../components/Navbar";
import Chat from "../components/Chat";
const Layout = () => {
 return(
<>
<Navbar/>
<Outlet />
</>
);
};

export default Layout;