//Write a function that accepts a square matrix (N x N 2D array) and returns the determinant of the matrix.
const minor = (M, j) =>
  M.map(m => m.filter((el, ind) => ind !== j)).slice(1, M.length);


const determinant = M => {
  let sum = 0
  let prod = 0;
  if (!isNaN(M)) return M[0][0];
  if (M.length > 1) {
    for (j = 0; j < M.length; j++) {
      prod += M[0][j] * Math.pow(-1, j) * determinant(minor(M, j));
    //  sum += prod;
    }
    return prod;
  } else return M[0];
};

m0=[[1]]
m1 = [
  [1, 3],
  [2, 5],
];

m2 = [
  [2, 5, 3],
  [1, -2, -1],
  [1, 3, 4],
];

console.log(determinant(m2));
