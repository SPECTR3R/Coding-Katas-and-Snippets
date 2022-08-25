class RomanNumeral {
  constructor() {
    this.romanV = {
      M: 1000,
      D: 500,
      C: 100,
      L: 50,
      X: 10,
      V: 5,
      I: 1,
    };
  }

  fromRoman(str) {
    return str.split('').forEach((char, index, array) => {
      
    });
  }
}
const RomanNumerals = new RomanNumeral(); // ReferenceError

console.log(RomanNumerals.fromRoman('MDCLXIX'));
