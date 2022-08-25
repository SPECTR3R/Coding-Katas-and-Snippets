// Half diamond star pattern using a forEach iterator

let printStars = howMany => console.log(' *'.repeat(howMany));

[1, 2, 3, 4, 5, 6, 7].forEach((num, it, arr) => {
  let p = Math.round(arr.length / 2) - 1;
  if (it < p) {
    printStars(arr[it] + num - 1);
    printStars(arr[it] + num);
  } else if (it === p && (arr.length - 1) % 2 !== 0) {
    printStars(arr[arr.length - 1] - 1);
  } else if (it === p && (arr.length - 1) % 2 == 0) {
    printStars(arr[arr.length - 1]);
  } else {
    printStars(2 * (arr.length - it));
    printStars(2 * (arr.length - it) - 1);
  }
});
