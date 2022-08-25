// Write a function that flattens an Array of Array objects into a flat Array. Our function must flatten any given array input
// por ejemplo [1, 2, [3, 4]];

// Erwin Bastardo16:36
// [1, 2, 3, 4]
// [1, 2, [3, 4]]; : [1, 2, 3, 4]
// [1, 2, [3, 4, [5, 6]]]: [1, 2,3, 4, 5, 6]
// [1, 2, [3, 4, [5, 6, [7, 8, [9, 10]]]]];:
// [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

function flattenedArray(arr) {
  const newArr = []
  arr.forEach((element) => {
    if (Array.isArray(element)) {
      newArr.push(...flattenedArray(element))
    } else {
      newArr.push(element)
    }
  })
  return newArr
}
let myArr = [1, 2, 3, 4]
myArr = [1, 2, [3, 4]]
myArr = [1, 2, [3, 4, [5, 6]]]
myArr = [1, 2, [3, 4, [5, 6, [7, 8, [9, 10]]]]]

const result = flattenedArray(myArr)
console.log(result)
