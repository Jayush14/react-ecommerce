import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import { fetchAllProducts,fetchAllProductsByFilters } from './ProductAPI';

const initialState = {
  Products: [],
  status: 'idle',
  totalItems:0,
};

export const fetchAllProductsByFiltersAsync= createAsyncThunk(
  'Product/fetchAllProductsByFilters',
    async ({filter,sort,pagination}) => {
      const response = await fetchAllProductsByFilters(filter,sort,pagination);
      
    return response.data;
  }
);

export const fetchAllProductsAsync = createAsyncThunk(
  'Product/fetchAllProducts',
  async () => {
    const response = await fetchAllProducts();
    
    return response.data;
  }
);

export const ProductSlice = createSlice({
  name: 'Product',
  initialState,
  reducers: {
    increment: (state) => {
      
     // state.value += 1;
    },
  },
 
  extraReducers: (builder) => {
    builder
      .addCase(fetchAllProductsAsync.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(fetchAllProductsAsync.fulfilled, (state, action) => {
        state.status = 'idle';
        state.Products = action.payload;
      })
      .addCase(fetchAllProductsByFiltersAsync.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(fetchAllProductsByFiltersAsync.fulfilled, (state, action) => {
        state.status = 'idle';
        state.Products = action.payload;
     //   state.totalItems = action.payload.totalItems;
      });
  },
});

export const { increment} = ProductSlice.actions;

 
export const selectAllProducts = (state) => state.Product.Products;
//export const selectTotalItems = (state) => state.product.totalItems;

export default ProductSlice.reducer;
