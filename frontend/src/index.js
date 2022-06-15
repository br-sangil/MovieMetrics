import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Layout from './pages/Layout';
import Home from './pages/Home';
import Random from './pages/Random';
import './index.css'; 

export default function App(){
  return (
    <BrowserRouter>
    <Routes>
      <Route path="/" element={<Layout />}>
        <Route index element={<Home />} />
        <Route path="random" element={<Random/>} />
      </Route>
    </Routes>
  </BrowserRouter>

  );
}
ReactDOM.render(<App />, document.getElementById("root"));