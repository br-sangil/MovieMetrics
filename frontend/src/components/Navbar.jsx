import React from "react";
import NavItem from "./NavItem";
export default function Navbar() {
  return (
    <nav className="p-2 flex bg-black justify-between">
        <a href="/">
          <img className="w-40" src="/moviemetrics.png" alt="logo" />
        </a>
      <ul className="place-items-center flex ">
        <NavItem content="Home" href="/" />
        <NavItem content="Surprise Me!" href="/random" />
        <NavItem content="About" href="/about" />
        <NavItem content="Register" href="/register" />
        <NavItem content="Login" href="/login" />
      </ul>
    </nav>
  );
  
}