import React from 'react';

import { Link, Route, Routes } from 'react-router-dom';
import { Rest } from './pages/Rest';
import { Home } from './pages/Home';


function App() {
  return (
    <>
    <header>
      <Link to = "/" className='m-10 text-center text-6xl text-white'>Home</Link>
      <Link to = "/rest" className='m-10 text-center text-6xl text-white'>Rest</Link>
    </header>

    <Routes>
      <Route path = "/" element = {<Home/>}></Route>
      <Route path = "/rest" element = {<Rest/>}></Route>
    </Routes>
    </>
  )
}

export default App;
