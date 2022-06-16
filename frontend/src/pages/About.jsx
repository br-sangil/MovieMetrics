import React from 'react'

export default function About() {
  return (
    <div className="flex flex-col h-screen">
      
      <div className="flex flex-row h-full">
          <div className="bg-red-600 w-1/2 font-sans text-white text-lg md:text-4xl p-16 text-center">Welcome to MovieMetrics <p className="text-sm md:text-lg md:px-44"> A movie platform to help you pick your favorite flix</p> 
          </div>
          <img className="w-1/2 object-cover" src="/movies.jpeg" alt="logo" />
      </div>

    <div className="flex min-h-[384px]">
        <div className="flex  justify-center items-center bg-red-300 w-1/3   ">
          <text className="text-white font-bold origin-center -rotate-90 text-xl md:text-6xl font-sans"> About us </text> </div>
        <div className="bg-slate-300 md:w-1/3 p-4 md:p-10 font-sans text-sm md:text-md font-semibold text-center" >MovieMetrics was founded on June 2022 by a couple of young college kids that were passionate about movies. Blah blah blah</div>
        <div className="bg-red-500 md:block md:w-1/3 hidden"></div>
    </div>

    <div className="flex h-full">
    <img className="w-1/2 object-cover" src="/movies3.jpeg" alt="logo" />
        <div className="bg-black w-1/2 text-white text-sm font-sans text-center p-10"> We need text here pls</div>
    </div>

    </div>
  )
}
