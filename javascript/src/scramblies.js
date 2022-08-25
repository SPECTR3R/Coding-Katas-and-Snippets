// Complete the function scramble(str1, str2) that returns true if a portion of str1 characters can be rearranged to match str2, otherwise returns false.
// Only lower case letters will be used (a-z). No punctuation or digits will be included.
// Performance needs to be considered
// Input strings s1 and s2 are null terminated.

const countChars = (array, count = {}) => {
  array.split('').forEach((elm, idx) => (count[elm] = count[elm] ? count[elm] + 1 : 1));
  return count;
};

const scramble = (s1, s2) => {
  const count1 = countChars(s1);
  const count2 = countChars(s2);
  for (let key in count2) {
    if (!(key in count1) || count1[key] < count2[key]) {
      return false;
    }
  }
  return true;
};
console.log(scramble('cedewaraaossoqqyt', 'codewars'));
