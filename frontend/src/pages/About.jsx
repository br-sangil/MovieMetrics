import React from 'react'

export default function About() {
  return (
    <div className="flex flex-col h-screen">
      
      <div className="flex flex-row">
          <div className="bg-red-600 w-1/2 font-sans text-white p-4 font-bold">WELCOME TO MOVIEMETRICS
          </div>
          <img className="w-1/2" src="/movies.jpeg" alt="logo" />
      </div>

    <div className="flex h-96">
        <div className="bg-red-300 w-1/3 text-6xl font-sans text-white p-4 font-bold">About us</div>
        <div className="bg-slate-300 w-1/3">4</div>
        <div className="bg-red-500 w-1/3">5</div>
    </div>

    <div className="flex">
    <img className="w-1/2" src="/movies3.jpeg" alt="logo" />
        <div className="bg-black w-1/2"></div>
    </div>

    </div>
  )
}
