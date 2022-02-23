let createPassBut=document.getElementById("createPass");
let copyPass=document.getElementById("passwordLabel");
let addRecord=document.getElementById("addRecord");




// Generate Password
function generatePassword() {
    var pass= Math.random().toString(36).slice(2) + Math.random().toString(36).slice(2)
    createLabelAndAppendPassword(pass)
}


function copyText() {
    
    var copy=document.getElementsByTagName('pre')[0].innerHTML;
    navigator.clipboard.writeText(copy);
    copyPass.innerHTML = "Copied";
    
    
}

function createLabelAndAppendPassword(generratedPassword) {
    var passwordlabel=document.getElementById("passwordLabel");
    passwordLabel.innerHTML=""
    passwordlabel.innerHTML = generratedPassword;

}




// Add Records function
function addRecordfunc() {
var myHeaders = new Headers();
myHeaders.append("Content-Type", "application/json");
// myHeaders.append("Access-Control-Allow-Origin", "*")

var raw = JSON.stringify({
  "email": "emirlerogluhalil@gmail.com",
  "password": "password",
  "domain": "glinkedin.com"
});

var requestOptions = {
  method: 'POST',
  headers: myHeaders,
  body: raw,
  redirect: 'follow'
};

fetch("http://localhost:8080/addRecord", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
}

createPassBut.addEventListener("click",generatePassword)
copyPass.addEventListener("click",copyText)
addRecord.addEventListener("click",addRecordfunc)
