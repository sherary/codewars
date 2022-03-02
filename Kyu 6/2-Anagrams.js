/**
 * What is an anagram? 
 * Well, two words are anagrams of each other if they both contain the same letters. 
 * 
 * For example:
 * 'abba' & 'baab' == true
 * 'abba' & 'bbaa' == true
 * 'abba' & 'abbba' == false
 * 'abba' & 'abca' == false
 * 
 * Write a function that will find all the anagrams of a word from 
 * a list. You will be given two inputs a word and an array with words. 
 * You should return an array of all the anagrams or an empty array 
 * if there are none. For example:
 * 
 * anagrams('abba', ['aabb', 'abcd', 'bbaa', 'dada']) => ['aabb', 'bbaa']
 * anagrams('racer', ['crazer', 'carer', 'racar', 'caers', 'racer']) => ['carer', 'racer']
 * anagrams('laser', ['lazing', 'lazy',  'lacer']) => []
 */

// function anagrams (word, arr) {
//     let tmp = [];
//     let result = [];
//     let str = [];
//     word = word.split('').sort().join('')
//     arr.forEach(element => {
//         element = element.split('').sort().join('')
//         tmp.push(element)
//     });

//     for (let i = 0; i < tmp.length; i++) {
//         if (tmp[i] == word) {
//             str.push(i)
//         }
//     }

//     for (let i = 0; i < str.length; i++) {
//         result.push(arr[str[i]])
//     }
    
//     return result
// }

//alternative
function anagrams(word, words) {
    word = word.split('').sort().join('');
    return words.filter(function(v) {return word == v.split('').sort().join('');});
}

console.log(anagrams('abba', ['aabb', 'abcd', 'bbaa', 'dada'])) //['aabb', 'bbaa']
console.log(anagrams('racer', ['crazer', 'carer', 'racar', 'caers', 'racer']))// ['carer', 'racer']
console.log(anagrams('laser', ['lazing', 'lazy',  'lacer']))// []