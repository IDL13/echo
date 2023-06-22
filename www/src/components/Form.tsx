import React from "react";
import axios from "axios";


export function Form() {

  function  handlerClick() {
    const number = document.getElementById('number') as HTMLInputElement
    const date = document.getElementById('date') as HTMLInputElement
    const cvv = document.getElementById('cvv') as HTMLInputElement

    const card = `"number":${number.value}, "date":${date.value},"CVV":${cvv.value}`
    var Post = JSON.stringify(card)
    
    axios.post("/addOne").then((r)=>{console.log(r)})
  }

  return (
    <form onSubmit = {handlerClick} className="w-96 mt-auto ml-auto mb-auto mr-auto">

  <div className="mb-6">
    <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Number your card</label>
    <input id="number" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="4000001234567899"></input>
  </div>

  <div className="mb-6">
    <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Date</label>
    <input id="date" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="12.12" required></input>
  </div>

  <div className="mb-6">
    <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">CVV</label>
    <input type="password" id="cvv" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="123" required></input>
  </div>

  <button type = "submit" className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Save</button>
</form>
  )
}