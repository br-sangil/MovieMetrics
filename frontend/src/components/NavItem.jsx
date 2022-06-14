import React from 'react'

export default function NavItem({ content, href }) {
  return (
    <li className="px-4 text-slate-400 hover:text-white text-sm md:text-xl">
        <a href={href}>{content}</a>
    </li>
  );
}
