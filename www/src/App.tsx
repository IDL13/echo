import React from 'react';

import { Link, Route, Routes } from 'react-router-dom';
import { Update } from './pages/Update';
import { Home } from './pages/Home';
import { NotFound } from './pages/NotFound';
import { Del } from './pages/Delete';
import { Select } from './pages/SelectAll';
import { RedisSet } from './pages/RedisSet'
import { RedisGet } from './pages/RedisGet'
import { RedisLog } from './pages/RedisLog'


function App() {
  return (
    <>
    <header>
      <nav>
        <ul>
          <li className='m-4 bg-neutral-900'>
            <Link to = "/" className='m-16 text-center text-2xl text-white'>Home</Link>
            <Link to = "/update" className='m-16 text-center text-2xl text-white'>Update</Link>
            <Link to = "/delete" className='m-16 text-center text-2xl text-white'>Delete</Link>
            <Link to = "/select_all" className='m-16 text-center text-2xl text-white'>Select All</Link>
            <Link to = "/redis_set" className='m-16 text-center text-2xl text-white'>Redis Set</Link>
            <Link to = "/redis_get" className='m-16 text-center text-2xl text-white'>Redis Get</Link>
            <Link to = "/log" className='m-16 text-center text-2xl text-orange-500'>Login/Logout</Link>
          </li>
        </ul>
      </nav>
    </header>

    <Routes>
      <Route path = "/" element = {<Home/>}></Route>
      <Route path = "/update" element = {<Update/>}></Route>
      <Route path = "/delete" element = {<Del/>}></Route>
      <Route path = "/select_all" element = {<Select/>}></Route>
      <Route path = "/redis_set" element = {<RedisSet/>}></Route>
      <Route path = "/redis_get" element = {<RedisGet/>}></Route>
      <Route path = "/log" element = {<RedisLog/>}></Route>
      <Route path = "*" element = {<NotFound/>}></Route>
    </Routes>
    </>
  )
}

export default App;
