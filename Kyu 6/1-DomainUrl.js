/**
 * Write a function that when given a URL as a string, 
 * parses out just the domain name and returns it as a string. 
 * 
 * For example:
 * 
 * domainName("http://github.com/carbonfive/raygun") == "github" 
 * domainName("http://www.zombie-bites.com") == "zombie-bites"
 * domainName("https://www.cnet.com") == "cnet"
 */

// function domainName (url) {
//     const words = ['http', 'https', '://', 'www'].join('\\b|\\b');
//     url = url.replace(new RegExp(words, 'gi'), '').split('.').filter((str) => {
//         return /\S/.test(str)
//     })

//     return url[0]
// }

//alternative
function domainName(url){
    return url.match(/(?:http(?:s)?:\/\/)?(?:w{3}\.)?([^\.]+)/i)[1];
}

console.log(domainName("http://google.com")) // "google"
console.log(domainName("http://google.co.jp")) // "google"
console.log(domainName("www.xakep.ru")) // "xakep"
console.log(domainName("https://youtube.com")) // "youtube"
console.log(domainName("http://github.com/carbonfive/raygun")) // "github" 
console.log(domainName("http://www.zombie-bites.com")) // "zombie-bites"
console.log(domainName("https://www.cnet.com")) // "cnet"