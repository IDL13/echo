import React from "react";
import axios from "axios";


export function Log() {

  function  handlerClick() {
    var username = document.getElementById('username') as HTMLInputElement
    var password = document.getElementById('password') as HTMLInputElement

    var user = {"username" : username.value, "password" : password.value}
    var Post = JSON.stringify(user)
    
    axios.post("/reg", Post).then((r)=>{console.log(r.data)})
  }

  return (
    <form onSubmit = {handlerClick} className="w-96 mt-auto ml-auto mb-auto mr-auto">

  <div className="mb-6">
    <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Username</label>
    <input id="username" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="User"></input>
  </div>

  <div className="mb-6">
    <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Password</label>
    <input id="password" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="1234" required></input>
  </div>

  <button type = "submit" className=" hover:font-bold text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Log in</button>
</form>
  )
}