// const isBalanced = (s, caps) => {
//   let stack = [];
//   let flag = false;
//   caps.split('').forEach((cap, index) => {
//     if (index % 2 === 0 && cap !== caps[index + 1]) {
//       console.log(cap,caps[index+1],'masu')
//       s.split('').forEach(char => {
//         if (char === cap) {
//           stack.push(cap);
//           console.log('1',stack);
//         }
//       });
//     } else {
//       s.split('').forEach((char, idx) => {
//         if (char === cap) {
//           let temp = stack.pop()
//           console.log('2',temp, caps[index - 1]);
//           flag = temp === caps[index - 1];
//         }
//       });
//     }
//   });
//   console.log(caps, flag, stack, '<---3');

//   return !stack.length && flag;
// };

// const opposites = { '{': '}', '(': ')', "'": "'", '[': ']', '-': '-' };

// const isBalanced = (s, caps) => {
//   let stack = s.split('').filter(char => caps.includes(char));

//   if (stack.length % 2) return [false, 'a'];
//   for (let i = 0; i < Math.floor(stack.length / 2); i++) {
//     console.log(opposites[stack[i]], stack[stack.length - 1 - i]);
//     console.log(opposites[stack[i]] === stack[stack.length - 1 - i]);

//     if (opposites[stack[i]] !== stack[stack.length - 1 - i]) return false;
//   }
//   return true;

// };

// const opposite = { '{': '}', '(': ')', "'": "'", '[': ']', '-': '-' };

// const isBalanced = (s, caps) => {
//   let stack = [];
//   let flag = true;

//   s.split('').forEach((value, index, array) => {
//     if (value === '{' || value === '(' || value === "'" || value === '[' || value === '-')
//       stack.push(value);
//     else if (value === '}' || value === ')' || value === "'" || value === ']' || value === '-') {
//       let temp = stack.pop();
//      // console.log(opposite[temp], value);
//       if (opposite[temp] !== value) flag = false;
//     }
//   });

//   return !(stack.length % 2) && flag;
// };

// console.log(isBalanced('(Sensei says yes!)', '()'));
// console.log(isBalanced('(Sensei says no!', '()'));

// console.log(isBalanced('(Sensei [says] yes!)', '()[]'));
// console.log(isBalanced('(Sensei [says) no!]', '()[]'));

// console.log(isBalanced('Sensei says -yes-!', '--'));
// console.log(isBalanced('Sensei -says no!', '--'));



const opposite = { '{': '}', '(': ')', "'": "'", '[': ']', '-': '-' };

const isBalanced = (s, caps) => {
  let stack = [];
  let flag = true;

  s.split('').forEach((value, index, array) => {
    if (value === '{' || value === '(' || value === "'" || value === '[' || value === '-')
      stack.push(value);
    else if (value === '}' || value === ')' || value === "'" || value === ']' || value === '-') {
      let temp = stack.pop();
      console.log(opposite[temp], value);
      if (opposite[temp] !== value) flag = false;
    }
  });

  return !(stack.length % 2) && flag;
};

console.log(isBalanced('(Sensei says yes!)', '()'));
console.log(isBalanced('(Sensei says no!', '()'));

console.log(isBalanced('(Sensei [says] yes!)', '()[]'));
console.log(isBalanced('(Sensei [says) no!]', '()[]'));

console.log(isBalanced('Sensei says -yes-!', '--'));
console.log(isBalanced('Sensei -says no!', '--'));
