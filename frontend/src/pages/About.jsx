import React from 'react'

export default function About() {
  return (
    <div className="flex flex-col h-screen">
      
      <div className="flex flex-row h-full">
          <div className="bg-red-600 w-1/2 font-sans text-white text-lg md:text-4xl p-16 text-center">Welcome to MovieMetrics <p className="text-sm md:text-lg md:px-44"> 
          Our online movie platform helps pick your favorite films! Our proprietary algorithm, based on the factors such as the movie title. genre, actors, and director, will pick the movies that we believe are the most relevant to you!
          </p> 
          </div>
          <img className="w-1/2 object-cover" src="/movies.jpeg" alt="logo" />
      </div>

    <div className="flex min-h-[384px]">
        <div className="flex  justify-center items-center bg-red-300 w-1/3   ">
          <text className="text-white font-bold origin-center -rotate-90 text-xl md:text-6xl font-sans"> About Us </text> </div>
        <div className="bg-slate-300 md:w-1/3 p-4 md:p-10 font-sans text-sm md:text-md font-semibold text-center" >MovieMetrics was founded on June 2022 by a couple of Red Ventures interns that were passionate about helping consumers pick the best and most relevant movies for them. We believe in helping consumers make the most informed decisions, which led us to creating our platform and algorithm - we want you to consume the most enjoyable content!</div>
        <div className="bg-red-500 md:block md:w-1/3 hidden"></div>
    </div>

    <div className="flex h-full">
    <img className="w-1/2 object-cover" src="/movies3.jpeg" alt="logo" />
        <div className="bg-black w-1/2 text-white text-lg font-sans text-center p-10"> 
        <h1>Our Proprietary Algorithm</h1>
        <p>When you search for a movie, our algorithm selects the most "relevant" films for you, in order, based on the following categories ranked in order of impotance:</p>
        <ol>
          <li>- Actors</li>
          <li>- Genre</li>
          <li>- Directors</li>
          <li>- Movie Title</li>
        </ol>
        <p>We believe that we have found the perfect formula, using these categories, to allow you - the consumer - to make the most informed choices when deciding the movie that you want to watch!</p>
        </div>
    </div>

    </div>
  )
}
