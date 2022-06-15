import React, {useState, useEffect} from 'react'
import axios from 'axios';

export default function Home() {
  const [search, setSearch] = useState("");

  const handleSearch = async (event) => {
    event.preventDefault();
    const movie = await axios.get(`localhost:8081/?t=${encodeURI(search)}`);
    console.log(movie);
  }


  return (
    <div className="bg-movie-posters flex h-screen justify-center items-center">
        <div className="flex justify-center">

          <form onSubmit={handleSearch}>
            <input value={search} onChange={(e)=>setSearch(e.target.value)} type="text" className=" px-4 inline-flex rounded-l w-40 md:w-64 lg:w-96 h-10 bg-gray-200" placeholder="Search Movie by Title..."></input>
            <button type="submit" className =" w-12 rounded-r bg-red-600 hover:bg-red-700">
            <svg xmlns="http://www.w3.org/2000/svg" className="m-auto h-5 w-5" viewBox="0 0 20 20" fill="white">
  <path fillRule="evenodd" d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z" clipRule="evenodd" />
</svg>
            </button>
          </form>
        </div>
        </div>
    
  );
}
