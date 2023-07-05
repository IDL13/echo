import React from 'react';
import {Form} from "./components/Form"
import {Section1} from "./components/Section1"
import { Footer } from './components/Footer';
// import {Routes, Route} from "react-router-dom";

// import {Rest} from "./pages/Rest";
// import {Home} from "./pages/Home";

function App() {
  return (
    <>
    {/* <header>
    <nav>
      <li>
      <a href = "/" className='text-center text-2xl text-white mx-9'> Rest </a>
      </li>
        <li>
          <a href = "/rest" className='text-center text-2xl text-white mx-9'> Rest </a>
        </li>
      </nav>
    </header> */}

    <div>
    <Section1 />
    <Form />
    <Footer />
    </div>
    {/* <Routes>
    <Route path = "/" element = {<Home/>}></Route>
      <Route path = "/rest" element = {<Rest/>}></Route>
    </Routes> */}
    </>
  )
}

export default App;
