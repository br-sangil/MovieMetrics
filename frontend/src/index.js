import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from './pages/Home';
import Random from './pages/Random';
import Layout from './pages/Layout';
import About from './pages/About';
export default function App(){
  return (
    <BrowserRouter>
    <Routes>
        <Route path="/" element={<Layout />}>
        <Route index element={<Home />} />
        <Route path="/random" element={<Random/>} />
        <Route path="/about" element={<About/>} />
      </Route>
    </Routes>
  </BrowserRouter>

  );
}
ReactDOM.render(<App />, document.getElementById("root"));
