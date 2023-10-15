export function fetchAllProducts() {
  return new Promise(async (resolve) => {
    //todo: we will nor hardcore server url here
    const response = await fetch("http://localhost:8080/products");
    const data = await response.json();
    resolve({ data });
  });
}

export function fetchAllProductsByFilters(filter, sort, pagination) {
  // filter ={category: [smartphone,laptops]}
  // sort ={_sort:price, _order:desc}
  //pagination ={_page:1,_limit =10}
  console.log("asdjksld");
  let queryString = "";
  for (let key in filter) {
    const categoryValues = filter[key];

    if (categoryValues.length) {
      console.log("hello");
      const lastCotegoryValue = categoryValues[categoryValues.length - 1];
      console.log(lastCotegoryValue);
      queryString += `${key}=${lastCotegoryValue}&`;
    }
  }
  for (let key in sort) {
    queryString += `${key}=${sort[key]}&`;
  }
  for (let key in pagination) {
    queryString += `${key}=${pagination[key]}&`;
  }
  return new Promise(async (resolve) => {
    //todo: we will nor hardcore server url here
    // todo
    const response = await fetch(
      "http://localhost:8080/products?" + queryString
    );
    const data = await response.json();
    resolve({ data });
    //console.log({data})
    // const totalItems = await response.headers.get("X-Total-Count");
    // resolve({ data: { products: data, totalItems: +totalItems } });
  });
}
