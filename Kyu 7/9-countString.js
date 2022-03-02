/**
 * The main idea is to count all the occurring characters in a string. 
 * If you have a string like aba, then the result should be {'a': 2, 'b': 1}.
 * 
 * What if the string is empty? Then the result should be empty object literal, {}.
 * 
 */

// function count (string) {  
//     let obj = {}
//     string = string.split('').sort()

//     for (let i = 0; i < string.length; i++) {
//         let num = 0;
//         for (let j = 0; j < string.length + 1; j++) {
//             if (string[j] == string[i]) {
//                 num++
//                 obj[string[j]] = num
//             }
//         }
//     }

//     return obj
// }

// alternative 
function count (string) {  
    var count = {};
    string.split('').forEach(function(s) {
       count[s] ? count[s]++ : count[s] = 1;
    });
    return count;
  }

console.log(count("aba")) // { a: 2, b: 1 }); 
console.log(count("")) // {});
