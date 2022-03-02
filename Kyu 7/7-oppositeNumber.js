/**
 * Very simple, given an integer or a floating-point number, find its opposite.
 * Examples:
 * 
 * 1: -1
 * 14: -14
 * -34: 34
 */

// function opposite (number) {
//     let x = String(number)
//     if (x.includes('-')) {
//         x = Math.abs(x)
//     } else {
//         x = Math.abs(x) * -1
//     }
//     return x
// }

// alternative 1 
function opposite (number) {
    return -number
}

console.log(opposite(1)) // -1
console.log(opposite(14)) // -14
console.log(opposite(-34)) // 34
  