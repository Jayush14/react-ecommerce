import React from 'react';

import { Counter } from './features/counter/Counter';
import './App.css';

import SignupPage from './Pages/SignupPage';
import Home from './Pages/Home';
import LoginPage from './Pages/LoginPage';
import CartPage from './Pages/CartPage';
import CheckOutPage from './Pages/CheckOutPage';
import ProductDetailPage from './Pages/ProductDetailPage';
 
import { createRoot } from "react-dom/client";
import {
  createBrowserRouter,
  RouterProvider,
  Route,
  Link,
} from "react-router-dom";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home></Home>,
  },
  {
    path: "/login",
    element: <LoginPage></LoginPage>,
  },
  {
    path: "/signup",
    element: <SignupPage> </SignupPage>,
  },
  {
    path: "/cart",
    element: <CartPage></CartPage>,
  },
  {
    path: "/checkout",
    element: <CheckOutPage></CheckOutPage>,
  },
  {
    path: "/product-detail",
    element: <ProductDetailPage></ProductDetailPage>,
  },
]);




function App() {
  return (
    <div className="App">
     <RouterProvider router={router} />
    </div>
  );
}

export default App;
 