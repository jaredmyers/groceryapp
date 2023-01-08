import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import Auth from './components/Auth'
import SignIn from './components/SignIn'
import NewItem from './components/NewItem'
import ShowPath from './components/ShowPath'
import Dashboard from './components/Dashboard'
import GroceryList from './components/GroceryList'
import reportWebVitals from './reportWebVitals';
import { createBrowserRouter, RouterProvider, Route } from 'react-router-dom'

const router = createBrowserRouter ([
		{
				path: "/",
				element: <SignIn />,
		},
		{
				path: "/dashboard",
				element: <Dashboard />,
		},
		{
				path: "/grocerylist",
				element: <GroceryList />,
		},
		{
				path: "/newitem",
				element: <NewItem />,
		},
		{
				path: "/showpath",
				element: <ShowPath />,
		},

])

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
		<RouterProvider router={router}/>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
