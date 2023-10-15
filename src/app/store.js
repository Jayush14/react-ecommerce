import { configureStore } from '@reduxjs/toolkit';
import ProductReducer from "../features/Product/ProductSlice";

export const store = configureStore({
  reducer: {
    Product : ProductReducer,
  },
});
