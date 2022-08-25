// Un arreglo contiene N valores numéricos desordenados que pueden llegar a repetirse más de una vez, cree una función que entre todos esos valores encuentre el segundo número más grande.

// Ejemplo:

const arr = [3, 1, 4, 5, 0, -1, 2, 4, 2];
const arr1 = [23, 1022, 1022, 2013, 2013, 2012, 4, 56];
const arr2 = [35, 51, 45, 54, 50, 54, 51, 4, 51];

const secondLargestNumber = arr => {
  const setStruc = new Set(arr);
  const setArr = [...setStruc];
  const maxIdx = setArr.indexOf(Math.max(...setStruc));
  setArr.splice(maxIdx, 1);
  return Math.max(...setArr);
};

console.log(secondLargestNumber(arr));
console.log(secondLargestNumber(arr1));
console.log(secondLargestNumber(arr2));
