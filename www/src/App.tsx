import React from 'react';

import { Link, Route, Routes } from 'react-router-dom';
import { Update } from './pages/Update';
import { Home } from './pages/Home';
import { NotFound } from './pages/NotFound';


function App() {
  return (
    <>
    <header>
      <nav>
        <ul>
          <li className='m-4 bg-neutral-900'>
            <Link to = "/" className='m-10 text-center text-2xl text-white'>Home</Link>
            <Link to = "/update" className='m-10 text-center text-2xl text-white'>Update</Link>
            <Link to = "/delete" className='m-10 text-center text-2xl text-white'>Delete</Link>
            <Link to = "/select_all" className='m-10 text-center text-2xl text-white'>Select All</Link>
          </li>
        </ul>
      </nav>
    </header>

    <Routes>
      <Route path = "/" element = {<Home/>}></Route>
      <Route path = "/update" element = {<Update/>}></Route>
      <Route path = "*" element = {<NotFound/>}></Route>
    </Routes>
    </>
  )
}

export default App;
